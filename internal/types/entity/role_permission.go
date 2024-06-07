package entity

type RolePermission struct {
	RoleId       int32 `gorm:"column:role_id;index:idx_role_id"`
	PermissionId int32 `gorm:"column:permission_id;index:idx_permission_id"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
