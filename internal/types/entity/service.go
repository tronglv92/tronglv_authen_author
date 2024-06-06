package entity

import "github.com/lib/pq"

type Service struct {
	IDModel
	Name           string         `gorm:"column:name;type:varchar(200)"`
	Status         int            `gorm:"column:status"`
	Urls           pq.StringArray `gorm:"column:redirect_urls;not null;type:text[]"`
	Permissions    []*Permission
	ServiceActions []ServiceAction
}

func (Service) TableName() string {
	return "services"
}
