package entity

type UserGroup struct {
	IDModel
	EmployeeId string `gorm:"column:employee_id;type:varchar(20);index:idx_employee_id"`
	GroupId    int32  `gorm:"column:group_id;index:idx_group_id"`
}

func (UserGroup) TableName() string {
	return "user_groups"
}
