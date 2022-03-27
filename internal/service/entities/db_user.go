package entities

import "gorm.io/gorm"

type Users []User
type User struct {
	gorm.Model
	Name               string `gorm:"type:varchar"`
	UUID               string `gorm:"type:uuid"`
	Todo               Todo   `gorm:"foreignKey:AuthorID"`
	UserAuthentication *UserAuthentication
}
