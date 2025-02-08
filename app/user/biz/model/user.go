package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(100) uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(100) not null"`
}

func (User) TableName() string {
	return "user"
}
