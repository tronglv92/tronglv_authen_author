package request

import (
	"context"
	"github/tronglv_authen_author/internal/types/define"

	validation "github.com/itgelo/ozzo-validation/v4"
)

type ClientListReq struct {
	PaginationReq
	SortOrderReq
}

func (req *ClientListReq) SetDefault(ctx context.Context) error {
	if len(req.SortBy) == 0 {
		req.SortBy = define.SortByDefault
	}

	if len(req.SortOrder) == 0 {
		req.SortOrder = define.SortOrderDefault
	}
	return nil
}

type ClientSaveReq struct {
	Name         string   `json:"name"`
	Status       int      `json:"status"`
	Public       bool     `json:"public"`
	RedirectUrls []string `json:"redirect_urls"`
	Grants       []string `json:"grants"`
	RoleIds      []int32  `json:"role_ids,optional"`
	Scopes       []string `json:"scopes,optional"` // will remove in new version
}

func (req ClientSaveReq) Validate(ctx context.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.RedirectUrls, validation.Required),
	)
}
