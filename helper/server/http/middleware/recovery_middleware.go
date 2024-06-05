package middleware

import (
	"net/http"
	"github/tronglv_authen_author/helper/define"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/server/core"
	"github/tronglv_authen_author/helper/server/http/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoveryMiddleware struct {
	EnvMode string
}

func NewRecoveryMiddleware(mode string) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		EnvMode: mode,
	}
}

func (m *RecoveryMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if result := recover(); result != nil {
				if m.EnvMode == define.ProdEnv {
					logx.ErrorStack(result)
				} else {
					core.PrintStack()
				}
				response.Error(r.Context(), w, errors.InternalServer(
					errors.ToError(result),
					errors.WithReport(),
					errors.WithStack(core.SprintStack()),
				))
				return
			}
		}()
		next(w, r)
	}
}
