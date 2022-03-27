package telegram

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

func (t *Telegram) HandleCreateTodo(ctx telebot.Context) error {
	userAuth := entities.UserAuthentication{
		AuthType:   entities.TelegramAuthType,
		ExternalID: ctx.Sender().Recipient(),
	}

	taskNames := t.handleTaskNames(ctx.Text())
	todos, users, err := t.usecases.Todo.CreateMany(context.Background(), taskNames, userAuth)
	if err != nil {
		log.Println(err)
		return ctx.Send(entities.ErrorUnknownMessage)
	}

	answer := "Задача создана!"
	if len(taskNames) > 1 {
		answer = "Задачи созданы!"
	}
	markup := t.todosToButtons(todos)
	updatingMessage := fmt.Sprintf("%s обновил(а) список задач.", ctx.Sender().FirstName)
	t.emitAll(ctx.Sender().Recipient(), users, updatingMessage, markup)

	return ctx.Send(answer, markup)
}

func (t *Telegram) handleTaskNames(message string) []string {
	rawNames := strings.Split(message, ",")
	names := make([]string, 0, len(rawNames))

	for _, name := range rawNames {
		if name != "" {
			names = append(names, name)
		}
	}

	return names
}
