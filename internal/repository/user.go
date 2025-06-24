package repository

import (
	"context"
	"time"
	"yourapp/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, obj *model.User) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, obj *model.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.User, error)
	// Thêm các method đặc thù nếu muốn
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	DeleteUnverifiedUsersCreatedBefore(ctx context.Context, t time.Time) error
}

type userRepository struct {
	db *gorm.DB
}

var (
	userRepoInstance *userRepository
)

// NewUserRepository returns the singleton instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	if userRepoInstance == nil {
		userRepoInstance = &userRepository{
			db: db,
		}
	}
	return userRepoInstance
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, obj *model.User) error {
	return r.db.WithContext(ctx).Create(obj).Error
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var obj model.User
	err := r.db.WithContext(ctx).First(&obj, id).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, obj *model.User) error {
	return r.db.WithContext(ctx).Save(obj).Error
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

// List retrieves all users
func (r *userRepository) List(ctx context.Context) ([]*model.User, error) {
	var objs []*model.User
	err := r.db.WithContext(ctx).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var obj model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// DeleteUnverifiedUsersCreatedBefore deletes unverified users created before a given time
func (r *userRepository) DeleteUnverifiedUsersCreatedBefore(ctx context.Context, t time.Time) error {
	return r.db.WithContext(ctx).
		Where("is_active = ? AND created_at < ?", false, t).
		Delete(&model.User{}).Error
}
