package fosite

import (
	"context"
	"encoding/json"
	"fmt"
	"github/tronglv_authen_author/helper/cache"
	db "github/tronglv_authen_author/helper/database"
	"github/tronglv_authen_author/helper/define"

	"github/tronglv_authen_author/helper/util"
	rp "github/tronglv_authen_author/internal/repository"
	"github/tronglv_authen_author/internal/types/entity"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ory/fosite"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizeCode struct {
	ID                string           `json:"id"`
	RequestedAt       time.Time        `json:"requested_at"`
	Client            *entity.Client   `json:"client"`
	RequestedScope    fosite.Arguments `json:"scopes"`
	GrantedScope      fosite.Arguments `json:"granted_scopes"`
	Form              url.Values       `json:"form"`
	Session           *PortalSession   `json:"session"`
	RequestedAudience fosite.Arguments `json:"requested_audience"`
	GrantedAudience   fosite.Arguments `json:"granted_audience"`
}

type myClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.StandardClaims
}

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}
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
	fmt.Println("GetClient")
	client, err := s.clientRepo.First(ctx, s.clientRepo.WithClientId(clientId))
	if err != nil {
		return nil, err
	}
	if client.Status != int(define.ActiveStatus) {
		return nil, fmt.Errorf("the client account is inactive")
	}
	return client, nil
}

func (s *gormStorage) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	fmt.Println("ClientAssertionJWTValid")
	return fosite.ErrInvalidClient
}

func (s *gormStorage) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {
	fmt.Println("SetClientAssertionJWT")
	return nil
}

func (s *gormStorage) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	fmt.Println("CreateAccessTokenSession")
	return nil
}

func (s *gormStorage) DeleteAccessTokenSession(ctx context.Context, signature string) (err error) {
	fmt.Println("DeleteAccessTokenSession")
	return nil
}

func (s *gormStorage) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	fmt.Println("GetAccessTokenSession")
	return nil, nil
}

func (s *gormStorage) CreateRefreshTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	fmt.Println("CreateRefreshTokenSession")
	return nil
}

func (s *gormStorage) DeleteRefreshTokenSession(ctx context.Context, signature string) (err error) {
	fmt.Println("DeleteRefreshTokenSession")
	return nil
}

func (s *gormStorage) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	fmt.Println("GetRefreshTokenSession")
	return nil, nil
}

func (s *gormStorage) InvalidateAuthorizeCodeSession(ctx context.Context, code string) (err error) {
	fmt.Println("InvalidateAuthorizeCodeSession")
	return nil
}

func (s *gormStorage) RevokeAccessToken(ctx context.Context, requestID string) error {
	fmt.Println("RevokeAccessToken")
	return nil
}

func (s *gormStorage) RevokeRefreshToken(ctx context.Context, requestID string) error {
	fmt.Println("RevokeRefreshToken")
	return nil
}

func (s *gormStorage) RevokeRefreshTokenMaybeGracePeriod(ctx context.Context, requestID string, signature string) error {
	fmt.Println("RevokeRefreshTokenMaybeGracePeriod")
	return nil
}

func (s *gormStorage) generateToken() (*string, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		TokenPayload{
			UId:   1,
			URole: "user",
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(5)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte("12345"))
	if err != nil {
		return nil, err
	}
	return &myToken, nil
}

func (s *gormStorage) CreateAuthorizeCodeSession(ctx context.Context, code string, request fosite.Requester) error {
	fmt.Println("CreateAuthorizeCodeSession")
	// info, err := auth.GetAuthData(ctx)
	// if err != nil {
	// 	return err
	// }

	client := request.GetClient()

	// fmt.Println("info: ", info)

	eClient, ok := client.(*entity.Client)
	if !ok {
		return fmt.Errorf("the client context unknown")
	}
	fmt.Println("client: ", client)
	eSession, ok := request.GetSession().(*PortalSession)
	if !ok {
		return fmt.Errorf("the portal session context unknown")
	}
	accessToken, err := s.generateToken()
	if err != nil {
		return err
	}

	if err := eSession.SetToken(*accessToken); err != nil {
		return err
	}

	authorizeCode := AuthorizeCode{
		ID:                request.GetID(),
		RequestedAt:       request.GetRequestedAt(),
		Client:            eClient,
		RequestedScope:    request.GetRequestedScopes(),
		GrantedScope:      request.GetGrantedScopes(),
		Form:              request.GetRequestForm(),
		Session:           eSession,
		RequestedAudience: request.GetRequestedAudience(),
		GrantedAudience:   request.GetGrantedAudience(),
	}

	val := util.Marshal(authorizeCode)
	fmt.Println("CreateAuthorizeCodeSession authorizeCode:", val)
	if err := s.cacheClient.SetWithExpire(s.authorizeCodeKey(client.GetID(), code), val, 5*time.Minute); err != nil {
		fmt.Println("SetWithExpire err: ", err)
		logx.Error(err)
		return err
	}
	return nil
}

func (s *gormStorage) authorizeCodeKey(key, code string) string {
	fmt.Println("authorizeCodeKey")
	return fmt.Sprintf(
		"%s_%s",
		strings.ToLower(s.replaceAlb(key)),
		util.Md5Hash(s.replaceAlb(code)),
	)
}

func (s *gormStorage) replaceAlb(key string) string {
	rex := regexp.MustCompile("[^a-zA-Z0-9]")
	return rex.ReplaceAllString(key, "")
}

func (s *gormStorage) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (fosite.Requester, error) {
	fmt.Println("GetAuthorizeCodeSession")
	clientId, ok := ctx.Value("client-id").(string)
	if !ok {
		return nil, fmt.Errorf("client id can not be blank")
	}

	var result string
	keyCache := s.authorizeCodeKey(clientId, code)

	if err := s.cacheClient.Get(keyCache, &result); err != nil {
		fmt.Println("vao trong nay", err)
		return nil, err
	}
	defer func(cacheClient cache.Cache, key string) {
		if err := cacheClient.Del(key); err != nil {
			logx.Error(err)
		}
	}(s.cacheClient, keyCache)

	if len(result) == 0 {
		return nil, fmt.Errorf("the authorization code is invalid or has expired")
	}

	fmt.Println("GetAuthorizeCodeSession result:", result)
	var reqData AuthorizeCode
	if err := json.Unmarshal([]byte(result), &reqData); err != nil {
		return nil, err
	}

	return &fosite.Request{
		ID:                reqData.ID,
		RequestedAt:       reqData.RequestedAt,
		Client:            reqData.Client,
		RequestedScope:    reqData.RequestedScope,
		GrantedScope:      reqData.GrantedScope,
		Form:              reqData.Form,
		Session:           reqData.Session,
		RequestedAudience: reqData.RequestedAudience,
		GrantedAudience:   reqData.GrantedAudience,
	}, nil

}
