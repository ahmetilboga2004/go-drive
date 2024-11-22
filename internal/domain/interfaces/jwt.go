package interfaces

import "github.com/golang-jwt/jwt/v5"

type IJwtService interface {
	GenerateAccessToken(userId uint, claims map[string]any) (string, error)
	GenerateRefreshToken(userId uint) (string, error)
	ValidateAccessToken(tokenString string) (*jwt.Token, error)
	ValidateRefreshToken(tokenString string) (*jwt.Token, error)
	GetTokenClaims(token *jwt.Token) (map[string]any, error)
}
