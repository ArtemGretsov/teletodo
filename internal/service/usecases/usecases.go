package usecases

import "gorm.io/gorm"

type Usecases struct {
	Todo Todo
}

func NewUsecases(db *gorm.DB) *Usecases {
	return &Usecases{
		Todo: NewTodo(db),
	}
}
