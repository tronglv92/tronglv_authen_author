package repository

import (
	db "github/tronglv_authen_author/helper/database"
	baseRepo "github/tronglv_authen_author/helper/sqldata/repository"
	"github/tronglv_authen_author/internal/types/entity"
)

type clientRepo struct {
	baseRepo.BaseRepository[entity.Client]
}

type ClientRepository interface {
}

func NewClientRepository(db db.Database) ClientRepository {
	return &clientRepo{
		baseRepo.NewBaseRepository[entity.Client](db),
	}
}
