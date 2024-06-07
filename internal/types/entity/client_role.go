package entity

type ClientRole struct {
	ClientId int32 `gorm:"column:client_id;index:idx_client_id"`
	RoleId   int32 `gorm:"column:role_id;index:idx_role_id"`
}

func (ClientRole) TableName() string {
	return "client_roles"
}
