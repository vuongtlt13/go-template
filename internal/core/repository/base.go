package repository

import (
	"context"

	"gorm.io/gorm"
)

type QueryOptions struct {
	Preloads []string
	Where    map[string]interface{}
	OrderBys []string
	Limit    int
	Offset   int
}

type BaseRepository[T any] interface {
	First(ctx context.Context, opts *QueryOptions, db *gorm.DB) (*T, error)
	All(ctx context.Context, opts *QueryOptions, db *gorm.DB) ([]*T, error)
	FindByID(ctx context.Context, id any, opts *QueryOptions, db *gorm.DB) (*T, error)
	FindOrFail(ctx context.Context, id any, errMsg string, opts *QueryOptions, db *gorm.DB) (*T, error)
	Create(ctx context.Context, obj *T, db *gorm.DB, flush bool) error
	Update(ctx context.Context, obj *T, db *gorm.DB, flush bool) error
	Delete(ctx context.Context, obj *T, db *gorm.DB, flush bool) error
}
