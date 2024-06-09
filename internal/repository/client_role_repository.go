package repository

import (
	"context"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/model"
	baseRepo "github/tronglv_authen_author/helper/sqldata/repository"
	"github/tronglv_authen_author/internal/types/entity"

	"gorm.io/gorm"
)

type clientRoleRepo struct {
	baseRepo.BaseRepository[entity.ClientRole]
}

type ClientRoleRepository interface {
	First(ctx context.Context, opts ...baseRepo.QueryOpt) (*entity.ClientRole, error)
	Find(ctx context.Context, opts ...baseRepo.QueryOpt) ([]*entity.ClientRole, error)
	FindWithPagination(ctx context.Context, limit int, page int, opts ...baseRepo.QueryOpt) ([]*entity.ClientRole, *model.Pagination, error)
	Bulk(ctx context.Context, entity []*entity.ClientRole) error
	CreateWithReturn(ctx context.Context, model *entity.ClientRole) (*entity.ClientRole, error)
	UpdateWithReturn(ctx context.Context, input any, opts ...baseRepo.QueryOpt) (*entity.ClientRole, error)
	Delete(ctx context.Context, opts ...baseRepo.QueryOpt) error
	WithClientId(clientId int32) baseRepo.QueryOpt
}

func NewClientRoleRepository(db db.Database) ClientRoleRepository {
	return &clientRoleRepo{
		baseRepo.NewBaseRepository[entity.ClientRole](db),
	}
}

func (r *clientRoleRepo) WithClientId(clientId int32) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("client_id=?", clientId)
	}
}

func (r *clientRoleRepo) Bulk(ctx context.Context, entity []*entity.ClientRole) error {
	if err := r.GetDB(ctx).Create(entity).Error; err != nil {
		return err
	}
	return nil
}
