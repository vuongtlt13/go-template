package service

import (
	"context"
	"time"
	"yourapp/internal/domain/model"
	"yourapp/pkg/auth"
	coreRepo "yourapp/pkg/core/repository"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUnverifiedUsersCreatedBefore(ctx context.Context, cutoffTime time.Time) error {
	args := m.Called(ctx, cutoffTime)
	return args.Error(0)
}

// MockJWTManager is a mock implementation of auth.JWTManagerInterface
type MockJWTManager struct {
	mock.Mock
}

// GenerateToken mocks the GenerateToken method
func (m *MockJWTManager) GenerateToken(userID uint64) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

// VerifyToken mocks the VerifyToken method
func (m *MockJWTManager) VerifyToken(tokenStr string) (uint64, error) {
	args := m.Called(tokenStr)
	return args.Get(0).(uint64), args.Error(1)
}

// ValidateToken mocks the ValidateToken method
func (m *MockJWTManager) ValidateToken(tokenStr string) (*auth.Claims, error) {
	args := m.Called(tokenStr)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.Claims), args.Error(1)
}

// MockRoleRepository is a mock implementation of RoleRepository
type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) Create(ctx context.Context, role *model.Role, db *gorm.DB) error {
	args := m.Called(ctx, role, db)
	return args.Error(0)
}

func (m *MockRoleRepository) FindByID(ctx context.Context, id any, opts *coreRepo.QueryOptions, db *gorm.DB) (*model.Role, error) {
	args := m.Called(ctx, id, opts, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

func (m *MockRoleRepository) Update(ctx context.Context, role *model.Role, db *gorm.DB) error {
	args := m.Called(ctx, role, db)
	return args.Error(0)
}

func (m *MockRoleRepository) Delete(ctx context.Context, role *model.Role, db *gorm.DB) error {
	args := m.Called(ctx, role, db)
	return args.Error(0)
}

func (m *MockRoleRepository) FindByCode(ctx context.Context, code string, db *gorm.DB) (*model.Role, error) {
	args := m.Called(ctx, code, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

func (m *MockRoleRepository) All(ctx context.Context, opts *coreRepo.QueryOptions, db *gorm.DB) ([]*model.Role, error) {
	args := m.Called(ctx, opts, db)
	return args.Get(0).([]*model.Role), args.Error(1)
}

func (m *MockRoleRepository) Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error) {
	args := m.Called(ctx, conditions, db)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRoleRepository) AssignPermissions(ctx context.Context, roleID uint64, permissionIDs []uint64, db *gorm.DB) error {
	args := m.Called(ctx, roleID, permissionIDs, db)
	return args.Error(0)
}

func (m *MockRoleRepository) GetRolePermissions(ctx context.Context, roleID uint64, db *gorm.DB) ([]model.Permission, error) {
	args := m.Called(ctx, roleID, db)
	return args.Get(0).([]model.Permission), args.Error(1)
}

func (m *MockRoleRepository) FindOrFail(ctx context.Context, id any, field string, opts *coreRepo.QueryOptions, db *gorm.DB) (*model.Role, error) {
	args := m.Called(ctx, id, field, opts, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

func (m *MockRoleRepository) First(ctx context.Context, opts *coreRepo.QueryOptions, db *gorm.DB) (*model.Role, error) {
	args := m.Called(ctx, opts, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

func (m *MockRoleRepository) WithTransaction(ctx context.Context, tx *gorm.DB, fn func(*gorm.DB) error) error {
	args := m.Called(ctx, tx, fn)
	return args.Error(0)
}

// MockPermissionRepository is a mock implementation of PermissionRepository
type MockPermissionRepository struct {
	mock.Mock
}

func (m *MockPermissionRepository) Create(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	args := m.Called(ctx, permission, db)
	return args.Error(0)
}

func (m *MockPermissionRepository) FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*model.Permission, error) {
	args := m.Called(ctx, id, preloads, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Permission), args.Error(1)
}

func (m *MockPermissionRepository) Update(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	args := m.Called(ctx, permission, db)
	return args.Error(0)
}

func (m *MockPermissionRepository) Delete(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	args := m.Called(ctx, permission, db)
	return args.Error(0)
}

func (m *MockPermissionRepository) FindByServiceAndMethod(ctx context.Context, service, method string, db *gorm.DB) (*model.Permission, error) {
	args := m.Called(ctx, service, method, db)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Permission), args.Error(1)
}

func (m *MockPermissionRepository) FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]model.Permission, error) {
	args := m.Called(ctx, offset, limit, preloads, db)
	return args.Get(0).([]model.Permission), args.Error(1)
}

func (m *MockPermissionRepository) Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error) {
	args := m.Called(ctx, conditions, db)
	return args.Get(0).(int64), args.Error(1)
}
