package middleware

import (
	"fmt"

	"github/tronglv_authen_author/helper/auth"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/locale"
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/helper/util"

	"net/http"
)

type AuthMiddleware struct {
	authSvc auth.Validate
}


func NewAuthMiddleware(opts ...auth.Option) *AuthMiddleware {
	return &AuthMiddleware{
		authSvc: auth.New(opts...),
	}
}

func (s *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := util.GetTokenFromHeader(r)
		if len(tkn) == 0 {
			response.Error(r.Context(), w, errors.BadRequest(
				fmt.Errorf("%s", "Header token is missing"),
			))
			return
		}

		data, err := s.authSvc.Validate(r.Context(), tkn)
		if err != nil {
			response.Error(r.Context(), w, errors.Unauthorized(err, errors.WithReason(locale.NoAuthMsg.Message)))
			return
		}
		next(w, r.WithContext(s.authSvc.AssignToContext(r.Context(), data)))
	}
}
