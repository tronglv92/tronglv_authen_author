package server

import (
	"github/tronglv_authen_author/helper/server/http/middleware"

	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func MustSetup(c Config) {
	Initialize()
	if !c.StatLog {
		stat.DisableLog()
	}
	if !c.LoadLog {
		load.DisableLog()
	}
}

func Initialize() {
	for _, v := range Providers() {
		v.Register()
	}
}

func NewHttpServer(c Config, h RestHandler, opts ...rest.RunOption) *rest.Server {
	MustSetup(c)
	srv := rest.MustNewServer(c.Http, opts...)
	srv.Use(middleware.NewRecoveryMiddleware(c.Env).Handle)

	h.Register(srv)
	return srv
}

func NewGrpcServer(c Config, h GrpcHandler, opts ...grpc.ServerOption) *zrpc.RpcServer {
	MustSetup(c)
	s := zrpc.MustNewServer(c.Grpc, func(grpcServer *grpc.Server) {
		h.Register(grpcServer)
		reflection.Register(grpcServer)
	})
	s.AddOptions(opts...)
	s.AddUnaryInterceptors(
	// interceptor.NewLogInterceptor().Unary(),
	// interceptor.NewTraceInterceptor(c.GetId(), c.GetGrpcName()).Unary(),
	)
	h.Interceptors(s)
	return s
}
