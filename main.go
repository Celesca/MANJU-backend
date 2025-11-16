package main

import (
	"fmt"
	"log"
	"manju/backend/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		// no-op; env might be set already
	}

	r := router.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Bind to loopback by default to avoid Windows wildcard socket permission
	// issues. Override with HOST env var (e.g. 0.0.0.0) if you need external access.
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("🚀 Server running on %s", addr)
	if err := r.Listen(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
