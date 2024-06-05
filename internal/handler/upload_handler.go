package handler

import (
	"net/http"
	"github/tronglv_authen_author/helper/server/http/response"
	"github/tronglv_authen_author/internal/registry"
)

type UploadHandler interface {
	UploadFileS3() http.HandlerFunc
}

type uploadHandler struct {
	reg *registry.ServiceContext
}

func NewUploadHandler(reg *registry.ServiceContext) UploadHandler {
	return &uploadHandler{
		reg: reg,
	}
}

func (p *uploadHandler) UploadFileS3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// response.Error(r.Context(), w, err)
		// var data []string
		response.OkJson(r.Context(), w, "Upload Success", nil)

	}
}
