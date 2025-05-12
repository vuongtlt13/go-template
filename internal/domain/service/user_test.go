package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"yourapp/internal/domain/model"
)

func TestUserService_CreateUser(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		user    *model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "successful creation",
			user: &model.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").
					Return(nil, gorm.ErrRecordNotFound)
				mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "duplicate email",
			user: &model.User{
				Email:    "existing@example.com",
				Password: "password123",
			},
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "existing@example.com").
					Return(&model.User{Email: "existing@example.com"}, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.CreateUser(context.Background(), tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		id      uint
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "user found",
			id:   1,
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(1)).
					Return(&model.User{ID: 1, Email: "test@example.com"}, nil)
			},
			want:    &model.User{ID: 1, Email: "test@example.com"},
			wantErr: false,
		},
		{
			name: "user not found",
			id:   999,
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(999)).
					Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			user, err := service.GetUserByID(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.ID, user.ID)
				assert.Equal(t, tt.want.Email, user.Email)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		email   string
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name:  "user found",
			email: "test@example.com",
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").
					Return(&model.User{ID: 1, Email: "test@example.com"}, nil)
			},
			want:    &model.User{ID: 1, Email: "test@example.com"},
			wantErr: false,
		},
		{
			name:  "user not found",
			email: "nonexistent@example.com",
			mock: func() {
				mockUserRepo.On("GetByEmail", mock.Anything, "nonexistent@example.com").
					Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			user, err := service.GetUserByEmail(context.Background(), tt.email)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.ID, user.ID)
				assert.Equal(t, tt.want.Email, user.Email)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		user    *model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "successful update",
			user: &model.User{
				ID:    1,
				Email: "updated@example.com",
			},
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(1)).
					Return(&model.User{ID: 1}, nil)
				mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.User")).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			user: &model.User{
				ID:    999,
				Email: "nonexistent@example.com",
			},
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(999)).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.UpdateUser(context.Background(), tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		id      uint
		mock    func()
		wantErr bool
	}{
		{
			name: "successful deletion",
			id:   1,
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(1)).
					Return(&model.User{ID: 1}, nil)
				mockUserRepo.On("Delete", mock.Anything, uint(1)).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			id:   999,
			mock: func() {
				mockUserRepo.On("GetByID", mock.Anything, uint(999)).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.DeleteUser(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_ListUsers(t *testing.T) {
	// Setup
	mockUserRepo := new(MockUserRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewUserService(mockDB, mockUserRepo)

	tests := []struct {
		name    string
		mock    func()
		want    []*model.User
		wantErr bool
	}{
		{
			name: "successful listing",
			mock: func() {
				mockUserRepo.On("List", mock.Anything).
					Return([]*model.User{
						{ID: 1, Email: "user1@example.com"},
						{ID: 2, Email: "user2@example.com"},
					}, nil)
			},
			want: []*model.User{
				{ID: 1, Email: "user1@example.com"},
				{ID: 2, Email: "user2@example.com"},
			},
			wantErr: false,
		},
		{
			name: "empty list",
			mock: func() {
				mockUserRepo.On("List", mock.Anything).
					Return([]*model.User{}, nil)
			},
			want:    []*model.User{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			users, err := service.ListUsers(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.want), len(users))
				for i, user := range users {
					assert.Equal(t, tt.want[i].ID, user.ID)
					assert.Equal(t, tt.want[i].Email, user.Email)
				}
			}
		})
	}
}
