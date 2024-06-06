package entity

type Permission struct {
	IDModel
	ServiceId int32  `gorm:"column:service_id;uniqueIndex:idx_service_endpoint"`
	GroupName string `gorm:"column:group_name;type:varchar(150)"`
	Code      string `gorm:"column:code;index;type:varchar(50)"`
	Name      string `gorm:"column:name;type:varchar(200)"`
	Path      string `gorm:"column:path;type:varchar(50);uniqueIndex:idx_service_endpoint"`
	Method    string `gorm:"column:method;type:varchar(50);uniqueIndex:idx_service_endpoint"`
	System    bool   `gorm:"column:system"`
	Service   Service
}

func (Permission) TableName() string {
	return "permissions"
}
