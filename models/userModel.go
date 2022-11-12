package models

type User struct {
	// gorm.Model
	ID       uint   `gorm:"not null;auto_increment;primary_key"`
	Username string `gorm:"unique"`
	Password string
}
