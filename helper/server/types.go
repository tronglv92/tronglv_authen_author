package server

import (
	"github/tronglv_authen_author/helper/locale"
	"github/tronglv_authen_author/helper/server/core"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type RestHandler interface {
	Register(svr *rest.Server)
}

type GrpcHandler interface {
	Register(svr *grpc.Server)
	Interceptors(rpc *zrpc.RpcServer)
}



func Providers() []core.Service {
	return []core.Service{
		locale.NewLocalizer(),
	}
}
