package repository

import (
	"context"
	"fmt"
	util "github/tronglv_authen_author/helper"
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/logify"
	"github/tronglv_authen_author/helper/model"
	"math"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const DBContextKey = "DB"

type QueryOpt func(*gorm.DB) *gorm.DB
type PreloadOpt any

type baseRepository[T any] struct {
	db    *gorm.DB
	cache cache.Cache
}

type BaseRepository[T any] interface {
	GetDB(ctx context.Context, opts ...QueryOpt) *gorm.DB
	Transaction(fc func(tx *gorm.DB) error) error
	BeginTx(ctx context.Context) (*gorm.DB, context.Context)
	First(ctx context.Context, opts ...QueryOpt) (*T, error)
	FirstWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) (*T, error)
	Find(ctx context.Context, opts ...QueryOpt) ([]*T, error)
	FindWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) ([]*T, error)
	FindWithPagination(ctx context.Context, limit int, page int, opts ...QueryOpt) ([]*T, *model.Pagination, error)
	FindPaginationWithCache(ctx context.Context, key string, expiration time.Duration, limit int, page int, opts ...QueryOpt) ([]*T, *model.Pagination, error)
	Count(ctx context.Context, opts ...QueryOpt) (int64, error)
	CountWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) (int64, error)
	Create(ctx context.Context, entity *T) error
	CreateWithReturn(ctx context.Context, entity *T) (*T, error)
	Update(ctx context.Context, params any, opts ...QueryOpt) error
	UpdateWithReturn(ctx context.Context, params any, opts ...QueryOpt) (*T, error)
	UpdateWithCache(ctx context.Context, keys []string, params any, opts ...QueryOpt) error
	UpdateWithReturnCache(ctx context.Context, keys []string, params any, opts ...QueryOpt) (*T, error)
	Delete(ctx context.Context, opts ...QueryOpt) error
	DeleteWithCache(ctx context.Context, keys []string, opts ...QueryOpt) error
	Paginate(value interface{}, pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB
	WithPreload(relation string, opts ...PreloadOpt) QueryOpt
	WithPreloads(relations ...string) QueryOpt
	WithId(id uint32) QueryOpt
	WithUid(uid string) QueryOpt
	WithSoftDelete() QueryOpt
	WithCreatedDate(t time.Time) QueryOpt
	WithOrder(sortBy string, sortOrder string, fields ...string) QueryOpt
	SortAble() map[string]string
}

func NewBaseRepository[T any](db db.Database) BaseRepository[T] {
	return &baseRepository[T]{
		db:    db.GetGormClient(),
		cache: db.GetCacheClient(),
	}
}

func (r *baseRepository[T]) GetDB(ctx context.Context, opts ...QueryOpt) *gorm.DB {
	l := r.db.WithContext(ctx)
	if tx, ok := ctx.Value(DBContextKey).(*gorm.DB); ok {
		l = tx.WithContext(ctx)
	}
	for _, opt := range opts {
		l = opt(l)
	}
	return l
}

func (r *baseRepository[T]) WithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, DBContextKey, db)
}

func (r *baseRepository[T]) First(ctx context.Context, opts ...QueryOpt) (*T, error) {
	var result T
	if err := r.GetDB(ctx, opts...).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *baseRepository[T]) FirstWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) (*T, error) {
	var result T
	if err := r.cache.Get(key, &result); err == nil {
		return &result, nil
	}
	record, err := r.First(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if err = r.cache.SetWithExpire(key, record, expiration); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return record, nil
}

func (r *baseRepository[T]) Find(ctx context.Context, opts ...QueryOpt) ([]*T, error) {
	var result []*T
	if err := r.GetDB(ctx, opts...).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *baseRepository[T]) FindWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) ([]*T, error) {
	var results []*T
	if err := r.cache.Get(key, &results); err == nil {
		return results, nil
	}

	records, err := r.Find(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if err = r.cache.SetWithExpire(key, records, expiration); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return records, nil
}

func (r *baseRepository[T]) FindWithPagination(ctx context.Context, limit int, page int, opts ...QueryOpt) ([]*T, *model.Pagination, error) {
	var results []*T
	var pagination model.Pagination
	pagination.Limit = limit
	pagination.Page = page

	query := r.GetDB(ctx, opts...)
	err := query.Scopes(r.Paginate(results, &pagination, query)).
		Find(&results).
		Error
	if err != nil {
		return nil, nil, err
	}

	return results, &pagination, nil
}

func (r *baseRepository[T]) FindPaginationWithCache(ctx context.Context, key string, expiration time.Duration, limit int, page int, opts ...QueryOpt) ([]*T, *model.Pagination, error) {
	var results []*T
	var pagination model.Pagination
	if err := r.cache.Get(key, &results); err == nil {
		if err = r.cache.Get(fmt.Sprintf("%s:paging", key), &pagination); err == nil {
			return results, &pagination, nil
		}
	}

	records, paging, err := r.FindWithPagination(ctx, limit, page, opts...)
	if err != nil {
		return nil, nil, err
	}
	if err = r.cache.SetWithExpire(key, records, expiration); err == nil {
		if err = r.cache.SetWithExpire(fmt.Sprintf("%s:paging", key), paging, expiration); err != nil {
			logify.New().ErrorCtx(ctx, err)
		}
	}
	return records, paging, nil
}

func (r *baseRepository[T]) Count(ctx context.Context, opts ...QueryOpt) (int64, error) {
	var m T
	var total int64
	if err := r.GetDB(ctx, opts...).Model(&m).Find(&total).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *baseRepository[T]) CountWithCache(ctx context.Context, key string, expiration time.Duration, opts ...QueryOpt) (int64, error) {
	var total int64
	if err := r.cache.Get(key, &total); err == nil {
		return total, nil
	}

	count, err := r.Count(ctx, opts...)
	if err != nil {
		return 0, err
	}
	if err = r.cache.SetWithExpire(key, count, expiration); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return count, nil
}

func (r *baseRepository[T]) Create(ctx context.Context, model *T) error {
	if err := r.GetDB(ctx).Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *baseRepository[T]) CreateWithReturn(ctx context.Context, model *T) (*T, error) {
	if err := r.GetDB(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, params any, opts ...QueryOpt) error {
	var result T
	if err := r.GetDB(ctx, opts...).Model(&result).Updates(params).Error; err != nil {
		return err
	}
	return nil
}

func (r *baseRepository[T]) UpdateWithReturn(ctx context.Context, params any, opts ...QueryOpt) (*T, error) {
	var result T
	if err := r.GetDB(ctx, opts...).Clauses(clause.Returning{}).Model(&result).Updates(params).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *baseRepository[T]) UpdateWithCache(ctx context.Context, keys []string, params any, opts ...QueryOpt) error {
	if err := r.cache.DelCtx(ctx, keys...); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return r.Update(ctx, params, opts...)
}

func (r *baseRepository[T]) UpdateWithReturnCache(ctx context.Context, keys []string, params any, opts ...QueryOpt) (*T, error) {
	if err := r.cache.DelCtx(ctx, keys...); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return r.UpdateWithReturn(ctx, params, opts...)
}

func (r *baseRepository[T]) Delete(ctx context.Context, opts ...QueryOpt) error {
	var result T
	return r.GetDB(ctx, opts...).Delete(&result).Error
}

func (r *baseRepository[T]) DeleteWithCache(ctx context.Context, keys []string, opts ...QueryOpt) error {
	if err := r.cache.DelCtx(ctx, keys...); err != nil {
		logify.New().ErrorCtx(ctx, err)
	}
	return r.Delete(ctx, opts...)
}

func (r *baseRepository[T]) Transaction(fc func(tx *gorm.DB) error) error {
	return r.db.Transaction(fc)
}

func (r *baseRepository[T]) BeginTx(ctx context.Context) (*gorm.DB, context.Context) {
	tx := r.db.Begin()
	return tx, r.WithContext(ctx, tx)
}

func (r *baseRepository[T]) Paginate(value interface{}, pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	db.Model(value).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetLimit())))
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

func (r *baseRepository[T]) WithPreload(relation string, opts ...PreloadOpt) QueryOpt {
	return func(g *gorm.DB) *gorm.DB {
		var preloadOpts []interface{}
		for _, opt := range opts {
			preloadOpts = append(preloadOpts, opt)
		}
		return g.Preload(relation, preloadOpts...)
	}
}

func (r *baseRepository[T]) WithPreloads(relations ...string) QueryOpt {
	return func(g *gorm.DB) *gorm.DB {
		if len(relations) == 0 {
			return g
		}
		for _, relation := range relations {
			g = g.Preload(relation)
		}
		return g
	}
}

func (r *baseRepository[T]) WithUid(uid string) QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uid=?", uid)
	}
}

func (r *baseRepository[T]) WithId(id uint32) QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id=?", id)
	}
}

func (r *baseRepository[T]) WithOrder(sortBy string, sortOrder string, fields ...string) QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		m := r.SortAble()
		if len(fields) > 0 {
			for _, v := range fields {
				m[v] = v
			}
		}
		if val, ok := m[sortBy]; ok {
			return db.Order(fmt.Sprintf("%s %s", val, util.SortOrder(sortOrder)))
		}
		return db
	}
}

func (r *baseRepository[T]) WithCreatedDate(t time.Time) QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("created_at::date = ?", t.Format(time.DateOnly))
	}
}

func (r *baseRepository[T]) SortAble() map[string]string {
	return map[string]string{
		"id":         "id",
		"created_at": "created_at",
	}
}

func (r *baseRepository[T]) WithSoftDelete() QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_deleted = ?", true)
	}
}
