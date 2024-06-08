package handler

import (
	"fmt"
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
	"github/tronglv_authen_author/internal/types/fosite"
	"net/http"

	fs "github.com/ory/fosite"
)

type OAuthHandler interface {
	PortalAuthorize() http.HandlerFunc
	PortalToken() http.HandlerFunc
	Token() http.HandlerFunc
}

type oauthHandler struct {
	reg     *registry.ServiceContext
	fs      fs.OAuth2Provider
	authSvc service.AuthenticationService
}

func NewOAuthHandler(reg *registry.ServiceContext) OAuthHandler {
	return &oauthHandler{
		reg:     reg,
		fs:      reg.Fosite,
		authSvc: service.NewAuthenticationService(reg),
	}
}

func (p *oauthHandler) PortalAuthorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ar, err := p.fs.NewAuthorizeRequest(ctx, r)
		// logx.Error("vao trong nay", err)
		if err != nil {
			fmt.Println("vao trong nay 123")
			fosite.WriteAuthorizeError(ctx, p.fs, w, ar, err)
			return
		}
		response.OkJson(ctx, w, "Success", nil)
	}
}

func (p *oauthHandler) PortalToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (p *oauthHandler) Token() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
