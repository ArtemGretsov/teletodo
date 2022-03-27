package usecases

import (
	"context"

	"gorm.io/gorm"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
	"github.com/ArtemGretsov/teletodo/internal/service/repositories"
)

type Todo interface {
	Start(ctx context.Context, user entities.User) (entities.Todos, error)
	CreateMany(ctx context.Context, names []string, userAuth entities.UserAuthentication) (entities.Todos, entities.Users, error)
	Delete(ctx context.Context, name string) (entities.Todos, entities.Users, error)
	GetAllWithAuthorName(ctx context.Context) (entities.TodosWithAuthor, error)
	GetAll(ctx context.Context) (entities.Todos, error)
}

type todo struct {
	todoRepository repositories.Todo
	userRepository repositories.User
}

func NewTodo(db *gorm.DB) Todo {
	return &todo{
		todoRepository: repositories.NewTodo(db),
		userRepository: repositories.NewUser(db),
	}
}

func (t *todo) Start(ctx context.Context, user entities.User) (entities.Todos, error) {
	_, err := t.userRepository.CreateAuth(ctx, user)
	if err != nil {
		return nil, err
	}

	return t.todoRepository.GetAll(ctx)
}

func (t *todo) CreateMany(ctx context.Context, names []string, userAuth entities.UserAuthentication) (entities.Todos, entities.Users, error) {
	user, err := t.userRepository.GetByAuth(ctx, userAuth)
	if err != nil {
		return nil, nil, err
	}

	_, err = t.todoRepository.CreateMany(ctx, names, user.ID)
	if err != nil {
		return nil, nil, err
	}

	users, err := t.userRepository.GetAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	todos, err := t.todoRepository.GetAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	return todos, users, nil
}

func (t *todo) Delete(ctx context.Context, name string) (entities.Todos, entities.Users, error) {
	err := t.todoRepository.DeleteByName(ctx, name)
	if err != nil {
		return nil, nil, err
	}

	users, err := t.userRepository.GetAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	todos, err := t.todoRepository.GetAll(ctx)
	if err != nil {
		return nil, nil, err
	}

	return todos, users, nil
}

func (t *todo) GetAllWithAuthorName(ctx context.Context) (entities.TodosWithAuthor, error) {
	return t.todoRepository.GetAllWithAuthorName(ctx)
}

func (t *todo) GetAll(ctx context.Context) (entities.Todos, error) {
	return t.todoRepository.GetAll(ctx)
}
