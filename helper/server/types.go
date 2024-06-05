package server

import (
	"github/tronglv_authen_author/helper/locale"
	"github/tronglv_authen_author/helper/server/core"

	"github.com/zeromicro/go-zero/rest"
)

type RestHandler interface {
	Register(svr *rest.Server)
}

type Config struct {
	Id      int           `json:",default=0,optional"`
	Env     string        `json:",default=production,optional"`
	StatLog bool          `json:"stat-log,default=false"`
	LoadLog bool          `json:"load-log,default=false"`
	Http    rest.RestConf `json:"http,optional"`
}

func Providers() []core.Service {
	return []core.Service{
		locale.NewLocalizer(),
	}
}
