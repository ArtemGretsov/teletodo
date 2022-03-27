package entities

import (
	"gorm.io/gorm"
)

type Todos []Todo
type Todo struct {
	gorm.Model
	Name     string `gorm:"index:,unique,where:deleted_at is null"`
	AuthorID uint
}
