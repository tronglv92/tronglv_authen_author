package jwt

import (
	"fmt"
	"github/tronglv_authen_author/helper/tokenprovider"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenPayloadImp struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayloadImp) UserId() int {
	return p.UId
}

func (p TokenPayloadImp) Role() string {
	return p.URole
}

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider() *jwtProvider {
	return &jwtProvider{}
}

type myClaims struct {
	Payload TokenPayloadImp `json:"payload"`
	jwt.StandardClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		TokenPayloadImp{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

// func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
// 	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(j.secret), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	// validate the token
// 	if !res.Valid {
// 		return nil, tokenprovider.ErrInvalidToken
// 	}

// 	claims, ok := res.Claims.(*myClaims)

// 	if !ok {
// 		return nil, tokenprovider.ErrInvalidToken
// 	}

// 	// return the token
// 	return claims.Payload, nil
// }
