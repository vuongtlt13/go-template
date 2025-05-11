package repository

import (
	"context"
	"time"
	"yourapp/internal/domain/model"
	"yourapp/pkg/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.User, error)
	DeleteUnverifiedUsersCreatedBefore(ctx context.Context, cutoffTime time.Time) error
}

type userRepository struct {
	db *gorm.DB
}

var (
	userRepoInstance *userRepository
)

// NewUserRepository returns the singleton instance of UserRepository
func NewUserRepository() UserRepository {
	if userRepoInstance == nil {
		userRepoInstance = &userRepository{
			db: database.GetDatabase(),
		}
	}
	return userRepoInstance
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

// List retrieves all users
func (r *userRepository) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUnverifiedUsersCreatedBefore deletes unverified users created before the cutoff time
func (r *userRepository) DeleteUnverifiedUsersCreatedBefore(ctx context.Context, cutoffTime time.Time) error {
	return r.db.WithContext(ctx).
		Where("is_verified = ? AND created_at < ?", false, cutoffTime).
		Delete(&model.User{}).Error
}
