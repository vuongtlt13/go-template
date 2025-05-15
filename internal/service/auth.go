package service

import (
	"context"
	"errors"

	"yourapp/internal/model"
	"yourapp/internal/repository"
	authpkg "yourapp/pkg/auth"

	"gorm.io/gorm"
)

var (
	ErrUserInactive = errors.New("user is inactive")
	ErrUserExists   = errors.New("user already exists")
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
	jwtManager authpkg.JWTManagerInterface
}

// NewAuthService creates a new auth service with the given dependencies
func NewAuthService(db *gorm.DB, userRepo repository.UserRepository, jwtManager authpkg.JWTManagerInterface) AuthService {
	return &authService{
		db:         db,
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (a *authService) Login(ctx context.Context, cred Credential) (string, error) {
	user, err := a.userRepo.GetByEmail(ctx, cred.Email)
	if err != nil {
		return "", err
	}

	if !user.IsActive {
		return "", ErrUserInactive
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
	// Check if user already exists
	existingUser, err := a.userRepo.GetByEmail(ctx, cred.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return ErrUserExists
	}

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
