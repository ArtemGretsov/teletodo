package repositories

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

type Todo interface {
	GetAllWithAuthorName(ctx context.Context) (entities.TodosWithAuthor, error)
	DeleteByName(ctx context.Context, name string) error
	CreateMany(ctx context.Context, names []string, authorID uint) (entities.Todos, error)
	GetAll(ctx context.Context) (todos entities.Todos, err error)
}

type todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &todo{
		db: db,
	}
}

func (t *todo) GetAll(ctx context.Context) (todos entities.Todos, err error) {
	err = t.db.WithContext(ctx).Find(&todos).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return todos, nil
}

func (t *todo) GetAllWithAuthorName(ctx context.Context) (todos entities.TodosWithAuthor, err error) {
	err = t.db.
		WithContext(ctx).
		Model(entities.Todo{}).
		Joins("left join users on users.id = todos.author_id").
		Select("todos.*, users.name as author_name").
		Scan(&todos).
		Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return todos, nil
}

func (t *todo) DeleteByName(ctx context.Context, name string) error {
	err := t.db.WithContext(ctx).Delete(&entities.Todo{}, "name = ?", name).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *todo) CreateMany(ctx context.Context, names []string, authorID uint) (entities.Todos, error) {
	todos := make(entities.Todos, len(names), len(names))

	for i, todoName := range names {
		todos[i] = entities.Todo{Name: todoName, AuthorID: authorID}
	}

	err := t.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&todos).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return todos, nil
}
