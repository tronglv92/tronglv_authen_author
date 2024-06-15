package response

import (
	"context"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/locale"
	"github/tronglv_authen_author/helper/logify"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/trace"
)

func setHttpResponse(ctx context.Context, result any, paging any, err errors.Error) any {
	dt := data{}
	var msgKey, msgResp = locale.SuccessMsg.Key, locale.SuccessMsg.Message
	if err != nil {
		msgKey, msgResp = err.GetMetaCode(), err.GetReason()
	}
	return responseHttp{
		Meta: metaResponse{
			TradeId: trace.SpanFromContext(ctx).SpanContext().TraceID().String(),
			Code:    msgKey,
			Message: msgResp,
		},
		Data: dt.SetData(result, paging),
	}
}

func Error(ctx context.Context, w http.ResponseWriter, err error) {
	status, e := parseError(ctx, err)
	httpx.WriteJsonCtx(ctx, w, status, setHttpResponse(ctx, nil, nil, e))
	return
}

func parseError(ctx context.Context, err error) (int, errors.Error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	e := errors.From(err)
	if e.HasReport() {
		logify.NewReport().Send(ctx, e)
	}
	if span.IsRecording() {
		span.RecordError(err)
	}
	return e.GetCode(), e
}

func OkJson(ctx context.Context, w http.ResponseWriter, result any, paging any) {
	httpx.OkJsonCtx(ctx, w, setHttpResponse(ctx, result, paging, nil))
	return
}
func Write(w http.ResponseWriter, status int, resp any) {
	httpx.WriteJson(w, status, resp)
	return
}
