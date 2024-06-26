package handler

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
	"github/tronglv_authen_author/internal/types/define"
	"github/tronglv_authen_author/internal/types/fosite"
	"net/http"

	fs "github.com/ory/fosite"
	"github.com/ory/fosite/handler/oauth2"
)

type OAuth2Handler interface {
	PortalAuthorize() http.HandlerFunc
	PortalToken() http.HandlerFunc
	Token() http.HandlerFunc
}

type oauth2Handler struct {
	reg     *registry.ServiceContext
	fs      fs.OAuth2Provider
	authSvc service.AuthenticationService
}

func NewOAuth2Handler(reg *registry.ServiceContext) OAuth2Handler {
	return &oauth2Handler{
		reg:     reg,
		fs:      reg.Fosite,
		authSvc: service.NewAuthenticationService(reg),
	}
}

func (p *oauth2Handler) PortalAuthorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ar, err := p.fs.NewAuthorizeRequest(ctx, r)

		if err != nil {

			fosite.WriteAuthorizeError(ctx, p.fs, w, ar, err)
			return
		}

		ar.GrantScope("offline")

		s := new(fosite.PortalSession)
		resp, err := p.fs.NewAuthorizeResponse(ctx, ar, s)
		if err != nil {
			fosite.WriteAuthorizeError(ctx, p.fs, w, ar, err)
			return
		}

		redirectUrl := fmt.Sprintf("%s?%s", ar.GetRedirectURI().String(), resp.GetParameters().Encode())
		response.OkJson(ctx, w, fosite.AuthorizeResp{RedirectUrl: redirectUrl}, nil)
	}
}

func (p *oauth2Handler) PortalToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, define.OAuthClientKey, r.FormValue("client_id"))

		s := new(fosite.PortalSession)
		ar, err := p.fs.NewAccessRequest(ctx, r, s)
		if err != nil {
			p.fs.WriteAccessError(ctx, w, ar, err)
			return
		}

		pSession, ok := ar.GetSession().(*fosite.PortalSession)
		if !ok {
			p.fs.WriteAccessError(ctx, w, ar, fmt.Errorf("the portal session context unknown"))
			return
		}

		if pSession.Token == nil {
			p.fs.WriteAccessError(ctx, w, ar, fmt.Errorf("the token does not exist"))
			return
		}

		response.Write(w, http.StatusOK, pSession.GetToken())

	}
}

func (p *oauth2Handler) Token() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, define.OAuthClientKey, r.FormValue("client_id"))
		s := new(oauth2.JWTSession)

		ar, err := p.fs.NewAccessRequest(ctx, r, s)
		if err != nil {
			p.fs.WriteAccessError(ctx, w, ar, err)
			return
		}

		claims, e := p.authSvc.ClientClaims(r.Context(), ar.GetClient().GetID())
		if e != nil {
			p.fs.WriteAccessError(ctx, w, ar, e)
			return
		}

		s.JWTClaims = claims

		ar.SetSession(s)
		if ar.GetGrantTypes().ExactOne(define.GrantClientCredential) {
			for _, scope := range claims.Scope {
				ar.GrantScope(scope)
			}
		}

		// for _, scope := range claims.Scope {
		// 	ar.GrantScope(scope)
		// }

		fmt.Println("Token ar:", ar)

		resp, err := p.fs.NewAccessResponse(ctx, ar)
		if err != nil {
			p.fs.WriteAccessError(ctx, w, ar, err)
			return
		}

		fmt.Println("Token ar.GetGrantTypes():", ar.GetGrantedScopes())
		p.fs.WriteAccessResponse(ctx, w, ar, resp)
	}
}
