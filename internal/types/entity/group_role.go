package entity

type GroupRole struct {
	GroupId int32 `gorm:"column:group_id;index:idx_group_id"`
	RoleId  int32 `gorm:"column:role_id;index:idx_role_id"`
}

func (GroupRole) TableName() string {
	return "group_roles"
}
