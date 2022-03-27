package telegram

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

func (t *Telegram) HandleDeleteTodo(ctx telebot.Context) error {
	task := t.removeMarkFromTodoName(ctx.Message().Text)
	todos, users, err := t.usecases.Todo.Delete(context.Background(), task)
	if err != nil {
		log.Println(err)
		return ctx.Send(entities.ErrorUnknownMessage)
	}

	markup := t.todosToButtons(todos)
	answerMessage := fmt.Sprintf("Задача \"%s\" выполнена!", task)

	updatingMessage := fmt.Sprintf("%s пометил(а) задачу \"%s\" как выполненную.", ctx.Sender().FirstName, task)
	t.emitAll(ctx.Sender().Recipient(), users, updatingMessage, markup)

	return ctx.Send(answerMessage, markup)
}
