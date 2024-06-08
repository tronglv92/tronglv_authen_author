package service

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/api/authorization"

	"github/tronglv_authen_author/internal/registry"
	rp "github/tronglv_authen_author/internal/repository"
)

type AuthorizationService interface {
	Permissions(ctx context.Context, req *authorization.PermissionListReq) (*authorization.PermissionListResp, error)
	Permission(ctx context.Context, req *authorization.PermissionReq) (*authorization.PermissionResp, error)
	PermissionByCode(ctx context.Context, req *authorization.PermissionByCodeReq) (*authorization.PermissionResp, error)
}

type authorizationSvcImpl struct {
	reg            *registry.ServiceContext
	clientRepo     rp.ClientRepository
	permissionRepo rp.PermissionRepository
}

func NewAuthorizationService(reg *registry.ServiceContext) AuthorizationService {
	return &authorizationSvcImpl{
		reg:            reg,
		clientRepo:     reg.ClientRepo,
		permissionRepo: reg.PermissionRepo,
	}
}

func (s *authorizationSvcImpl) Permissions(ctx context.Context, req *authorization.PermissionListReq) (*authorization.PermissionListResp, error) {

	fmt.Println("AuthorizationService Permissions")
	return nil, nil
}

func (s *authorizationSvcImpl) Permission(ctx context.Context, req *authorization.PermissionReq) (*authorization.PermissionResp, error) {
	fmt.Println("AuthorizationService Permission")
	return nil, nil
}

func (s *authorizationSvcImpl) PermissionByCode(ctx context.Context, req *authorization.PermissionByCodeReq) (*authorization.PermissionResp, error) {
	fmt.Println("AuthorizationService PermissionByCode")
	return nil, nil
}
