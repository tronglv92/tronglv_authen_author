package entity

import (
	"github.com/lib/pq"
	"github.com/ory/fosite"
)

type Client struct {
	IDModel
	Name             string         `gorm:"column:name;type:varchar(100)"`
	ClientId         string         `gorm:"column:client_id;uniqueIndex;type:varchar(50)"`
	ClientSecret     string         `gorm:"column:client_secret;index:idx_client_secret;type:varchar(100)"`
	ClientSecretHash string         `gorm:"column:client_secret_hash;type:varchar(100)"`
	Status           int            `gorm:"column:status"`
	Public           *bool          `gorm:"column:public"`
	Scopes           pq.StringArray `gorm:"column:scopes;type:text[]"`
	Grants           pq.StringArray `gorm:"column:grants;not null;type:text[]"`
	Audiences        pq.StringArray `gorm:"column:audiences;not null;type:text[]"`
	RedirectUrls     pq.StringArray `gorm:"column:redirect_urls;not null;type:text[]"`
	ResponseTypes    pq.StringArray `gorm:"column:response_types;not null;type:text[]"`
	Roles            []*Role        `gorm:"many2many:client_roles;"`
}

func (Client) TableName() string {
	return "clients"
}

func (c Client) GetID() string {
	return c.ClientId
}

func (c Client) GetHashedSecret() []byte {
	return []byte(c.ClientSecretHash)
}

func (c Client) GetRedirectURIs() []string {
	return c.RedirectUrls
}

func (c Client) GetGrantTypes() fosite.Arguments {
	return fosite.Arguments(c.Grants)
}

func (c Client) GetResponseTypes() fosite.Arguments {
	return fosite.Arguments(c.ResponseTypes)
}

func (c Client) GetScopes() fosite.Arguments {
	return fosite.Arguments(c.Scopes)
}

func (c Client) IsPublic() bool {
	return *c.Public
}

func (c Client) GetAudience() fosite.Arguments {
	return fosite.Arguments(c.Audiences)
}

type ClientPermission struct {
	ServiceId int32
	Code      string
	Name      string
	Path      string
	Method    string
	System    bool
}
