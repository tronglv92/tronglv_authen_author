package auth

import (
	"slices"

	"github/tronglv_authen_author/helper/util/token"
)

func ParseUser(t, secretKey string) (UserData, error) {
	user, err := token.NewTokenParser(token.WithSecretKey(secretKey)).Parse(t)
	if err != nil {
		return nil, err
	}
	return &userData{
		Id:   int32(user.GetInt("sub")),
		Name: user.GetString("name"),
	}, nil
}

type UserData interface {
	GetId() int32
	GetEmail() string
	GetName() string
	GetRoles() []string
}
type userData struct {
	Id    int32    `json:"id"`
	Email string   `json:"email,omitempty"`
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"portal_roles,omitempty"`
}

func (u *userData) GetId() int32 {
	return u.Id
}

func (u *userData) GetEmail() string {
	return u.Email
}

func (u *userData) GetName() string {
	return u.Name
}

func (u *userData) GetRoles() []string {
	return u.Roles
}

func (u *userData) HasRole(role string) bool {
	return slices.Contains(u.Roles, role)
}
