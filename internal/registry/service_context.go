package registry

import (
	"github/tronglv_authen_author/helper/auth"
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/tokenprovider"
	"github/tronglv_authen_author/helper/tokenprovider/jwt"

	mdh "github/tronglv_authen_author/helper/server/http/middleware"
	"github/tronglv_authen_author/internal/config"
	"github/tronglv_authen_author/internal/middleware"
	rp "github/tronglv_authen_author/internal/repository"
	"github/tronglv_authen_author/internal/types/entity"

	"github.com/ory/fosite"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                 config.Config
	Fosite                 fosite.OAuth2Provider
	CacheClient            cache.Cache
	ClientRepo             rp.ClientRepository
	PermissionRepo         rp.PermissionRepository
	ClientRoleRepo         rp.ClientRoleRepository
	UserRepo               rp.UserRepository
	JwtProvider            tokenprovider.Provider
	AuthMiddleware         rest.Middleware
	AuthInternalMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	cacheClient := cache.New(c.Cache)
	sqlConn := db.Must(
		&c.Database,
		db.WithGormMigrate(entity.RegisterMigrate()),
		db.WithCache(cacheClient),
	)
	userRepo := rp.NewUserRepository(sqlConn)

	return &ServiceContext{
		Config:         c,
		Fosite:         NewFositeContext(c.OAuth, sqlConn, cacheClient),
		CacheClient:    cacheClient,
		ClientRepo:     rp.NewClientRepository(sqlConn),
		PermissionRepo: rp.NewPermissionRepository(sqlConn),
		ClientRoleRepo: rp.NewClientRoleRepository(sqlConn),
		UserRepo:       userRepo,
		JwtProvider:    jwt.NewTokenJWTProvider(c.JWT),
		AuthMiddleware: mdh.NewAuthMiddleware(
			auth.WithPublicKey(c.OAuth.PublicKey)).Handle,
		AuthInternalMiddleware: middleware.NewAuthInternalMiddleware(userRepo, c.JWT.HashSecret).Handle,
	}
}
