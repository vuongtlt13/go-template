package repository

import (
	"context"
	"yourapp/internal/domain/model"
	"yourapp/pkg/core/repository"

	"gorm.io/gorm"
)

// RoleRepository defines the interface for role operations
type RoleRepository interface {
	repository.BaseRepository[model.Role]
	FindByCode(ctx context.Context, code string, db *gorm.DB) (*model.Role, error)
	AssignPermissions(ctx context.Context, roleID uint64, permissionIDs []uint64, db *gorm.DB) error
	GetRolePermissions(ctx context.Context, roleID uint64, db *gorm.DB) ([]model.Permission, error)
	Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error)
}

// roleRepository implements RoleRepository
type roleRepository struct {
	*repository.BaseRepositoryImpl[model.Role]
}

// FindByCode finds a role by its code
func (r *roleRepository) FindByCode(ctx context.Context, code string, db *gorm.DB) (*model.Role, error) {
	var role model.Role
	if err := db.WithContext(ctx).Where("code = ?", code).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// AssignPermissions assigns permissions to a role
func (r *roleRepository) AssignPermissions(ctx context.Context, roleID uint64, permissionIDs []uint64, db *gorm.DB) error {
	role := &model.Role{ID: roleID}
	if err := db.WithContext(ctx).First(role).Error; err != nil {
		return err
	}

	var permissions []model.Permission
	if err := db.WithContext(ctx).Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		return err
	}

	return db.WithContext(ctx).Model(role).Association("Permissions").Replace(permissions)
}

// GetRolePermissions gets all permissions for a role
func (r *roleRepository) GetRolePermissions(ctx context.Context, roleID uint64, db *gorm.DB) ([]model.Permission, error) {
	role := &model.Role{ID: roleID}
	if err := db.WithContext(ctx).First(role).Error; err != nil {
		return nil, err
	}

	var permissions []model.Permission
	if err := db.WithContext(ctx).Model(role).Association("Permissions").Find(&permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}

// Count counts the number of roles matching the given conditions
func (r *roleRepository) Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error) {
	var count int64
	query := db.WithContext(ctx).Model(&model.Role{})
	if conditions != nil {
		query = query.Where(conditions)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// NewRoleRepository creates a new role repository
func NewRoleRepository() RoleRepository {
	return &roleRepository{
		BaseRepositoryImpl: repository.NewBaseRepository[model.Role](model.Role{}),
	}
}
