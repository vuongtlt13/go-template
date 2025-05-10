package repository

import (
	"context"

	"gorm.io/gorm"
)

// BaseRepository defines the common operations for all repositories
type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T, db *gorm.DB) error
	FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*T, error)
	Update(ctx context.Context, entity *T, db *gorm.DB) error
	Delete(ctx context.Context, entity *T, db *gorm.DB) error
	FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]T, error)
	Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error)
}

// BaseRepositoryImpl implements BaseRepository
type BaseRepositoryImpl[T any] struct {
	model T
}

// NewBaseRepository creates a new base repository
func NewBaseRepository[T any](model T) *BaseRepositoryImpl[T] {
	return &BaseRepositoryImpl[T]{
		model: model,
	}
}

// Create creates a new entity
func (r *BaseRepositoryImpl[T]) Create(ctx context.Context, entity *T, db *gorm.DB) error {
	return db.WithContext(ctx).Create(entity).Error
}

// FindByID finds an entity by ID
func (r *BaseRepositoryImpl[T]) FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*T, error) {
	var entity T
	query := db.WithContext(ctx).Model(&r.model)

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update updates an entity
func (r *BaseRepositoryImpl[T]) Update(ctx context.Context, entity *T, db *gorm.DB) error {
	return db.WithContext(ctx).Save(entity).Error
}

// Delete deletes an entity
func (r *BaseRepositoryImpl[T]) Delete(ctx context.Context, entity *T, db *gorm.DB) error {
	return db.WithContext(ctx).Delete(entity).Error
}

// FindAll finds all entities with pagination
func (r *BaseRepositoryImpl[T]) FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]T, error) {
	var entities []T
	query := db.WithContext(ctx).Model(&r.model)

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

// Count counts entities with optional conditions
func (r *BaseRepositoryImpl[T]) Count(ctx context.Context, conditions map[string]interface{}, db *gorm.DB) (int64, error) {
	var count int64
	query := db.WithContext(ctx).Model(&r.model)

	if conditions != nil {
		query = query.Where(conditions)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
