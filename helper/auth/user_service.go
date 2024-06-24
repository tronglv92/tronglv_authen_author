package auth

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/helper/httpc"
)

// xem xoa class nay
type UserResponse struct {
	Result    string   `json:"result,omitempty"`
	Data      userData `json:"data"`
	Error     string   `json:"error"`
	ErrorCode string   `json:"errorCode,omitempty"`
}

type UserService interface {
	Auth(ctx context.Context, token string) (UserData, error)
}

type userSvc struct {
	restClient httpc.Service
	gatewayUrl string
}

func NewUserService(serviceName, gatewayUrl string) UserService {
	return &userSvc{
		gatewayUrl: gatewayUrl,
		restClient: httpc.New(serviceName),
	}
}

func (r *userSvc) Auth(ctx context.Context, token string) (UserData, error) {
	resp, err := r.restClient.Get(ctx,
		fmt.Sprintf("%s/users/me", r.gatewayUrl),
		httpc.WithAuthToken(token),
	)
	if err != nil {
		return nil, err
	}

	var authResp UserResponse
	if err = httpc.ParseJsonBody(resp, &authResp); err != nil {
		return nil, err
	}
	return &authResp.Data, nil
}
