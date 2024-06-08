package server

import (
	"context"
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
)

type AuthenticationServer struct {
	reg               *registry.ServiceContext
	authenticationSvc service.AuthenticationService
	authentication.UnimplementedAuthenticationServer
}

func NewAuthenticationServer(reg *registry.ServiceContext) *AuthenticationServer {
	return &AuthenticationServer{
		reg:               reg,
		authenticationSvc: service.NewAuthenticationService(reg),
	}
}

func (s *AuthenticationServer) Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error) {
	return s.authenticationSvc.Validate(ctx, req)
}
