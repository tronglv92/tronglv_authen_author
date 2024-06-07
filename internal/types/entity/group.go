package entity

type Group struct {
	IDModel
	Name  string  `gorm:"column:name;type:varchar(200)"`
	Roles []*Role `gorm:"many2many:group_roles;"`
}

func (Group) TableName() string {
	return "groups"
}
