package tokenprovider

import "time"

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() string
}

type Token interface {
	GetToken() string
	GetCreated() time.Time
	GetExpiry() int
}
