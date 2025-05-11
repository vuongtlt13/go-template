package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"yourapp/internal/domain/model"
)

func TestAuthService_Login(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	service := NewAuthService().(*authService)
	service.userRepo = mockUserRepo

	// Create a test user with hashed password
	testUser := &model.User{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashed_password",
		IsActive: true,
	}

	tests := []struct {
		name    string
		cred    Credential
		mock    func()
		wantErr bool
	}{
		{
			name: "successful login",
			cred: Credential{
				Email:    "test@example.com",
				Password: "testpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").
					Return(testUser, nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			cred: Credential{
				Email:    "nonexistent@example.com",
				Password: "testpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "nonexistent@example.com").
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
		{
			name: "invalid password",
			cred: Credential{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").
					Return(testUser, nil)
			},
			wantErr: true,
		},
		{
			name: "inactive user",
			cred: Credential{
				Email:    "inactive@example.com",
				Password: "testpassword",
			},
			mock: func() {
				inactiveUser := *testUser
				inactiveUser.Email = "inactive@example.com"
				inactiveUser.IsActive = false
				mockUserRepo.On("GetByEmail", mock.Anything, "inactive@example.com").
					Return(&inactiveUser, nil)
			},
			wantErr: true,
		},
		{
			name: "repository error on GetByEmail",
			cred: Credential{
				Email:    "error@example.com",
				Password: "testpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "error@example.com").
					Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			token, err := service.Login(context.Background(), tt.cred)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func TestAuthService_Register(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	service := NewAuthService().(*authService)
	service.userRepo = mockUserRepo

	tests := []struct {
		name    string
		cred    Credential
		mock    func()
		wantErr bool
	}{
		{
			name: "successful registration",
			cred: Credential{
				Email:    "newuser@example.com",
				Password: "newpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "newuser@example.com").
					Return(nil, gorm.ErrRecordNotFound)
				mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "email already exists",
			cred: Credential{
				Email:    "existing@example.com",
				Password: "password",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "existing@example.com").
					Return(&model.User{Email: "existing@example.com"}, nil)
			},
			wantErr: true,
		},
		{
			name: "repository error on GetByEmail",
			cred: Credential{
				Email:    "error@example.com",
				Password: "password",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "error@example.com").
					Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name: "repository error on Create",
			cred: Credential{
				Email:    "newuser2@example.com",
				Password: "newpassword",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "newuser2@example.com").
					Return(nil, gorm.ErrRecordNotFound)
				mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).
					Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.Register(context.Background(), tt.cred)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
