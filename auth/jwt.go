package auth

import (
	"encoding/json"
	"github.com/Nerzal/gocloak"
	"net/http"
	"os"
	"strings"
)

var (
	clientId     string // = "microportal"
	clientSecret string // = "b56f5388-f2b6-4f9c-be59-48fed7294f43"
	realm        string // = "MicroportalRealm"
	keycloakUrl  string // = "http://localhost:7000/"
	pathPrefix   string
)

func init() {
	clientId = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm = os.Getenv("KEYCLOAK_REALM")
	keycloakUrl = os.Getenv("KEYCLOAK_URL")
	pathPrefix = os.Getenv("PATH_PREFIX")
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if pathPrefix == "" {
			pathPrefix = "/menu-service"
		}
		notAuth := []string{pathPrefix + "/public"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if strings.HasPrefix(requestPath, value) {
				next.ServeHTTP(w, r)
				return
			}
		}

		client := gocloak.NewClient(keycloakUrl)

		accessToken := r.Header.Get("Authorization")

		introspect, err := client.RetrospectToken(accessToken[7:], clientId, clientSecret, realm)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(err.Error())
			return
		}

		if !introspect.Active {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Invalid token!")
			return
		}

		next.ServeHTTP(w, r)
	})
}
