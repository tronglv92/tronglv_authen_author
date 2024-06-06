package registry

import (
	"context"

	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/rsa"
	"github/tronglv_authen_author/internal/config"
	"time"
	fs "github/tronglv_authen_author/internal/types/fosite"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"

	"github.com/zeromicro/go-zero/core/logx"
)

func NewFositeContext(c config.OAuthConfig, sqlConn db.Database, cacheClient cache.Cache) fosite.OAuth2Provider {
	conf := &fosite.Config{
		GlobalSecret:               []byte(c.HashSecret),
		AccessTokenLifespan:        time.Duration(c.AccessTokenLifespan) * time.Hour,
		RefreshTokenLifespan:       time.Duration(c.RefreshTokenLifespan) * time.Hour,
		AuthorizeCodeLifespan:      time.Duration(c.AuthorizeCodeLifespan) * time.Minute,
		SendDebugMessagesToClients: c.Debug,
	}

	keyGetter := func(ctx context.Context) (interface{}, error) {
		privateKey, err := rsa.ParsePKFromPEM(c.PrivateKey)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		return privateKey, nil
	}
	return compose.Compose(
		conf,
		fs.NewGormStore(sqlConn, cacheClient),
		compose.NewOAuth2JWTStrategy(keyGetter, compose.NewOAuth2HMACStrategy(conf), conf),
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.OAuth2AuthorizeExplicitFactory,
	)
}
