package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

func (User) TableName() string {
	return "users"
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
