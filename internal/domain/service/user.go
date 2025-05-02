package service

import (
	"yourapp/internal/domain/repository"
)

// UserService defines the user service interface
type UserService interface {
}

// userServiceImpl implements the UserService interface
type userServiceImpl struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService() UserService {
	return &userServiceImpl{
		repo: repository.NewUserRepository(),
	}
}
