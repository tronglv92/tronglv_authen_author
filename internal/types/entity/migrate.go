package entity

func RegisterMigrate() []any {
	return []any{
		&Role{},
		&Client{},
		&ClientRole{},
		&Service{},
		&Permission{},
		&RolePermission{},
		&Group{},
		&GroupRole{},
		&UserRole{},
		&UserGroup{},
	}
}
