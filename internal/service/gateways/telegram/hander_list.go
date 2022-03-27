package telegram

import (
	"context"
	"log"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

func (t *Telegram) HandleList(ctx telebot.Context) error {
	todos, err := t.usecases.Todo.GetAll(context.Background())
	if err != nil {
		log.Println(err)
		return ctx.Send(entities.ErrorUnknownMessage)
	}

	markup := t.todosToButtons(todos)

	return ctx.Send("Задачи:", markup)
}
