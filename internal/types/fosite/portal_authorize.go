package fosite

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/helper/server/http/response"
	"net/http"

	
	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func getLangFromRequester(requester fosite.Requester) language.Tag {
	lang := language.English
	g11nContext, ok := requester.(fosite.G11NContext)
	if ok {
		lang = g11nContext.GetLang()
	}

	return lang
}

func WriteAuthorizeError(ctx context.Context, foa fosite.OAuth2Provider, rw http.ResponseWriter, ar fosite.AuthorizeRequester, err error) {
	f, ok := foa.(*fosite.Fosite)
	if !ok {
		response.Error(ctx, rw, fmt.Errorf("the client context unknown"))
		return
	}

	rfCert := fosite.ErrorToRFC6749Error(err).WithExposeDebug(f.Config.GetSendDebugMessagesToClients(ctx)).WithLocalizer(f.Config.GetMessageCatalog(ctx), getLangFromRequester(ar))
	v, err := util.AnyToStruct[AuthorizeResp](rfCert)
	if err != nil {
		response.Error(ctx, rw, err)
		return
	}
	v.Debug = f.Config.GetSendDebugMessagesToClients(ctx)
	response.Error(ctx, rw, v)
	return
}
