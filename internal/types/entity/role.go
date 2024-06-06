package entity

type Role struct {
	IDModel
	Code        string        `gorm:"column:code;uniqueIndex;type:varchar(50)"`
	Name        string        `gorm:"column:name;type:varchar(200)"`
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}

func (Role) TableName() string {
	return "roles"
}
