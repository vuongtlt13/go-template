package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"yourapp/internal/domain/model"
)

func TestRoleService_CreateRole(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name        string
		role        *model.Role
		mock        func()
		wantErr     bool
		wantErrType error
	}{
		{
			name: "successful creation",
			role: &model.Role{
				Name: "Admin",
				Code: "ADMIN",
			},
			mock: func() {
				mockRoleRepo.On("FindByCode", mock.Anything, "ADMIN", mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
				mockRoleRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Role"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "duplicate code",
			role: &model.Role{
				Name: "Admin",
				Code: "ADMIN",
			},
			mock: func() {
				mockRoleRepo.On("FindByCode", mock.Anything, "ADMIN", mock.Anything).
					Return(&model.Role{Code: "ADMIN"}, nil)
			},
			wantErr:     true,
			wantErrType: ErrDuplicateRoleCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			err := service.CreateRole(context.Background(), tt.role)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErrType, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoleService_GetRoleByID(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name    string
		id      uint64
		mock    func()
		want    *model.Role
		wantErr bool
	}{
		{
			name: "role found",
			id:   1,
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1, Name: "Admin", Code: "ADMIN"}, nil)
			},
			want:    &model.Role{ID: 1, Name: "Admin", Code: "ADMIN"},
			wantErr: false,
		},
		{
			name: "role not found",
			id:   999,
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			role, err := service.GetRoleByID(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, role)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.ID, role.ID)
				assert.Equal(t, tt.want.Name, role.Name)
				assert.Equal(t, tt.want.Code, role.Code)
			}
		})
	}
}

func TestRoleService_UpdateRole(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name    string
		role    *model.Role
		mock    func()
		wantErr bool
	}{
		{
			name: "successful update",
			role: &model.Role{
				ID:   1,
				Name: "Updated Admin",
				Code: "ADMIN",
			},
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1, Name: "Admin", Code: "ADMIN"}, nil)
				mockRoleRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Role"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "role not found",
			role: &model.Role{
				ID:   999,
				Name: "Non-existent",
				Code: "NONEXISTENT",
			},
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			err := service.UpdateRole(context.Background(), tt.role)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoleService_DeleteRole(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

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
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1}, nil)
				mockRoleRepo.On("Delete", mock.Anything, mock.AnythingOfType("*model.Role"), mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "role not found",
			id:   999,
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			err := service.DeleteRole(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoleService_ListRoles(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name      string
		page      int
		pageSize  int
		mock      func()
		want      []model.Role
		wantTotal int64
		wantErr   bool
	}{
		{
			name:     "successful listing",
			page:     1,
			pageSize: 10,
			mock: func() {
				mockRoleRepo.On("All", mock.Anything, mock.Anything, mock.Anything).
					Return([]*model.Role{
						{ID: 1, Name: "Admin", Code: "ADMIN"},
						{ID: 2, Name: "User", Code: "USER"},
					}, nil)
				mockRoleRepo.On("Count", mock.Anything, mock.Anything, mock.Anything).
					Return(int64(2), nil)
			},
			want: []model.Role{
				{ID: 1, Name: "Admin", Code: "ADMIN"},
				{ID: 2, Name: "User", Code: "USER"},
			},
			wantTotal: 2,
			wantErr:   false,
		},
		{
			name:     "empty list",
			page:     1,
			pageSize: 10,
			mock: func() {
				mockRoleRepo.On("All", mock.Anything, mock.Anything, mock.Anything).
					Return([]*model.Role{}, nil)
				mockRoleRepo.On("Count", mock.Anything, mock.Anything, mock.Anything).
					Return(int64(0), nil)
			},
			want:      []model.Role{},
			wantTotal: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			roles, total, err := service.ListRoles(context.Background(), tt.page, tt.pageSize)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, roles)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTotal, total)
				assert.Equal(t, len(tt.want), len(roles))
				for i, role := range roles {
					assert.Equal(t, tt.want[i].ID, role.ID)
					assert.Equal(t, tt.want[i].Name, role.Name)
					assert.Equal(t, tt.want[i].Code, role.Code)
				}
			}
		})
	}
}

func TestRoleService_AssignPermissions(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name          string
		roleID        uint64
		permissionIDs []uint64
		mock          func()
		wantErr       bool
		wantErrType   error
	}{
		{
			name:          "successful assignment",
			roleID:        1,
			permissionIDs: []uint64{1, 2},
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1}, nil)
				mockPermRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Permission{ID: 1}, nil)
				mockPermRepo.On("FindByID", mock.Anything, uint64(2), mock.Anything, mock.Anything).
					Return(&model.Permission{ID: 2}, nil)
				mockRoleRepo.On("AssignPermissions", mock.Anything, uint64(1), []uint64{1, 2}, mock.Anything).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name:          "role not found",
			roleID:        999,
			permissionIDs: []uint64{1, 2},
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr:     true,
			wantErrType: ErrRoleNotFound,
		},
		{
			name:          "permission not found",
			roleID:        1,
			permissionIDs: []uint64{999},
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1}, nil)
				mockPermRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr:     true,
			wantErrType: ErrPermissionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			err := service.AssignPermissions(context.Background(), tt.roleID, tt.permissionIDs)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErrType, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoleService_GetRolePermissions(t *testing.T) {
	// Setup
	mockRoleRepo := new(MockRoleRepository)
	mockPermRepo := new(MockPermissionRepository)
	var mockDB *gorm.DB // nil is fine for tests since we mock all repo methods
	service := NewRoleService(mockDB, mockRoleRepo, mockPermRepo)

	tests := []struct {
		name    string
		roleID  uint64
		mock    func()
		want    []model.Permission
		wantErr bool
	}{
		{
			name:   "successful retrieval",
			roleID: 1,
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(1), mock.Anything, mock.Anything).
					Return(&model.Role{ID: 1}, nil)
				mockRoleRepo.On("GetRolePermissions", mock.Anything, uint64(1), mock.Anything).
					Return([]model.Permission{
						{ID: 1, Name: "read", Code: "READ"},
						{ID: 2, Name: "write", Code: "WRITE"},
					}, nil)
			},
			want: []model.Permission{
				{ID: 1, Name: "read", Code: "READ"},
				{ID: 2, Name: "write", Code: "WRITE"},
			},
			wantErr: false,
		},
		{
			name:   "role not found",
			roleID: 999,
			mock: func() {
				mockRoleRepo.On("FindByID", mock.Anything, uint64(999), mock.Anything, mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRoleRepo.ExpectedCalls = nil // reset mock
			mockPermRepo.ExpectedCalls = nil // reset mock
			tt.mock()
			permissions, err := service.GetRolePermissions(context.Background(), tt.roleID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, permissions)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.want), len(permissions))
				for i, perm := range permissions {
					assert.Equal(t, tt.want[i].ID, perm.ID)
					assert.Equal(t, tt.want[i].Name, perm.Name)
					assert.Equal(t, tt.want[i].Code, perm.Code)
				}
			}
		})
	}
}
