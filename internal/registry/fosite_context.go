package registry

import (
	"context"

	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/rsa"
	"github/tronglv_authen_author/internal/config"
	fs "github/tronglv_authen_author/internal/types/fosite"
	"time"

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
		// privateKey, err := rsa.ReadPrivateKeyFromFile("private_key1.pem")
		// fmt.Println("c.PrivateKey", c.PrivateKey)
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
		//server-to-server communication
		// 1. The client sends a request to the token endpoint with its client ID and secret.
		// 2. The authorization server validates the credentials.
		// 3. If valid, the server issues an access token to the client.
		compose.OAuth2ClientCredentialsGrantFactory,

		// client-to-server
		// 1. The client directs the user to the authorization server with a request for authorization.
		// 2. The user logs in and approves the request.
		// 3. The authorization server redirects the user back to the client with an authorization code.
		// 4. The client exchanges the authorization code for an access token.
		compose.OAuth2AuthorizeExplicitFactory,
	)
}
