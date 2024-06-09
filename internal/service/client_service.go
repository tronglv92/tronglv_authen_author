package service

import (
	"context"
	"github/tronglv_authen_author/helper/define"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/model"
	"github/tronglv_authen_author/helper/util"
	"github/tronglv_authen_author/internal/registry"
	rp "github/tronglv_authen_author/internal/repository"
	df "github/tronglv_authen_author/internal/types/define"
	"github/tronglv_authen_author/internal/types/entity"
	"github/tronglv_authen_author/internal/types/request"
	"github/tronglv_authen_author/internal/types/response"

	"golang.org/x/crypto/bcrypt"
)

const (
	ClientCacheKey        = "clients"
	ClientIdMaxLength     = 36
	ClientSecretMaxLength = 36
)

var ResponseTypes = map[string][]string{
	df.GrantAuthorizationCode: {"token", "code"},
	df.GrantClientCredential:  {"token"},
}

type ClientService interface {
	List(ctx context.Context, input request.ClientListReq) ([]*response.ClientResponse, *model.Pagination, error)
	Create(ctx context.Context, input request.ClientSaveReq) (*response.ClientResponse, error)
}

type clientSvcImpl struct {
	reg            *registry.ServiceContext
	clientRepo     rp.ClientRepository
	clientRoleRepo rp.ClientRoleRepository
}

func NewClientService(reg *registry.ServiceContext) ClientService {
	return &clientSvcImpl{
		reg:            reg,
		clientRepo:     reg.ClientRepo,
		clientRoleRepo: reg.ClientRoleRepo,
	}
}

func (s *clientSvcImpl) List(ctx context.Context, input request.ClientListReq) ([]*response.ClientResponse, *model.Pagination, error) {
	results, pagination, err := s.clientRepo.FindWithPagination(
		ctx,
		input.Limit,
		input.Page,
		s.clientRepo.WithOrder(input.SortBy, input.SortOrder),
	)
	if err != nil {
		return nil, nil, errors.From(err)
	}

	if pagination.TotalRecords == 0 {
		return nil, nil, errors.DataNotFound()
	}

	return response.ClientMapToResponses(results), pagination, nil
}

func (s *clientSvcImpl) Create(ctx context.Context, input request.ClientSaveReq) (*response.ClientResponse, error) {
	clientId, err := util.RandomString(ClientIdMaxLength)
	if err != nil {
		return nil, errors.From(err)
	}

	clientSecret, err := util.RandomString(ClientSecretMaxLength)
	if err != nil {
		return nil, errors.From(err)
	}

	m := entity.Client{
		Name:             input.Name,
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		ClientSecretHash: s.hashVal(clientSecret),
		Status:           input.Status,
		Public:           &input.Public,
		Scopes:           input.Scopes,
		Grants:           input.Grants,
		Audiences:        []string{},
		RedirectUrls:     input.RedirectUrls,
		ResponseTypes:    s.responseTypes(input.Grants),
	}
	result, err := s.clientRepo.CreateWithReturn(ctx, &m)
	if err != nil {
		return nil, errors.From(err)
	}

	if err = s.roles(ctx, result.Id, input.RoleIds); err != nil {
		return nil, err
	}
	return response.ClientMapToResponse(result), nil
}

func (s *clientSvcImpl) hashVal(val string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
	if err != nil {
		return define.EmptyString
	}
	return string(hash)
}

func (s *clientSvcImpl) responseTypes(grants []string) []string {
	m := make(map[string]string)
	for _, val := range grants {
		if ts, ok := ResponseTypes[val]; ok {
			for _, v := range ts {
				m[v] = v
			}
		}
	}
	v := make([]string, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}

func (s *clientSvcImpl) roles(ctx context.Context, clientId int32, ids []int32) error {
	if len(ids) == 0 {
		return nil
	}

	var models []*entity.ClientRole
	for _, id := range ids {
		models = append(models, &entity.ClientRole{
			ClientId: clientId,
			RoleId:   id,
		})
	}
	if err := s.clientRoleRepo.Bulk(ctx, models); err != nil {
		return errors.From(err)
	}
	return nil
}
