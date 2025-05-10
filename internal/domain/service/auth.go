package service

import (
	"context"
	"yourapp/internal/domain/model"
	"yourapp/internal/domain/repository"
	"yourapp/pkg/auth"
	"yourapp/pkg/database"

	"gorm.io/gorm"
)

type Credential struct {
	Email    string
	Password string
}

// AuthService defines the auth service interface
type AuthService interface {
	Login(ctx context.Context, cred Credential) (string, error)
	Register(ctx context.Context, cred Credential) error
}

// authServiceImpl implements the UserService interface
type authService struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	jwtManager *auth.JWTManager
}

var (
	authServiceInstance *authService
)

// NewAuthService creates a new auth service using singleton instances
func NewAuthService() AuthService {
	if authServiceInstance == nil {
		authServiceInstance = &authService{
			db:         database.GetDatabase(),
			userRepo:   repository.NewUserRepository(),
			jwtManager: auth.GetJWTManager(),
		}
	}
	return authServiceInstance
}

func (a *authService) Login(ctx context.Context, cred Credential) (string, error) {
	user, err := a.userRepo.GetByEmail(ctx, cred.Email)
	if err != nil {
		return "", err
	}

	if !user.IsActive {
		return "", err
	}

	// Verify password
	if err := user.CheckPassword(cred.Password); err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := a.jwtManager.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authService) Register(ctx context.Context, cred Credential) error {
	// Create user
	user := &model.User{
		Email:    cred.Email,
		Password: cred.Password,
		IsActive: true,
	}

	// Hash password
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	// Create user
	if err := a.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
