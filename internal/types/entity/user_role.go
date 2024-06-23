package entity

type UserRole struct {
	IDModel
	UserId string `gorm:"column:user_id;type:varchar(20);index:idx_employee_id"`
	RoleId int32  `gorm:"column:role_id;index:idx_role_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
