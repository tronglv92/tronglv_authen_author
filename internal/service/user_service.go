package service

import (
	"context"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/tokenprovider"
	"github/tronglv_authen_author/helper/tokenprovider/jwt"
	"github/tronglv_authen_author/helper/util"
	"github/tronglv_authen_author/internal/registry"
	rp "github/tronglv_authen_author/internal/repository"
	"github/tronglv_authen_author/internal/types/entity"
	"github/tronglv_authen_author/internal/types/request"
	"github/tronglv_authen_author/internal/types/response"
)

type UserService interface {
	Register(ctx context.Context, input request.RegisterReq) (*response.RegisterResponse, error)
	Login(ctx context.Context, input request.LoginReq) (*response.LoginResponse, error)
	Profile(ctx context.Context) (*response.UserResponse, error)
}

type userSvcImpl struct {
	reg         *registry.ServiceContext
	userRepo    rp.UserRepository
	jwtProvider tokenprovider.Provider
}

func NewUserService(reg *registry.ServiceContext) UserService {
	return &userSvcImpl{
		reg:         reg,
		userRepo:    reg.UserRepo,
		jwtProvider: reg.JwtProvider,
	}
}

// 1 check email exits
// 2 generate salt
// 3 hash password
// 4 create user with hash password and salt
// 5 generate access token
// 6 generate refresh token
func (s *userSvcImpl) Register(ctx context.Context, input request.RegisterReq) (*response.RegisterResponse, error) {

	user, _ := s.userRepo.First(ctx, s.userRepo.WithEmail(input.Email))
	if user != nil {
		return nil, errors.InternalServerReason("Email has existed")
	}

	salt, err := util.GenSalt(50)
	if err != nil {
		return nil, err
	}

	password := util.Hash(input.Password + salt)

	userEntity := entity.User{
		Email:     input.Email,
		Password:  password,
		Salt:      salt,
		LastName:  input.Password,
		FirstName: input.FirstName,
		Phone:     input.Phone,
	}
	resp, err := s.userRepo.CreateWithReturn(ctx, &userEntity)
	if err != nil {
		return nil, err
	}

	payload := &jwt.TokenPayloadImp{
		UId: resp.UId,
	}

	accessToken, err := s.jwtProvider.Generate(payload, s.reg.Config.JWT.AccessTokenLifespan)
	if err != nil {
		return nil, errors.BadRequest(err)
	}

	refreshToken, err := s.jwtProvider.Generate(payload, s.reg.Config.JWT.RefreshTokenLifespan)
	if err != nil {
		return nil, errors.BadRequest(err)
	}

	return &response.RegisterResponse{
		User:         response.UserMapToResponse(resp),
		AccessToken:  response.TokenMapToResponse(accessToken),
		RefreshToken: response.TokenMapToResponse(refreshToken),
	}, nil
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1 Access token and refresh token
// 4. Return tokens
func (s *userSvcImpl) Login(ctx context.Context, input request.LoginReq) (*response.LoginResponse, error) {
	user, err := s.userRepo.First(ctx, s.userRepo.WithEmail(input.Email))
	if err != nil {
		return nil, errors.InternalServerReason("Email or password invalid!")
	}

	passHashed := util.Hash(input.Password + user.Salt)

	if user.Password != passHashed {
		return nil, errors.InternalServerReason("Email or password invalid!")
	}

	payload := &jwt.TokenPayloadImp{
		UId: user.UId,
	}

	accessToken, err := s.jwtProvider.Generate(payload, s.reg.Config.JWT.AccessTokenLifespan)
	if err != nil {
		return nil, errors.BadRequest(err)
	}

	refreshToken, err := s.jwtProvider.Generate(payload, s.reg.Config.JWT.RefreshTokenLifespan)
	if err != nil {
		return nil, errors.BadRequest(err)
	}
	return &response.LoginResponse{
		User:         response.UserMapToResponse(user),
		AccessToken:  response.TokenMapToResponse(accessToken),
		RefreshToken: response.TokenMapToResponse(refreshToken),
	}, nil
}

func (s *userSvcImpl) Profile(ctx context.Context) (*response.UserResponse, error) {
	return nil, nil
}
