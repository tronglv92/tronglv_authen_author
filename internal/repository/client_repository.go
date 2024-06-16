package repository

import (
	"context"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/model"
	baseRepo "github/tronglv_authen_author/helper/sqldata/repository"
	"github/tronglv_authen_author/internal/types/entity"

	"gorm.io/gorm"
)

type clientRepo struct {
	baseRepo.BaseRepository[entity.Client]
}

type ClientRepository interface {
	First(ctx context.Context, opts ...baseRepo.QueryOpt) (*entity.Client, error)
	WithClientId(clientId string) baseRepo.QueryOpt
	FindWithPagination(ctx context.Context, limit int, page int, opts ...baseRepo.QueryOpt) ([]*entity.Client, *model.Pagination, error)
	WithOrder(sortBy string, sortOrder string, fields ...string) baseRepo.QueryOpt
	CreateWithReturn(ctx context.Context, entity *entity.Client) (*entity.Client, error)
	WithPreloads(relations ...string) baseRepo.QueryOpt
}

func NewClientRepository(db db.Database) ClientRepository {
	return &clientRepo{
		baseRepo.NewBaseRepository[entity.Client](db),
	}
}

func (r *clientRepo) WithClientId(clientId string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("client_id=?", clientId)
	}
}
