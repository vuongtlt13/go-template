package service

import (
	"context"
	"errors"
	"yourapp/internal/model"
	"yourapp/internal/repository"
	"yourapp/pkg/database"

	"gorm.io/gorm"
)

// PermissionService defines the interface for permission operations
type PermissionService interface {
	CreatePermission(ctx context.Context, permission *model.Permission) error
	GetPermissionByID(ctx context.Context, id uint64) (*model.Permission, error)
	UpdatePermission(ctx context.Context, permission *model.Permission) error
	DeletePermission(ctx context.Context, id uint64) error
	ListPermissions(ctx context.Context, page, pageSize int) ([]model.Permission, int64, error)
}

type permissionServiceImpl struct {
	db             *gorm.DB
	permissionRepo repository.PermissionRepository
}

// NewPermissionService creates a new permission service with default dependencies
func NewPermissionService() PermissionService {
	return &permissionServiceImpl{
		db:             database.GetDatabase(),
		permissionRepo: repository.NewPermissionRepository(),
	}
}

// NewPermissionServiceWithMocks creates a new permission service with injected dependencies (for tests)
func NewPermissionServiceWithMocks(db *gorm.DB, repo repository.PermissionRepository) PermissionService {
	return &permissionServiceImpl{
		db:             db,
		permissionRepo: repo,
	}
}

// CreatePermission creates a new permission
func (s *permissionServiceImpl) CreatePermission(ctx context.Context, permission *model.Permission) error {
	// Check for duplicate code (assuming service and method are required fields)
	existing, err := s.permissionRepo.FindByServiceAndMethod(ctx, permission.Service, permission.Method, s.db)
	if err == nil && existing != nil {
		return ErrDuplicatePermissionCode
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.permissionRepo.Create(ctx, permission, s.db)
}

// GetPermissionByID gets a permission by ID
func (s *permissionServiceImpl) GetPermissionByID(ctx context.Context, id uint64) (*model.Permission, error) {
	return s.permissionRepo.FindByID(ctx, id, nil, s.db)
}

// UpdatePermission updates a permission
func (s *permissionServiceImpl) UpdatePermission(ctx context.Context, permission *model.Permission) error {
	existing, err := s.permissionRepo.FindByID(ctx, permission.ID, nil, s.db)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.permissionRepo.Update(ctx, permission, s.db)
}

// DeletePermission deletes a permission
func (s *permissionServiceImpl) DeletePermission(ctx context.Context, id uint64) error {
	existing, err := s.permissionRepo.FindByID(ctx, id, nil, s.db)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.permissionRepo.Delete(ctx, existing, s.db)
}

// ListPermissions lists permissions with pagination
func (s *permissionServiceImpl) ListPermissions(ctx context.Context, page, pageSize int) ([]model.Permission, int64, error) {
	offset := (page - 1) * pageSize
	permissions, err := s.permissionRepo.FindAll(ctx, offset, pageSize, nil, s.db)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.permissionRepo.Count(ctx, nil, s.db)
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}
