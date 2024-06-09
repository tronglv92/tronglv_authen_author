package response

import "github/tronglv_authen_author/internal/types/entity"

type PermissionResponse struct {
	Id        int32  `json:"id"`
	UId       string `json:"uid"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	GroupName string `json:"group_name"`
}

func PermissionMapToResponse(data *entity.Permission) *PermissionResponse {
	return &PermissionResponse{
		Id:        data.Id,
		UId:       data.UId,
		Name:      data.Name,
		Code:      data.Code,
		Path:      data.Path,
		Method:    data.Method,
		GroupName: data.GroupName,
	}
}

func PermissionMapToResponses(items []*entity.Permission) []*PermissionResponse {
	var results []*PermissionResponse
	for _, val := range items {
		results = append(results, PermissionMapToResponse(val))
	}
	return results
}
