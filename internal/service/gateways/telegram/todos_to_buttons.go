package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

const TodoLabel = "✔"
const EmptyMessage = "Пусто :("

func (t *Telegram) todosToButtons(todos entities.Todos) *telebot.ReplyMarkup {
	markup := telebot.ReplyMarkup{ResizeKeyboard: true}
	rows := make([]telebot.Row, len(todos), len(todos))

	if len(todos) == 0 {
		markup.Reply([]telebot.Row{markup.Row(markup.Text(EmptyMessage))}...)

		return &markup
	}

	for i, todo := range todos {
		rows[i] = markup.Row(markup.Text(t.printTodoName(todo.Name)))
	}

	markup.Reply(rows...)

	return &markup
}

func (t *Telegram) printTodoName(name string) string {
	return fmt.Sprintf("%s %s", TodoLabel, name)
}

func (t *Telegram) removeMarkFromTodoName(name string) string {
	if len(name) == 0 {
		return ""
	}

	return string([]rune(name)[2:])
}
