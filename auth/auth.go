// auth.go
package main

import (
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, "loginHandler can't choose provider", http.StatusInternalServerError)
		}
		authUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, "loginHandler can't choose provider", http.StatusInternalServerError)
		}
		w.Header().Set("Location", authUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
