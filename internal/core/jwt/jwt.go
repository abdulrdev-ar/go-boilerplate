package jwt

import (
	"os"
	"strconv"
	"time"

	appJWT "github.com/golang-jwt/jwt/v5"
	"github.com/inienam06/go-boilerplate/internal/exception"
)

type JWTConfig struct {
	SecretKey string
	ExpiresIn int
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	appJWT.RegisteredClaims
}

func InitJWTConfig() *JWTConfig {
	expires := os.Getenv("JWT_EXPIRES_IN")
	exInt, _ := strconv.Atoi(expires)

	return &JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET"),
		ExpiresIn: exInt,
	}
}

// GenerateToken generates a signed JWT token
func GenerateToken(userID uint, email string) (string, error) {
	JWTConfig := InitJWTConfig()

	now := time.Now()
	expiresInSeconds := JWTConfig.ExpiresIn
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: appJWT.RegisteredClaims{
			IssuedAt:  appJWT.NewNumericDate(now),
			ExpiresAt: appJWT.NewNumericDate(now.Add(time.Duration(expiresInSeconds) * time.Second)),
			NotBefore: appJWT.NewNumericDate(now),
		},
	}
	token := appJWT.NewWithClaims(appJWT.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTConfig.SecretKey))
}

// ParseToken verifies token and returns claims
func ParseToken(tokenStr string) (*Claims, error) {
	JWTConfig := InitJWTConfig()
	token, err := appJWT.ParseWithClaims(tokenStr, &Claims{}, func(tok *appJWT.Token) (interface{}, error) {
		// enforce signing method
		if _, ok := tok.Method.(*appJWT.SigningMethodHMAC); !ok {
			return nil, exception.NewUnauthorizedException("invalid signing method")
		}
		return []byte(JWTConfig.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, exception.NewUnauthorizedException("invalid token")
}
