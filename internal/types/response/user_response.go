package response

import (
	"github/tronglv_authen_author/internal/types/entity"
)

type UserResponse struct {
	Id        int32  `json:"id"`
	UId       string `json:"uid"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
}

func UserMapToResponse(data *entity.User) *UserResponse {
	return &UserResponse{
		Id:        data.Id,
		UId:       data.UId,
		Email:     data.Email,
		LastName:  data.LastName,
		FirstName: data.FirstName,
		Phone:     data.Phone,
	}
}

type RegisterResponse struct {
	AccessToken  *TokenResponse `json:"access_token"`
	RefreshToken *TokenResponse `json:"refresh_token"`
	User         *UserResponse  `json:"user"`
}

type LoginResponse struct {
	AccessToken  *TokenResponse `json:"access_token"`
	RefreshToken *TokenResponse `json:"refresh_token"`
	User         *UserResponse  `json:"user"`
}
