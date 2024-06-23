package config

import (
	"flag"
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

func Load(file *string) Config {
	flag.Parse()
	var c Config
	conf.MustLoad(*file, &c, conf.UseEnv())
	return c
}

type Config struct {
	Server   ServerConfig `json:"server,optional"`
	Cache    cache.Config `json:"cache,optional"`
	OAuth    OAuthConfig  `json:"oauth,optional"`
	Database db.DBConfig  `json:"database,optional"`
	JWT      JWTConfig    `json:"jwt,optional"`
}

func (c Config) ServiceName() string {
	return c.Server.Http.Name
}

type OAuthConfig struct {
	Debug                 bool   `json:"debug,default=false"`
	HashSecret            string `json:"hash-secret,optional"`
	PrivateKey            string `json:"private-key,optional"`
	PublicKey             string `json:"public-key,optional"`
	AccessTokenLifespan   int    `json:"access-token-lifespan,default=1"`
	RefreshTokenLifespan  int    `json:"refresh-token-lifespan,default=24"`
	AuthorizeCodeLifespan int    `json:"authorize-code-lifespan,default=1"`
}

type JWTConfig struct {
	HashSecret           string `json:"hash-secret,optional"`
	AccessTokenLifespan  int    `json:"access-token-lifespan,default=1"`
	RefreshTokenLifespan int    `json:"refresh-token-lifespan,default=24"`
}

type ServerConfig struct {
	Id      int                `json:",default=0,optional"`
	Env     string             `json:",default=production,optional"`
	StatLog bool               `json:"stat-log,default=false"`
	LoadLog bool               `json:"load-log,default=false"`
	Http    rest.RestConf      `json:"http,optional"`
	Grpc    zrpc.RpcServerConf `json:"grpc,optional"`
}
