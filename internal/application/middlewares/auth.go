package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/ahmetilboga2004/internal/domain/interfaces"
	httphelper "github.com/ahmetilboga2004/internal/infrastructure/utils/httpHelper"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type authMiddleware struct {
	jwtService interfaces.IJwtService
}

func NewAuthMiddleware(jwtService interfaces.IJwtService) *authMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

func (m *authMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenClaims, err := m.jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		claims, ok := tokenClaims.Claims.(jwt.MapClaims)
		if !ok || !tokenClaims.Valid {
			next.ServeHTTP(w, r)
			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(UserIDKey)

		if userID == nil {
			httphelper.ErrorResponse(w, http.StatusUnauthorized, "You must be logged in to access this resource", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) GuestOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(UserIDKey) != nil {
			httphelper.ErrorResponse(w, http.StatusForbidden, "This resource is only accessible to guests", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
