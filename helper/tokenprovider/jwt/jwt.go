package jwt

import (
	"fmt"
	"github/tronglv_authen_author/helper/tokenprovider"
	"github/tronglv_authen_author/internal/config"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenPayloadImp struct {
	UId int32
}

func (p TokenPayloadImp) UserId() int32 {
	return p.UId
}

type TokenImp struct {
	Token   string
	Created time.Time
	Expiry  int
}

func (p TokenImp) GetToken() string {
	return p.Token
}
func (p TokenImp) GetCreated() time.Time {
	return p.Created
}
func (p TokenImp) GetExpiry() int {
	return p.Expiry
}

type jwtProvider struct {
	cf config.JWTConfig
}

func NewTokenJWTProvider(cf config.JWTConfig) tokenprovider.Provider {
	return &jwtProvider{cf: cf}
}

type myClaims struct {
	Payload TokenPayloadImp `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) SecretKey() string {
	return j.cf.HashSecret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		TokenPayloadImp{
			UId: data.UserId(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Minute * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.cf.HashSecret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &TokenImp{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.cf.HashSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// validate the token
	if !res.Valid {
		// return nil, tokenprovider.ErrInvalidToken
		return nil, nil
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		// return nil, tokenprovider.ErrInvalidToken
		return nil, nil
	}

	// return the token
	return claims.Payload, nil
}
