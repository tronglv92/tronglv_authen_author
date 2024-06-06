package fosite

import (
	"context"
	"fmt"
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	rp "github/tronglv_authen_author/internal/repository"
	"time"

	"github.com/ory/fosite"
)

type gormStorage struct {
	clientRepo  rp.ClientRepository
	cacheClient cache.Cache
}

func NewGormStore(sqlConn db.Database, cacheClient cache.Cache) fosite.Storage {
	return &gormStorage{
		cacheClient: cacheClient,
		clientRepo:  rp.NewClientRepository(sqlConn),
	}
}
func (s *gormStorage) GetClient(ctx context.Context, clientId string) (fosite.Client, error) {
	// client, err := s.clientRepo.First(ctx, s.clientRepo.WithClientId(clientId))
	// if err != nil {
	// 	return nil, err
	// }
	// if client.Status != int(define.ActiveStatus) {
	// 	return nil, fmt.Errorf("the client account is inactive")
	// }
	return nil, nil
}

func (s *gormStorage) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	return fosite.ErrInvalidClient
}

func (s *gormStorage) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {
	fmt.Println("SetClientAssertionJWT")
	return nil
}
