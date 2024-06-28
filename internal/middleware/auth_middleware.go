package middleware

import (
	"fmt"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/helper/util"
	"github/tronglv_authen_author/helper/util/token"
	"github/tronglv_authen_author/internal/repository"

	"net/http"
)

type AuthInternalMiddleware struct {
	userRepo  repository.UserRepository
	secretKey string
}

func NewAuthInternalMiddleware(userRepo repository.UserRepository, secretKey string) *AuthInternalMiddleware {
	return &AuthInternalMiddleware{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
}

func (s *AuthInternalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := util.GetTokenFromHeader(r)
		if len(tkn) == 0 {
			response.Error(r.Context(), w, errors.BadRequest(
				fmt.Errorf("%s", "Header token is missing"),
			))
			return
		}

		// payload, err := jwtProvider.Validate(tkn)
		// if err != nil {
		// 	panic(common.NewCusUnauthorizedError(err, "token invalid", "ErrTokenInvalid"))
		// }
		user, err := token.NewTokenParser(token.WithSecretKey(s.secretKey)).Parse(tkn)
		if err != nil {
			response.Error(r.Context(), w, errors.Unauthorized(err, errors.WithReason("Unauthorized")))
			return
		}
		fmt.Println("user ", user)

		// var items []*entity.Permission
		// if obj.GetIssuer() == define.ClientIssuer {
		// 	items, err = s.clientRepo.Permissions(r.Context(), obj.GetUid())
		// } else if obj.GetIssuer() == define.UserIssuer {
		// 	roleIds, err := s.userRepo.RoleIds(r.Context(), obj.GetUid())
		// 	if err == nil {
		// 		items, err = s.userRepo.Permissions(r.Context(), roleIds)
		// 	}
		// }
		// if err != nil || len(items) == 0 {
		// 	if err == nil {
		// 		err = fmt.Errorf("missing permisisons")
		// 	}
		// 	response.Error(r.Context(), w, errors.Forbidden(err, errors.WithReason(locale.NoPerMsg.Message)))
		// 	return
		// }

		// var pers []authz.PermissionData
		// for _, v := range items {
		// 	pers = append(pers, authz.NewPermissionData(
		// 		v.Id,
		// 		v.ServiceId,
		// 		v.Code,
		// 		v.Path,
		// 		v.Method,
		// 	))
		// }

		// if !authz.New().Validate(pers, r.Method, r.URL.Path) {
		// 	response.Error(r.Context(), w, errors.Forbidden(locale.NoPerMsg))
		// 	return
		// }
		next(w, r)
	}
}
