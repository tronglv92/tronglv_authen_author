package handler

import (
	"github.com/ory/fosite"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
	fs "github/tronglv_authen_author/internal/types/fosite"
	"net/http"
)

type OAuthHandler interface {
	PortalAuthorize() http.HandlerFunc
	PortalToken() http.HandlerFunc
	Token() http.HandlerFunc
}

type oauthHandler struct {
	reg     *registry.ServiceContext
	fs      fosite.OAuth2Provider
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
		if err != nil {
			fs.WriteAuthorizeError(ctx, p.fs, w, ar, err)
			return
		}
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
