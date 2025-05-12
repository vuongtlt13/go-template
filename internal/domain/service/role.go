package service

import (
	"context"
	"errors"
	"fmt"
	coreRepo "yourapp/internal/core/repository"
	"yourapp/internal/domain/model"
	"yourapp/internal/domain/repository"

	"gorm.io/gorm"
)

// RoleService defines the interface for role operations
type RoleService interface {
	CreateRole(ctx context.Context, role *model.Role) error
	GetRoleByID(ctx context.Context, id uint64) (*model.Role, error)
	UpdateRole(ctx context.Context, role *model.Role) error
	DeleteRole(ctx context.Context, id uint64) error
	AssignPermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) error
	GetRolePermissions(ctx context.Context, roleID uint64) ([]model.Permission, error)
	ListRoles(ctx context.Context, page, pageSize int) ([]model.Role, int64, error)
}

// roleServiceImpl implements RoleService
type roleServiceImpl struct {
	db             *gorm.DB
	roleRepo       repository.RoleRepository
	permissionRepo repository.PermissionRepository
}

// NewRoleService creates a new role service with the given dependencies
func NewRoleService(db *gorm.DB, roleRepo repository.RoleRepository, permRepo repository.PermissionRepository) RoleService {
	return &roleServiceImpl{
		db:             db,
		roleRepo:       roleRepo,
		permissionRepo: permRepo,
	}
}

// CreateRole creates a new role
func (s *roleServiceImpl) CreateRole(ctx context.Context, role *model.Role) error {
	existing, err := s.roleRepo.FindByCode(ctx, role.Code, s.db)
	if err == nil && existing != nil {
		return ErrDuplicateRoleCode
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.roleRepo.Create(ctx, role, s.db)
}

// GetRoleByID gets a role by ID
func (s *roleServiceImpl) GetRoleByID(ctx context.Context, id uint64) (*model.Role, error) {
	return s.roleRepo.FindByID(ctx, id, nil, s.db)
}

// UpdateRole updates a role
func (s *roleServiceImpl) UpdateRole(ctx context.Context, role *model.Role) error {
	existing, err := s.roleRepo.FindByID(ctx, role.ID, nil, s.db)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.roleRepo.Update(ctx, role, s.db)
}

// DeleteRole deletes a role
func (s *roleServiceImpl) DeleteRole(ctx context.Context, id uint64) error {
	existing, err := s.roleRepo.FindByID(ctx, id, nil, s.db)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	return s.roleRepo.Delete(ctx, existing, s.db)
}

// AssignPermissions assigns permissions to a role
func (s *roleServiceImpl) AssignPermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) error {
	// Check if role exists
	role, err := s.roleRepo.FindByID(ctx, roleID, nil, s.db)
	if err != nil {
		return fmt.Errorf("error finding role: %w", err)
	}
	if role == nil {
		return ErrRoleNotFound
	}

	// Check if all permissions exist
	for _, permID := range permissionIDs {
		perm, err := s.permissionRepo.FindByID(ctx, permID, nil, s.db)
		if err != nil {
			return fmt.Errorf("error finding permission: %w", err)
		}
		if perm == nil {
			return ErrPermissionNotFound
		}
	}

	return s.roleRepo.AssignPermissions(ctx, roleID, permissionIDs, s.db)
}

// GetRolePermissions gets all permissions for a role
func (s *roleServiceImpl) GetRolePermissions(ctx context.Context, roleID uint64) ([]model.Permission, error) {
	role, err := s.roleRepo.FindByID(ctx, roleID, nil, s.db)
	if err != nil {
		return nil, fmt.Errorf("error finding role: %w", err)
	}

	var permissions []model.Permission
	if err := s.db.Model(role).Association("Permissions").Find(&permissions); err != nil {
		return nil, fmt.Errorf("error finding role permissions: %w", err)
	}

	return permissions, nil
}

// ListRoles lists roles with pagination
func (s *roleServiceImpl) ListRoles(ctx context.Context, page, pageSize int) ([]model.Role, int64, error) {
	offset := (page - 1) * pageSize
	opts := &coreRepo.QueryOptions{
		Limit:  pageSize,
		Offset: offset,
	}

	roles, err := s.roleRepo.All(ctx, opts, s.db)
	if err != nil {
		return nil, 0, err
	}

	// Convert []*model.Role to []model.Role
	result := make([]model.Role, len(roles))
	for i, role := range roles {
		result[i] = *role
	}

	// Get total count using repository
	total, err := s.roleRepo.Count(ctx, nil, s.db)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
