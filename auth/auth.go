package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig *oauth2.Config

func init() {
	// Prefer REDIRECT_URI (from .env) for consistency with the project file,
	// fall back to OAUTH_REDIRECT_URL, then to a sensible default.
	redirect := strings.TrimSpace(os.Getenv("REDIRECT_URI"))
	if redirect == "" {
		redirect = strings.TrimSpace(os.Getenv("OAUTH_REDIRECT_URL"))
	}
	if redirect == "" {
		redirect = "http://localhost:8080/api/auth/callback/google"
	}
	clientID := strings.TrimSpace(os.Getenv("CLIENT_ID"))
	clientSecret := strings.TrimSpace(os.Getenv("CLIENT_SECRET"))
	googleOAuthConfig = &oauth2.Config{
		RedirectURL:  redirect,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	// Mask and log the client id and redirect for debugging (do not log secrets)
	cid := os.Getenv("CLIENT_ID")
	masked := cid
	if len(cid) > 8 {
		masked = cid[:4] + "..." + cid[len(cid)-4:]
	}
	log.Printf("OAuth CLIENT_ID=%s REDIRECT=%s", masked, redirect)
}

func generateState(c *fiber.Ctx) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// use RawURLEncoding to avoid padding (=) and keep the cookie a bit shorter
	state := base64.RawURLEncoding.EncodeToString(b)
	c.Cookie(&fiber.Cookie{ // set a short-lived cookie to verify state
		Name:    "oauthstate",
		Value:   state,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	})
	return state, nil
}

// Login starts the OAuth2 flow and redirects the user to Google's consent screen.
func Login(c *fiber.Ctx) error {
	// Diagnostic logging: log request header and cookie size to help debug 431 errors
	cookieHeader := c.Get("Cookie")
	totalHeaderLen := 0
	c.Request().Header.VisitAll(func(k, v []byte) {
		totalHeaderLen += len(k) + len(v)
	})
	log.Printf("Auth Login request headers total bytes=%d cookieHeaderBytes=%d", totalHeaderLen, len(cookieHeader))

	state, err := generateState(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to generate oauth state")
	}
	url := googleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	// Log the generated auth URL with client_id masked for diagnosis
	// mask client_id value in the URL
	maskedUrl := url
	if cid := os.Getenv("CLIENT_ID"); cid != "" {
		maskedCid := cid
		if len(cid) > 8 {
			maskedCid = cid[:4] + "..." + cid[len(cid)-4:]
		}
		maskedUrl = strings.ReplaceAll(maskedUrl, cid, maskedCid)
	}
	log.Printf("Auth URL: %s", maskedUrl)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// Callback handles the OAuth2 callback from Google, exchanges the code for a token
// and fetches basic user info. It returns the user info and token as JSON.
func Callback(c *fiber.Ctx) error {
	state := c.Query("state")
	cookieState := c.Cookies("oauthstate")
	if state == "" || cookieState == "" || state != cookieState {
		return c.Status(fiber.StatusBadRequest).SendString("invalid oauth state")
	}
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("code not found")
	}

	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to exchange token")
	}

	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to get userinfo")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to read userinfo")
	}

	var gu map[string]interface{}
	if err := json.Unmarshal(body, &gu); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to parse userinfo")
	}

	// At this point you would typically create or lookup the user in your DB
	// and create a session/jwt. For now, return the Google user info + token.
	return c.JSON(fiber.Map{
		"user":  gu,
		"token": token,
	})
}
