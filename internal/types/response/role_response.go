package response

import (
	"github/tronglv_authen_author/internal/types/entity"
	"time"
)

type RoleResponse struct {
	Id          int32                 `json:"id"`
	UId         string                `json:"uid"`
	Name        string                `json:"name"`
	Code        string                `json:"code"`
	CreatedBy   string                `json:"created_by"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedBy   string                `json:"updated_by"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Permissions []*PermissionResponse `json:"permissions"`
}

func RoleMapToResponse(data *entity.Role) *RoleResponse {
	return &RoleResponse{
		Id:          data.Id,
		UId:         data.UId,
		Name:        data.Name,
		Code:        data.Code,
		CreatedBy:   data.CreatedBy,
		CreatedAt:   data.CreatedAt,
		UpdatedBy:   data.UpdatedBy,
		UpdatedAt:   data.UpdatedAt,
		Permissions: PermissionMapToResponses(data.Permissions),
	}
}

func RoleMapToResponses(items []*entity.Role) []*RoleResponse {
	var results []*RoleResponse
	for _, val := range items {
		results = append(results, RoleMapToResponse(val))
	}

	// if results empty return [] instead null
	if len(results) == 0 {
		results = []*RoleResponse{}
	}

	return results
}
