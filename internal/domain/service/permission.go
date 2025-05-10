package service

import (
	"context"
	"yourapp/internal/domain/model"
	"yourapp/internal/domain/repository"
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

// permissionServiceImpl implements PermissionService
type permissionServiceImpl struct {
	db             *gorm.DB
	permissionRepo repository.PermissionRepository
}

var (
	permissionServiceInstance *permissionServiceImpl
)

// NewPermissionService creates a new permission service
func NewPermissionService() PermissionService {
	if permissionServiceInstance == nil {
		permissionServiceInstance = &permissionServiceImpl{
			db:             database.GetDatabase(),
			permissionRepo: repository.NewPermissionRepository(),
		}
	}
	return permissionServiceInstance
}

// CreatePermission creates a new permission
func (s *permissionServiceImpl) CreatePermission(ctx context.Context, permission *model.Permission) error {
	return s.permissionRepo.Create(ctx, permission, s.db)
}

// GetPermissionByID gets a permission by ID
func (s *permissionServiceImpl) GetPermissionByID(ctx context.Context, id uint64) (*model.Permission, error) {
	return s.permissionRepo.FindByID(ctx, id, nil, s.db)
}

// UpdatePermission updates a permission
func (s *permissionServiceImpl) UpdatePermission(ctx context.Context, permission *model.Permission) error {
	return s.permissionRepo.Update(ctx, permission, s.db)
}

// DeletePermission deletes a permission
func (s *permissionServiceImpl) DeletePermission(ctx context.Context, id uint64) error {
	permission := &model.Permission{ID: id}
	return s.permissionRepo.Delete(ctx, permission, s.db)
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
