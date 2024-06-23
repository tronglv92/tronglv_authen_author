package entity

type User struct {
	IDModel
	Email     string `gorm:"column:email;uniqueIndex;type:varchar(50)"`
	Password  string `gorm:"column:password;not null;type:varchar(50)"`
	Salt      string `gorm:"column:salt;not null;type:varchar(50)"`
	LastName  string `gorm:"column:salt;type:varchar(50)"`
	FirstName string `gorm:"column:salt;type:varchar(50)"`
	Phone     string `gorm:"column:salt;not null;type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}
