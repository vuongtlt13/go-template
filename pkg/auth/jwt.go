package auth

import (
	"errors"
	"sync"
	"time"

	"yourapp/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	config *config.JWTConfig
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"user_id"`
}

var (
	instance *JWTManager
	once     sync.Once
	initErr  error
)

// GetJWTManager returns the singleton instance of JWTManager
func GetJWTManager() *JWTManager {
	once.Do(func() {
		cfg := config.GetConfig()
		instance = &JWTManager{config: &cfg.JWT}
	})
	if initErr != nil {
		panic(initErr)
	}
	return instance
}

func (m *JWTManager) GenerateToken(userID uint64) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.config.ExpirePeriod)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.config.Secret))
}

func (m *JWTManager) VerifyToken(tokenStr string) (uint64, error) {
	claims, err := m.ValidateToken(tokenStr)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}

func (m *JWTManager) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(m.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
