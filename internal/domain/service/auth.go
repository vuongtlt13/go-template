package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"yourapp/internal/domain/repository"
	"yourapp/pkg/auth"
	"yourapp/pkg/database"
)

type Credential struct {
	Email    string
	Password string
}

// AuthService defines the auth service interface
type AuthService interface {
	Login(ctx context.Context, credential Credential) (string, error)
}

// authServiceImpl implements the UserService interface
type authServiceImpl struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	jwtManager *auth.JWTManager
}

func (a authServiceImpl) Login(ctx context.Context, credential Credential) (string, error) {
	user, err := a.userRepo.FindByEmail(ctx, credential.Email, a.db)
	if err != nil {
		return "", fmt.Errorf("error when finding user: %w", err)
	}
	token, err := a.jwtManager.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("error when generating token: %w", err)
	}
	return token, nil
}

// NewAuthService creates a new user service instance
func NewAuthService() AuthService {
	return &authServiceImpl{
		userRepo:   repository.NewUserRepository(),
		jwtManager: auth.GetJWTManager(),
		db:         database.GetDatabase(),
	}
}
