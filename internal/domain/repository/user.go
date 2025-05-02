package repository

import (
	"context"
	"fmt"
	"yourapp/internal/core/repository"
	"yourapp/internal/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	repository.BaseRepository[model.User]
	FindByEmail(ctx context.Context, email string, db *gorm.DB) (*model.User, error)
}

type userRepository struct {
}

func (u userRepository) First(ctx context.Context, opts *repository.QueryOptions, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) All(ctx context.Context, opts *repository.QueryOptions, db *gorm.DB) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByID(ctx context.Context, id any, opts *repository.QueryOptions, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindOrFail(ctx context.Context, id any, errMsg string, opts *repository.QueryOptions, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Create(ctx context.Context, obj *model.User, db *gorm.DB, flush bool) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(ctx context.Context, obj *model.User, db *gorm.DB, flush bool) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ctx context.Context, obj *model.User, db *gorm.DB, flush bool) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByEmail(ctx context.Context, email string, db *gorm.DB) (*model.User, error) {
	user := &model.User{}
	result := db.WithContext(ctx).First(user, "email = ?", email)
	if result.Error != nil {
		return nil, fmt.Errorf("error when finding user: %w", result.Error)
	}
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
