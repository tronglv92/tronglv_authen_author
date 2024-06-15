package tokenprovider

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}


