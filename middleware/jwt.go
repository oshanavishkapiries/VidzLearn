package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Cenzios/pf-backend/pkg/response"
	"github.com/Cenzios/pf-backend/pkg/utils"
)

// ContextKey is a custom type for context keys to avoid collisions
// [INFO] @better-comments:info
// Use ContextKeyUser for storing user info in context
type ContextKey string

const ContextKeyUser ContextKey = "user"

// JWTAuthMiddleware verifies JWT and sets user info in context
// [INFO] @better-comments:info
// Usage: Wrap protected handlers with this middleware
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(w, "Missing or invalid Authorization header")
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		jwtPayload, err := utils.VerifyJWT(token)
		if err != nil {
			response.Unauthorized(w, "Invalid or expired token")
			return
		}
		// Set user info in context
		ctx := context.WithValue(r.Context(), ContextKeyUser, jwtPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext extracts JWT payload from context
// [INFO] @better-comments:info
// Use in handlers to get authenticated user info
func GetUserFromContext(ctx context.Context) *utils.JWTPayload {
	user, ok := ctx.Value(ContextKeyUser).(*utils.JWTPayload)
	if !ok {
		return nil
	}
	return user
}
