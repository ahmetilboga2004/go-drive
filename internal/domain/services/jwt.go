package services

import (
	"errors"
	"time"

	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExp     time.Duration
	RefreshTokenExp    time.Duration
}

type jwtService struct {
	config *JwtConfig
}

func NewJwtService(config *JwtConfig) interfaces.IJwtService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateAccessToken(userId string, claims map[string]interface{}) (string, error) {
	now := time.Now()
	tokenClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     now.Add(s.config.AccessTokenExp).Unix(),
		"iat":     now.Unix(),
	}

	for key, value := range claims {
		tokenClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	return token.SignedString([]byte(s.config.AccessTokenSecret))
}

func (s *jwtService) GenerateRefreshToken(userId string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     now.Add(s.config.RefreshTokenExp).Unix(),
		"iat":     now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.RefreshTokenSecret))
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.AccessTokenSecret), nil
	})
}

func (s *jwtService) ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.RefreshTokenSecret), nil
	})
}

func (s *jwtService) GetTokenClaims(token *jwt.Token) (map[string]interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
