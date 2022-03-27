package telegram

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

func (t *Telegram) HandleStart(ctx telebot.Context) error {
	teleUser := ctx.Sender()
	user := entities.User{Name: teleUser.FirstName, UserAuthentication: &entities.UserAuthentication{
		AuthType:   entities.TelegramAuthType,
		ExternalID: teleUser.Recipient(),
	}}

	todos, err := t.usecases.Todo.Start(context.Background(), user)
	if err != nil {
		log.Println(err)
		return ctx.Send(entities.ErrorUnknownMessage)
	}

	markup := t.todosToButtons(todos)
	return ctx.Send(fmt.Sprintf("Привет, %s!", teleUser.FirstName), markup)
}
