package service

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/helper/define"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/internal/registry"
	rp "github/tronglv_authen_author/internal/repository"
	"github/tronglv_authen_author/internal/types/entity"

	"github.com/ory/fosite/token/jwt"
)

type AuthenticationService interface {
	Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error)
	ClientClaims(ctx context.Context, clientId string) (*jwt.JWTClaims, error)
}

type authSvcImpl struct {
	reg        *registry.ServiceContext
	clientRepo rp.ClientRepository
}

func NewAuthenticationService(reg *registry.ServiceContext) AuthenticationService {
	return &authSvcImpl{
		reg:        reg,
		clientRepo: reg.ClientRepo,
	}
}

func (s *authSvcImpl) Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error) {
	return nil, nil
}

func (s *authSvcImpl) ClientClaims(ctx context.Context, clientId string) (*jwt.JWTClaims, error) {
	client, err := s.clientRepo.First(
		ctx,
		s.clientRepo.WithClientId(clientId),
		s.clientRepo.WithPreloads("Roles", "Roles.Permissions"),
	)
	if err != nil {
		return nil, errors.From(err)
	}

	return &jwt.JWTClaims{
		Issuer:  define.ClientIssuer,
		Subject: client.UId,
		Scope:   client.GetScopes(),

		Extra: map[string]any{
			"id":    client.Id,
			"name":  client.Name,
			"roles": s.roles(client.Roles),
		},
	}, nil
}
func (s *authSvcImpl) roles(roleList []*entity.Role) map[string]string {
	var pers = make(map[int32]bool)
	var roles = make(map[string]string)
	for _, role := range roleList {
		for _, v := range role.Permissions {
			if _, ok := pers[v.Id]; ok {
				continue
			}
			pers[v.Id] = true
			k := fmt.Sprintf("%d", v.ServiceId)
			if _, ok := roles[k]; !ok {
				roles[k] = fmt.Sprintf("%d", v.Id)
				continue
			}
			roles[k] = fmt.Sprintf("%s,%d", roles[k], v.Id)
		}
	}
	return roles
}
