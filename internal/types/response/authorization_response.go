package response

import (
	"github/tronglv_authen_author/api/authorization"
	"github/tronglv_authen_author/internal/types/entity"
)

func AuthorizationMapToResponse(item *entity.Permission) *authorization.PermissionItem {
	return &authorization.PermissionItem{
		Id:        item.Id,
		ServiceId: item.ServiceId,
		Code:      item.Code,
		Path:      item.Path,
		Method:    item.Method,
	}
}

func AuthorizationMapToResponses(items []*entity.Permission) []*authorization.PermissionItem {
	var results []*authorization.PermissionItem
	for _, val := range items {
		results = append(results, AuthorizationMapToResponse(val))
	}
	return results
}
