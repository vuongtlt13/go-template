package service

import (
	"context"
	"errors"
	"time"
	"yourapp/internal/model"
	"yourapp/internal/repository"

	"gorm.io/gorm"
)

// UserService defines the interface for user operations
type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id uint) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context) ([]*model.User, error)
	CleanupUnverifiedUsers(ctx context.Context, olderThanDays int) error
}

// userServiceImpl implements UserService
type userServiceImpl struct {
	db       *gorm.DB
	userRepo repository.UserRepository
}

// NewUserService creates a new user service with the given dependencies
func NewUserService(db *gorm.DB, userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		db:       db,
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *userServiceImpl) CreateUser(ctx context.Context, user *model.User) error {
	existing, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existing != nil {
		return gorm.ErrDuplicatedKey
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.userRepo.Create(ctx, user)
}

// GetUserByID gets a user by ID
func (s *userServiceImpl) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetUserByEmail gets a user by email
func (s *userServiceImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// UpdateUser updates a user
func (s *userServiceImpl) UpdateUser(ctx context.Context, user *model.User) error {
	existing, err := s.userRepo.GetByID(ctx, uint(user.ID))
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user
func (s *userServiceImpl) DeleteUser(ctx context.Context, id uint) error {
	existing, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.userRepo.Delete(ctx, id)
}

// ListUsers lists all users
func (s *userServiceImpl) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.List(ctx)
}

// CleanupUnverifiedUsers deletes unverified users older than the specified number of days
func (s *userServiceImpl) CleanupUnverifiedUsers(ctx context.Context, olderThanDays int) error {
	cutoffTime := time.Now().AddDate(0, 0, -olderThanDays)
	return s.userRepo.DeleteUnverifiedUsersCreatedBefore(ctx, cutoffTime)
}
