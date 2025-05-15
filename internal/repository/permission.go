package repository

import (
	"context"
	"yourapp/internal/model"

	"gorm.io/gorm"
)

// PermissionRepository defines the interface for permission operations
type PermissionRepository interface {
	BaseRepository[model.Permission]
	FindByServiceAndMethod(ctx context.Context, service, method string, db *gorm.DB) (*model.Permission, error)
}

// permissionRepository implements PermissionRepository
type permissionRepository struct {
	BaseRepositoryImpl[model.Permission]
}

// NewPermissionRepository creates a new permission repository
func NewPermissionRepository() PermissionRepository {
	return &permissionRepository{
		BaseRepositoryImpl: *NewBaseRepository[model.Permission](model.Permission{}),
	}
}

// FindByServiceAndMethod finds a permission by service and method
func (r *permissionRepository) FindByServiceAndMethod(ctx context.Context, service, method string, db *gorm.DB) (*model.Permission, error) {
	var permission model.Permission
	if err := db.WithContext(ctx).Where("service = ? AND method = ?", service, method).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

// Create creates a new permission
func (r *permissionRepository) Create(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	return db.WithContext(ctx).Create(permission).Error
}

// FindByID finds a permission by ID
func (r *permissionRepository) FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*model.Permission, error) {
	var permission model.Permission
	query := db.WithContext(ctx).Model(&model.Permission{})

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&permission, id).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

// Update updates a permission
func (r *permissionRepository) Update(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	return db.WithContext(ctx).Save(permission).Error
}

// Delete deletes a permission
func (r *permissionRepository) Delete(ctx context.Context, permission *model.Permission, db *gorm.DB) error {
	return db.WithContext(ctx).Delete(permission).Error
}

// FindAll finds all permissions with pagination
func (r *permissionRepository) FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]model.Permission, error) {
	var permissions []model.Permission
	query := db.WithContext(ctx).Model(&model.Permission{})

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Offset(offset).Limit(limit).Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

// Count counts permissions with optional conditions
func (r *permissionRepository) Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error) {
	var count int64
	query := db.WithContext(ctx).Model(&model.Permission{})

	if conditions != nil {
		query = query.Where(conditions)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
