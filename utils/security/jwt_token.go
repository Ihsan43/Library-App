package security

import (
	"errors"
	"fmt"
	"library_app/config"
	"library_app/model"
	"library_app/model/dto"
	modelutil "library_app/utils/model_util"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *model.User) (dto.AuthResponseDto, error) {
	if user == nil {
		return dto.AuthResponseDto{}, errors.New("user cannot be nil")
	}
	if user.ID == "" || user.Role == "" {
		return dto.AuthResponseDto{}, errors.New("invalid user data: ID and Role are required")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if cfg.TokenConfig.JwtSigningMethod == nil || len(cfg.TokenConfig.JwtSignatureKey) == 0 {
		return dto.AuthResponseDto{}, errors.New("JWT configuration is not initialized")
	}
	if cfg.TokenConfig.AccessTokenLifeTime <= 0 {
		return dto.AuthResponseDto{}, errors.New("AccessTokenLifeTime must be greater than 0")
	}

	now := time.Now().UTC()
	end := now.Add(cfg.TokenConfig.AccessTokenLifeTime)

	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		UserId: user.ID,
		Role:   user.Role,
	}

	token := jwt.NewWithClaims(cfg.TokenConfig.JwtSigningMethod, claims)
	newToken, err := token.SignedString(cfg.TokenConfig.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{
		Token: newToken,
	}, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verifikasi metode signing
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method.Alg() != cfg.JwtSigningMethod.Name {
			return nil, fmt.Errorf("Invalid token signing method")
		}
		// Kembalikan kunci signature
		return []byte(cfg.JwtSignatureKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid parse token: %s", err.Error())
	}

	// Verifikasi klaim
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token: claims verification failed")
	}

	// Verifikasi issuer (iss)
	if claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("Invalid token: issuer mismatch")
	}

	return claims, nil
}
