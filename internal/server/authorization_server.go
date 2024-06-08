package server

import (
	"context"

	"github/tronglv_authen_author/api/authorization"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
)

type AuthorizationServer struct {
	reg              *registry.ServiceContext
	authorizationSvc service.AuthorizationService
	authorization.UnimplementedAuthorizationServer
}

func NewAuthorizationServer(reg *registry.ServiceContext) *AuthorizationServer {
	return &AuthorizationServer{
		reg:              reg,
		authorizationSvc: service.NewAuthorizationService(reg),
	}
}

func (s *AuthorizationServer) Permissions(ctx context.Context, req *authorization.PermissionListReq) (*authorization.PermissionListResp, error) {
	return s.authorizationSvc.Permissions(ctx, req)
}

func (s *AuthorizationServer) Permission(ctx context.Context, req *authorization.PermissionReq) (*authorization.PermissionResp, error) {
	return s.authorizationSvc.Permission(ctx, req)
}

func (s *AuthorizationServer) PermissionByCode(ctx context.Context, req *authorization.PermissionByCodeReq) (*authorization.PermissionResp, error) {
	return s.authorizationSvc.PermissionByCode(ctx, req)
}
