package fosite

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mohae/deepcopy"
	"github.com/ory/fosite"
	"golang.org/x/oauth2"
	"time"
)

type PortalSession struct {
	ExpiresAt map[fosite.TokenType]time.Time `json:"expires_at"`
	Username  string                         `json:"username"`
	Subject   string                         `json:"subject"`
	Extra     map[string]interface{}         `json:"extra"`
	Token     *oauth2.Token                  `json:"token"`
}

func (s *PortalSession) SetToken(t string) error {
	token, _, err := new(jwt.Parser).ParseUnverified(t, jwt.MapClaims{})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("can not parse jwt mapclaims")
	}
	if err = claims.Valid(); err != nil {
		return err
	}

	var expiry time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		expiry = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		expiry = time.Unix(v, 0)
	}

	s.Token = &oauth2.Token{
		AccessToken: t,
		Expiry:      expiry,
		TokenType:   fosite.BearerAccessToken,
	}
	return nil
}

func (s *PortalSession) GetToken() *oauth2.Token {
	return s.Token
}

func (s *PortalSession) SetExpiresAt(key fosite.TokenType, exp time.Time) {
	if s.ExpiresAt == nil {
		s.ExpiresAt = make(map[fosite.TokenType]time.Time)
	}
	s.ExpiresAt[key] = exp
}

func (s *PortalSession) GetExpiresAt(key fosite.TokenType) time.Time {
	if s.ExpiresAt == nil {
		s.ExpiresAt = make(map[fosite.TokenType]time.Time)
	}

	if _, ok := s.ExpiresAt[key]; !ok {
		return time.Time{}
	}
	return s.ExpiresAt[key]
}

func (s *PortalSession) GetUsername() string {
	if s == nil {
		return ""
	}
	return s.Username
}

func (s *PortalSession) SetSubject(subject string) {
	s.Subject = subject
}

func (s *PortalSession) GetSubject() string {
	if s == nil {
		return ""
	}

	return s.Subject
}

func (s *PortalSession) Clone() fosite.Session {
	if s == nil {
		return nil
	}
	return deepcopy.Copy(s).(fosite.Session)
}
