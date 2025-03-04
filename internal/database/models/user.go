package models

type UserModel struct {
	ID           uint   `gorm:"column:id"`
	Email        string `gorm:"column:email"`
	HashPassword string `gorm:"column:hash_password"`
	Firstname    string `gorm:"column:firstname"`
	Lastname     string `gorm:"column:lastname"`
	Age          int    `gorm:"column:age"`
	RoleID       uint   `gorm:"column:role_id"`
}

func (UserModel) TableName() string {
	return "users"
}
