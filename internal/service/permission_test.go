package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"yourapp/internal/model"
)

func TestPermissionService_CreatePermission(t *testing.T) {
	// Setup
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewPermissionServiceWithMocks(mockDB, mockPermRepo)

	tests := []struct {
		name        string
		permission  *model.Permission
		mock        func()
		wantErr     bool
		wantErrType error
	}{
		{
			name: "successful creation",
			permission: &model.Permission{
				Name:    "Create User",
				Code:    "CREATE_USER",
				Service: "user",
				Method:  "create",
			},
			mock: func() {
				mockPermRepo.On("FindByServiceAndMethod", mock.Anything, "user", "create", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
				mockPermRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Permission"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "duplicate code",
			permission: &model.Permission{
				Name:    "Create User",
				Code:    "CREATE_USER",
				Service: "user",
				Method:  "create",
			},
			mock: func() {
				mockPermRepo.On("FindByServiceAndMethod", mock.Anything, "user", "create", mock.Anything).
					Return(&model.Permission{Code: "CREATE_USER"}, nil)
			},
			wantErr:     true,
			wantErrType: ErrDuplicatePermissionCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			err := service.CreatePermission(context.Background(), tt.permission)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErrType, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPermissionService_GetPermissionByID(t *testing.T) {
	// Setup
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewPermissionServiceWithMocks(mockDB, mockPermRepo)

	tests := []struct {
		name    string
		id      uint64
		mock    func()
		want    *model.Permission
		wantErr bool
	}{
		{
			name: "permission found",
			id:   1,
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Permission{ID: 1, Name: "Create User", Code: "CREATE_USER"}, nil)
			},
			want:    &model.Permission{ID: 1, Name: "Create User", Code: "CREATE_USER"},
			wantErr: false,
		},
		{
			name: "permission not found",
			id:   999,
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			permission, err := service.GetPermissionByID(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, permission)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.ID, permission.ID)
				assert.Equal(t, tt.want.Name, permission.Name)
				assert.Equal(t, tt.want.Code, permission.Code)
			}
		})
	}
}

func TestPermissionService_UpdatePermission(t *testing.T) {
	// Setup
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewPermissionServiceWithMocks(mockDB, mockPermRepo)

	tests := []struct {
		name       string
		permission *model.Permission
		mock       func()
		wantErr    bool
	}{
		{
			name: "successful update",
			permission: &model.Permission{
				ID:   1,
				Name: "Updated Create User",
				Code: "CREATE_USER",
			},
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Permission{ID: 1, Name: "Create User", Code: "CREATE_USER"}, nil)
				mockPermRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Permission"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "permission not found",
			permission: &model.Permission{
				ID:   999,
				Name: "Non-existent",
				Code: "NONEXISTENT",
			},
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.UpdatePermission(context.Background(), tt.permission)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPermissionService_DeletePermission(t *testing.T) {
	// Setup
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewPermissionServiceWithMocks(mockDB, mockPermRepo)

	tests := []struct {
		name    string
		id      uint64
		mock    func()
		wantErr bool
	}{
		{
			name: "successful deletion",
			id:   1,
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Permission{ID: 1}, nil)
				mockPermRepo.On("Delete", mock.Anything, mock.AnythingOfType("*model.Permission"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "permission not found",
			id:   999,
			mock: func() {
				mockPermRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := service.DeletePermission(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPermissionService_ListPermissions(t *testing.T) {
	// Setup
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewPermissionServiceWithMocks(mockDB, mockPermRepo)

	tests := []struct {
		name      string
		page      int
		pageSize  int
		mock      func()
		want      []model.Permission
		wantTotal int64
		wantErr   bool
	}{
		{
			name:     "successful listing",
			page:     1,
			pageSize: 10,
			mock: func() {
				mockPermRepo.On("FindAll", mock.Anything, 0, 10, mock.Anything, mock.Anything).
					Return([]model.Permission{
						{ID: 1, Name: "Create User", Code: "CREATE_USER"},
						{ID: 2, Name: "Update User", Code: "UPDATE_USER"},
					}, nil)
				mockPermRepo.On("Count", mock.Anything, mock.Anything, mock.Anything).
					Return(int64(2), nil)
			},
			want: []model.Permission{
				{ID: 1, Name: "Create User", Code: "CREATE_USER"},
				{ID: 2, Name: "Update User", Code: "UPDATE_USER"},
			},
			wantTotal: 2,
			wantErr:   false,
		},
		{
			name:     "empty list",
			page:     1,
			pageSize: 10,
			mock: func() {
				mockPermRepo.On("FindAll", mock.Anything, 0, 10, mock.Anything, mock.Anything).
					Return([]model.Permission{}, nil)
				mockPermRepo.On("Count", mock.Anything, mock.Anything, mock.Anything).
					Return(int64(0), nil)
			},
			want:      []model.Permission{},
			wantTotal: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			permissions, total, err := service.ListPermissions(context.Background(), tt.page, tt.pageSize)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, permissions)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTotal, total)
				assert.Equal(t, len(tt.want), len(permissions))
				if len(tt.want) > 0 {
					for i, perm := range permissions {
						assert.Equal(t, tt.want[i].ID, perm.ID)
						assert.Equal(t, tt.want[i].Name, perm.Name)
						assert.Equal(t, tt.want[i].Code, perm.Code)
					}
				}
			}
		})
	}
}
