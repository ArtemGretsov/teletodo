package telegram

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

func (t *Telegram) HandleTextList(ctx telebot.Context) error {
	todos, err := t.usecases.Todo.GetAllWithAuthorName(context.Background())
	if err != nil {
		log.Println(err)
		return ctx.Send(entities.ErrorUnknownMessage)
	}

	if len(todos) == 0 {
		return nil
	}

	answer := ""

	for i, todo := range todos {
		date := todo.CreatedAt.Format("02.01.2006 15:04")
		answer += fmt.Sprintf("%d. %s (%s, %s)\n", i+1, todo.Name, todo.AuthorName, date)
	}

	return ctx.Send(answer)
}
