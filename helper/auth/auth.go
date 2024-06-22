package auth

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/api/authentication"

	"github/tronglv_authen_author/helper/client"
	"github/tronglv_authen_author/helper/client/identity"
	"github/tronglv_authen_author/helper/define"
	"github/tronglv_authen_author/helper/util/token"
)

type (
	Validate interface {
		Validate(ctx context.Context, tkn string) (AuthData, error)
		RemoteValidate(ctx context.Context, tkn string) (AuthData, error)
		LocalValidate(ctx context.Context, tkn string) (AuthData, error)
		AssignToContext(ctx context.Context, u AuthData) context.Context
	}

	AuthOption struct {
		idtClient  *identity.Client
		gatewayUrl string
		publicKey  string
		secretKey  string
	}

	Option func(s *AuthOption)
)

func WithSecretKey(key string) Option {
	return func(m *AuthOption) {
		m.secretKey = key
	}
}

func WithPublicKey(key string) Option {
	return func(m *AuthOption) {
		m.publicKey = key
	}
}

func WithGatewayUrl(url string) Option {
	return func(m *AuthOption) {
		m.gatewayUrl = url
	}
}

func WithClient(idt *identity.Client) Option {
	return func(m *AuthOption) {
		m.idtClient = idt
	}
}

type authSvc struct {
	*AuthOption
}

func New(opts ...Option) Validate {
	r := &authSvc{
		AuthOption: &AuthOption{},
	}
	for _, opt := range opts {
		opt(r.AuthOption)
	}
	return r
}

func (s *authSvc) Validate(ctx context.Context, tkn string) (AuthData, error) {
	// if s.idtClient != nil {
	// 	return s.RemoteValidate(ctx, tkn)
	// }
	return s.LocalValidate(ctx, tkn)
}

func (s *authSvc) RemoteValidate(ctx context.Context, tkn string) (AuthData, error) {
	resp, err := s.idtClient.Validate(
		client.WithAuthToken(ctx, tkn),
		&authentication.TokenAuthReq{},
	)
	if err != nil {
		return nil, err
	}
	return NewAuthData(tkn, resp.Client, resp.User), nil
}

func (s *authSvc) LocalValidate(ctx context.Context, tkn string) (AuthData, error) {
	claims, err := token.NewTokenParser().ParseUnverified(tkn)
	if err != nil {
		return nil, err
	}

	fmt.Println("claims: ", claims)

	if claims.GetString("iss") == define.ClientIssuer {
		fmt.Println("vao trong nay 1")
		clt, err := ParseClient(tkn, s.publicKey)
		if err != nil {
			return nil, err
		}
		return NewAuthData(tkn, clt, nil), nil
	}

	if len(s.secretKey) > 0 {
		fmt.Println("vao trong nay 2")
		ult, err := ParseUser(tkn, s.secretKey)
		if err != nil {
			return nil, err
		}
		return NewAuthData(tkn, nil, ult), nil
	}
	fmt.Println("vao trong nay tkn: ", tkn)
	// rst := NewUserService(util.ServiceName(ctx), s.gatewayUrl)
	// user, err := rst.Auth(ctx, tkn)
	// if err != nil {
	// 	return nil, err
	// }
	// return NewAuthData(tkn, nil, user), nil
	return nil, nil
}

func (s *authSvc) AssignToContext(ctx context.Context, u AuthData) context.Context {
	return SetAuthDataToContext(ctx, u)
}
