package response

import (
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/helper/auth"
)

func TokenAuthMapToResponse(data auth.AuthData) *authentication.TokenAuthResp {
	var userData *authentication.UserData
	var clientData *authentication.ClientData
	if data.GetClient() != nil {
		clientData = &authentication.ClientData{
			Id:     data.GetClient().GetId(),
			Uid:    data.GetClient().GetUid(),
			Name:   data.GetClient().GetName(),
			Scopes: data.GetClient().GetScopes(),
		}
	}
	if data.GetUser() != nil {
		userData = &authentication.UserData{
			Id:    data.GetUser().GetId(),
			Email: data.GetUser().GetEmail(),

			Name: data.GetUser().GetName(),

			Roles: data.GetUser().GetRoles(),
		}
	}
	return &authentication.TokenAuthResp{
		Client: clientData,
		User:   userData,
	}
}
