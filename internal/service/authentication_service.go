package service

import (
	"context"
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/internal/registry"

	"github.com/ory/fosite/token/jwt"
)

type AuthenticationService interface {
	Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error)
	ClientClaims(ctx context.Context, clientId string) (*jwt.JWTClaims, error)
}

type authSvcImpl struct {
	reg *registry.ServiceContext
}

func NewAuthenticationService(reg *registry.ServiceContext) AuthenticationService {
	return &authSvcImpl{
		reg: reg,
	}
}

func (s *authSvcImpl) Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error) {
	return nil, nil
}

func (s *authSvcImpl) ClientClaims(ctx context.Context, clientId string) (*jwt.JWTClaims, error) {
	return nil, nil
}
