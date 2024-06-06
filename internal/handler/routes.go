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
	registerOAuthHandler(svr, h.svc)
}

func globalMiddleware(_ *rest.Server, _ *registry.ServiceContext) {
}

func registerOAuthHandler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewOAuthHandler(svc)
	var path = "/oauth"
	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/portal-authorize", path),
					Handler: h.PortalAuthorize(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/portal-token", path),
					Handler: h.PortalToken(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/portal-token", path),
					Handler: h.PortalToken(),
				},
				{
					Method:  http.MethodPost,
					Path:    fmt.Sprintf("%s/token", path),
					Handler: h.Token(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)
}
