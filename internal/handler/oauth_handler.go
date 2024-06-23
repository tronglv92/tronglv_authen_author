package handler

import (
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
	"github/tronglv_authen_author/internal/types/request"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type OAuthHandler interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Profile() http.HandlerFunc
}

type oauthHandler struct {
	reg     *registry.ServiceContext
	userSvc service.UserService
}

func NewOauthHandler(reg *registry.ServiceContext) OAuthHandler {
	return &oauthHandler{
		reg:     reg,
		userSvc: service.NewUserService(reg),
	}
}

func (p *oauthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.RegisterReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		if err := req.Validate(r.Context()); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		resp, err := p.userSvc.Register(r.Context(), req)
		if err != nil {
			response.Error(r.Context(), w, err)
			return
		}
		response.OkJson(r.Context(), w, resp, nil)
	}
}

func (p *oauthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (p *oauthHandler) Profile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
