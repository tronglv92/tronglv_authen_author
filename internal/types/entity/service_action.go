package entity

type ServiceAction struct {
	IDModel
	ServiceId int32  `gorm:"column:service_id"`
	Name      string `gorm:"column:name;type:varchar(200)"`
	Code      string `gorm:"column:code;uniqueIndex;type:varchar(50)"`
}

func (ServiceAction) TableName() string {
	return "service_actions"
}
