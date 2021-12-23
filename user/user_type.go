package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30);not null;primaryKey;unique"`
	Password string `gorm:"size:255;not null"`
}
