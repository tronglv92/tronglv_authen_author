package response

import (
	"github/tronglv_authen_author/internal/types/entity"

	"github.com/lib/pq"
)

type ClientResponse struct {
	Id           int32           `json:"id"`
	UId          string          `json:"uid"`
	Name         string          `json:"name"`
	ClientId     string          `json:"client_id"`
	ClientSecret string          `json:"client_secret"`
	Status       int             `json:"status"`
	Public       bool            `json:"public"`
	Scopes       pq.StringArray  `json:"scopes"`
	Grants       pq.StringArray  `json:"grants"`
	Audiences    pq.StringArray  `json:"audiences"`
	RedirectUrls pq.StringArray  `json:"redirect_urls"`
	Roles        []*RoleResponse `json:"roles"`
}

func ClientMapToResponse(data *entity.Client) *ClientResponse {

	return &ClientResponse{
		Id:           data.Id,
		UId:          data.UId,
		Name:         data.Name,
		ClientId:     data.ClientId,
		ClientSecret: data.ClientSecret,
		Status:       data.Status,
		Public:       *data.Public,
		Scopes:       data.Scopes,
		Grants:       data.Grants,
		Audiences:    data.Audiences,
		RedirectUrls: data.RedirectUrls,
		Roles:        RoleMapToResponses(data.Roles),
	}
}

func ClientMapToResponses(items []*entity.Client) []*ClientResponse {
	var results []*ClientResponse
	for _, val := range items {
		results = append(results, ClientMapToResponse(val))
	}
	return results
}
