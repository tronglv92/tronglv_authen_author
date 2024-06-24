package repository

import (
	"context"
	db "github/tronglv_authen_author/helper/database"
	baseRepo "github/tronglv_authen_author/helper/sqldata/repository"
	"github/tronglv_authen_author/internal/types/entity"

	"gorm.io/gorm"
)

type userRepo struct {
	baseRepo.BaseRepository[entity.User]
}

type UserRepository interface {
	First(ctx context.Context, opts ...baseRepo.QueryOpt) (*entity.User, error)
	WithEmail(email string) baseRepo.QueryOpt
	WithUID(uid string) baseRepo.QueryOpt
	CreateWithReturn(ctx context.Context, entity *entity.User) (*entity.User, error)
}

func NewUserRepository(db db.Database) UserRepository {
	return &userRepo{
		baseRepo.NewBaseRepository[entity.User](db),
	}
}

func (r *userRepo) WithEmail(email string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email=?", email)
	}
}

func (r *userRepo) WithUID(uid string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uid=?", uid)
	}
}
