package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/GigaDesk/eardrum-server/pkg/jwt"
)

type favContextKey string

var k = favContextKey("user")

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if header == "" {
				log.Println("Allowed unauthenticated user")
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			credentials, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), k, credentials)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the token credentials from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *jwt.TokenCredentials {
	user, _ := ctx.Value(k).(*jwt.TokenCredentials)
	return user
}
