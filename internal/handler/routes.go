package handler

import (
	"fmt"
	"github/tronglv_authen_author/helper/server/http/handler"
	"github/tronglv_authen_author/internal/registry"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

const (
	BasePrefix = "/identity-svc"
	RestPrefix = BasePrefix + "/api/v1"
)

type RestHandler struct {
	svc *registry.ServiceContext
}

func NewRestHandler(svc *registry.ServiceContext) RestHandler {
	return RestHandler{svc: svc}
}

func (h RestHandler) Register(svr *rest.Server) {
	handler.RegisterSwaggerHandler(svr, BasePrefix)
	globalMiddleware(svr, h.svc)
	registerOAuth2Handler(svr, h.svc)
	registerOAuthHandler(svr, h.svc)
	registerClientHandler(svr, h.svc)
}

func globalMiddleware(_ *rest.Server, _ *registry.ServiceContext) {
}

func registerOAuth2Handler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewOAuth2Handler(svc)
	var path = "/oauth2"
	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/authorize", path),
					Handler: h.PortalAuthorize(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/token", path),
					Handler: h.Token(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/session-token", path),
					Handler: h.PortalToken(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)
}
func registerOAuthHandler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewOauthHandler(svc)
	var path = "/oauth"
	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/register", path),
					Handler: h.Register(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/login", path),
					Handler: h.Login(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)

	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{
				svc.AuthInternalMiddleware,
			},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/me", path),
					Handler: h.Profile(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)
}
func registerClientHandler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewClientHandler(svc)
	var (
		path = "/clients"
	)
	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{
				// svc.AuthMiddleware,
			},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    path,
					Handler: h.List(),
				},
				{
					Method:  http.MethodPost,
					Path:    path,
					Handler: h.Create(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)

}
