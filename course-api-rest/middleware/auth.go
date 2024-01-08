package middleware

import (
	"net/http"
	"strings"

	"api.com/go/rest-ws/models"
	"api.com/go/rest-ws/server"
	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEED = []string{
		"/",
		"/signup",
		"/login",
	}
)

func shouldCheckToken(route string) bool {
	for _, v := range NO_AUTH_NEED {
		if strings.Contains(route, v) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
