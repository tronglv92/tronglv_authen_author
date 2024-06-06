package config

import (
	"flag"
	"github/tronglv_authen_author/helper/cache"
	"github/tronglv_authen_author/helper/server"
db "github/tronglv_authen_author/helper/database"
	"github.com/zeromicro/go-zero/core/conf"
)

func Load(file *string) Config {
	flag.Parse()
	var c Config
	conf.MustLoad(*file, &c, conf.UseEnv())
	return c
}

type Config struct {
	Server   server.Config `json:"server,optional"`
	Cache    cache.Config  `json:"cache,optional"`
	OAuth    OAuthConfig   `json:"oauth,optional"`
	Database db.DBConfig   `json:"database,optional"`
}

func (c Config) ServiceName() string {
	return c.Server.Http.Name
}

type OAuthConfig struct {
	Debug                 bool   `json:"debug,default=false"`
	HashSecret            string `json:"hash-secret,optional"`
	PrivateKey            string `json:"private-key,optional"`
	AccessTokenLifespan   int    `json:"access-token-lifespan,default=1"`
	RefreshTokenLifespan  int    `json:"refresh-token-lifespan,default=24"`
	AuthorizeCodeLifespan int    `json:"authorize-code-lifespan,default=1"`
}
