package server

import (
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/api/authorization"
	"github/tronglv_authen_author/internal/registry"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	svc *registry.ServiceContext
}

func NewGrpcHandler(svc *registry.ServiceContext) GrpcHandler {
	return GrpcHandler{svc: svc}
}

func (h GrpcHandler) Interceptors(_ *zrpc.RpcServer) {
	//
}

func (h GrpcHandler) Register(svr *grpc.Server) {
	authentication.RegisterAuthenticationServer(svr, NewAuthenticationServer(h.svc))
	authorization.RegisterAuthorizationServer(svr, NewAuthorizationServer(h.svc))
}
