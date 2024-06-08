package repository

import (
	"context"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/model"
	"github/tronglv_authen_author/internal/types/entity"

	"github/tronglv_authen_author/helper/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	baseRepo "github/tronglv_authen_author/helper/sqldata/repository"
)

type permissionRepo struct {
	baseRepo.BaseRepository[entity.Permission]
}

type PermissionRepository interface {
	First(ctx context.Context, opts ...baseRepo.QueryOpt) (*entity.Permission, error)
	Find(ctx context.Context, opts ...baseRepo.QueryOpt) ([]*entity.Permission, error)
	FindWithPagination(ctx context.Context, limit int, page int, opts ...baseRepo.QueryOpt) ([]*entity.Permission, *model.Pagination, error)
	Upsert(ctx context.Context, input []*entity.Permission) (int64, error)
	Create(ctx context.Context, entity *entity.Permission) error
	CreateWithReturn(ctx context.Context, entity *entity.Permission) (*entity.Permission, error)
	UpdateWithReturn(ctx context.Context, params any, opts ...baseRepo.QueryOpt) (*entity.Permission, error)
	Delete(ctx context.Context, opts ...baseRepo.QueryOpt) error
	WithPreloads(relations ...string) baseRepo.QueryOpt
	WithUid(uid string) baseRepo.QueryOpt
	WithFilter(filter map[string]any) baseRepo.QueryOpt
	WithServiceId(serviceId int32) baseRepo.QueryOpt
	WithCode(code string) baseRepo.QueryOpt
	WithMethod(method string) baseRepo.QueryOpt
	WithPath(path string) baseRepo.QueryOpt
	WithOrder(sortBy string, sortOrder string, fields ...string) baseRepo.QueryOpt
}

func NewPermissionRepository(db db.Database) PermissionRepository {
	return &permissionRepo{
		baseRepo.NewBaseRepository[entity.Permission](db),
	}
}

func (r *permissionRepo) WithServiceId(serviceId int32) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("service_id=?", serviceId)
	}
}

func (r *permissionRepo) WithCode(code string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("code=?", code)
	}
}

func (r *permissionRepo) WithPath(path string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("path=?", path)
	}
}

func (r *permissionRepo) WithMethod(method string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("method=?", method)
	}
}

func (r *permissionRepo) WithFilter(filter map[string]any) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		query := db
		if val, ok := util.FilterNotEmpty(filter, "keyword"); ok {
			query = query.Where("name ILIKE ?", "%"+val+"%")
		}
		if val, ok := util.FilterNotZero(filter, "service_id"); ok {
			query = query.Where("service_id = ?", val)
		}
		return query
	}
}

func (r *permissionRepo) Upsert(ctx context.Context, input []*entity.Permission) (int64, error) {
	result := r.GetDB(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "service_id"}, {Name: "path"}, {Name: "method"}},
		UpdateAll: true,
	}).Create(&input)

	return result.RowsAffected, result.Error
}
