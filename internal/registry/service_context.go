package registry

import (
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"

	"github/tronglv_authen_author/internal/config"
	"github/tronglv_authen_author/internal/types/entity"

	"github.com/ory/fosite"
)

type ServiceContext struct {
	Config      config.Config
	Fosite      fosite.OAuth2Provider
	CacheClient cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	cacheClient := cache.New(c.Cache)
	sqlConn := db.Must(
		&c.Database,
		db.WithGormMigrate(entity.RegisterMigrate()),
		db.WithCache(cacheClient),
	)

	return &ServiceContext{
		Config:      c,
		Fosite:      NewFositeContext(c.OAuth, sqlConn, cacheClient),
		CacheClient: cacheClient,
	}
}
