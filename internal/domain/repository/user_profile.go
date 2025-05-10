package repository

import (
	"yourapp/internal/core/repository"
	"yourapp/internal/domain/model"
)

// UserProfileRepository defines the interface for user profile operations
type UserProfileRepository interface {
	repository.BaseRepository[model.UserProfile]
}

// userProfileRepository implements UserProfileRepository
type userProfileRepository struct {
	*repository.BaseRepositoryImpl[model.UserProfile]
}

// NewUserProfileRepository creates a new user profile repository
func NewUserProfileRepository() UserProfileRepository {
	return &userProfileRepository{
		BaseRepositoryImpl: repository.NewBaseRepository[model.UserProfile](model.UserProfile{}),
	}
}
