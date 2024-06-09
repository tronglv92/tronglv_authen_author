package handler

import (
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/internal/registry"
	"github/tronglv_authen_author/internal/service"
	"github/tronglv_authen_author/internal/types/request"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ClientHandler interface {
	List() http.HandlerFunc
	Create() http.HandlerFunc
}

type clientHandler struct {
	reg       *registry.ServiceContext
	clientSvc service.ClientService
}

func NewClientHandler(reg *registry.ServiceContext) ClientHandler {
	return &clientHandler{
		reg:       reg,
		clientSvc: service.NewClientService(reg),
	}
}

func (p *clientHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.ClientListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		if err := req.SetDefault(r.Context()); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		resp, paging, err := p.clientSvc.List(r.Context(), req)
		if err != nil {
			response.Error(r.Context(), w, err)
			return
		}
		response.OkJson(r.Context(), w, resp, paging)
	}
}

func (p *clientHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.ClientSaveReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		if err := req.Validate(r.Context()); err != nil {
			response.Error(r.Context(), w, err)
			return
		}

		resp, err := p.clientSvc.Create(r.Context(), req)
		if err != nil {
			response.Error(r.Context(), w, err)
			return
		}
		response.OkJson(r.Context(), w, resp, nil)
	}
}
