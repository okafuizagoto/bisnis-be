package auth

import (
	"bisnis-be/internal/config"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	cfg *config.Config
	// configSecret = string(cfg.JWT.Secret)
	// jwtSecret    = []byte(configSecret) // default bisa di-override dari config
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresAt   int64  `json:"expires_at"`
	TokenType   string `json:"token_type"`
}

// GenerateJWT membuat token JWT
func GenerateJWT(userID string, duration time.Duration) (Token, error) {
	now := time.Now()
	exp := now.Add(duration)

	claims := jwt.MapClaims{
		"sub":  userID,
		"user": userID,
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
		"exp":  exp.Unix(),
		"iss":  "BISNIS-BE",
	}
	cfg, _ = config.Get()
	jwtSecret := []byte(cfg.JWT.Secret)

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := tokenObj.SignedString(jwtSecret)
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken: tokenStr,
		ExpiresIn:   int64(duration.Seconds()),
		ExpiresAt:   exp.Unix(),
		TokenType:   "Bearer",
	}, nil
}

// ValidateJWT memvalidasi token dan mengembalikan claims
func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	fmt.Println("test1")
	cfg, _ = config.Get()
	jwtSecret := []byte(cfg.JWT.Secret)
	tokenObj, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		fmt.Println("test1-1")
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("test1-2")
			return []byte("BISNIS-BE"), errors.New("invalid signing method")
		}
		fmt.Println("test1-3")
		return jwtSecret, nil
	})
	fmt.Printf("test1-4 %+v", tokenObj)
	if err != nil {
		fmt.Println("test1-5", err)
		return nil, err
	}
	fmt.Println("test2")

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok || !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}
	fmt.Println("test3")

	return claims, nil
}
