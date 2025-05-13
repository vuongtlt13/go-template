package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// QueryOptions represents options for database queries
type QueryOptions struct {
	Preloads []string
	OrderBy  string
	Limit    int
	Offset   int
	Where    map[string]interface{}
}

// BaseRepository defines the base repository interface
type BaseRepository[T any] interface {
	First(ctx context.Context, opts *QueryOptions, db *gorm.DB) (*T, error)
	All(ctx context.Context, opts *QueryOptions, db *gorm.DB) ([]*T, error)
	FindByID(ctx context.Context, id any, opts *QueryOptions, db *gorm.DB) (*T, error)
	FindOrFail(ctx context.Context, id any, errMsg string, opts *QueryOptions, db *gorm.DB) (*T, error)
	Create(ctx context.Context, obj *T, db *gorm.DB) error
	Update(ctx context.Context, obj *T, db *gorm.DB) error
	Delete(ctx context.Context, obj *T, db *gorm.DB) error
	WithTransaction(ctx context.Context, db *gorm.DB, fn func(tx *gorm.DB) error) error
}

// BaseRepositoryImpl provides a base implementation of BaseRepository
type BaseRepositoryImpl[T any] struct {
	model T
}

// NewBaseRepository creates a new base repository instance
func NewBaseRepository[T any](model T) *BaseRepositoryImpl[T] {
	return &BaseRepositoryImpl[T]{model: model}
}

// WithTransaction executes the given function within a transaction
func (r *BaseRepositoryImpl[T]) WithTransaction(ctx context.Context, db *gorm.DB, fn func(tx *gorm.DB) error) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// applyQueryOptions applies the query options to the given query
func (r *BaseRepositoryImpl[T]) applyQueryOptions(query *gorm.DB, opts *QueryOptions) *gorm.DB {
	if opts == nil {
		return query
	}

	// Apply preloads
	for _, preload := range opts.Preloads {
		query = query.Preload(preload)
	}

	// Apply ordering
	if opts.OrderBy != "" {
		query = query.Order(opts.OrderBy)
	}

	// Apply pagination
	if opts.Limit > 0 {
		query = query.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		query = query.Offset(opts.Offset)
	}

	// Apply where conditions
	if opts.Where != nil {
		query = query.Where(opts.Where)
	}

	return query
}

// First finds the first record matching the query
func (r *BaseRepositoryImpl[T]) First(ctx context.Context, opts *QueryOptions, db *gorm.DB) (*T, error) {
	var result T
	query := db.WithContext(ctx).Model(&r.model)

	query = r.applyQueryOptions(query, opts)

	if err := query.First(&result).Error; err != nil {
		return nil, fmt.Errorf("error finding first record: %w", err)
	}
	return &result, nil
}

// All retrieves all records
func (r *BaseRepositoryImpl[T]) All(ctx context.Context, opts *QueryOptions, db *gorm.DB) ([]*T, error) {
	var results []*T
	query := db.WithContext(ctx).Model(&r.model)

	query = r.applyQueryOptions(query, opts)

	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("error finding all records: %w", err)
	}
	return results, nil
}

// FindByID finds a record by its ID
func (r *BaseRepositoryImpl[T]) FindByID(ctx context.Context, id any, opts *QueryOptions, db *gorm.DB) (*T, error) {
	var result T
	query := db.WithContext(ctx).Model(&r.model)

	query = r.applyQueryOptions(query, opts)

	if err := query.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error finding record by ID: %w", err)
	}
	return &result, nil
}

// FindOrFail finds a record by ID or returns an error
func (r *BaseRepositoryImpl[T]) FindOrFail(ctx context.Context, id any, errMsg string, opts *QueryOptions, db *gorm.DB) (*T, error) {
	result, err := r.FindByID(ctx, id, opts, db)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}
	return result, nil
}

// Create creates a new record
func (r *BaseRepositoryImpl[T]) Create(ctx context.Context, obj *T, db *gorm.DB) error {
	if err := db.WithContext(ctx).Create(obj).Error; err != nil {
		return fmt.Errorf("error creating record: %w", err)
	}
	return nil
}

// Update updates an existing record
func (r *BaseRepositoryImpl[T]) Update(ctx context.Context, obj *T, db *gorm.DB) error {
	if err := db.WithContext(ctx).Save(obj).Error; err != nil {
		return fmt.Errorf("error updating record: %w", err)
	}
	return nil
}

// Delete deletes a record
func (r *BaseRepositoryImpl[T]) Delete(ctx context.Context, obj *T, db *gorm.DB) error {
	if err := db.WithContext(ctx).Delete(obj).Error; err != nil {
		return fmt.Errorf("error deleting record: %w", err)
	}
	return nil
}
