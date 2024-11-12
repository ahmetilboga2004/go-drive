package interfaces

import "github.com/golang-jwt/jwt/v5"

type IJwtService interface {
	GenerateAccessToken(userId string, claims map[string]interface{}) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	ValidateAccessToken(tokenString string) (*jwt.Token, error)
	ValidateRefreshToken(tokenString string) (*jwt.Token, error)
	GetTokenClaims(token *jwt.Token) (map[string]interface{}, error)
}
