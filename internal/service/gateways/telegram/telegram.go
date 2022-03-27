package telegram

import (
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/telebot.v3/middleware"

	"github.com/ArtemGretsov/teletodo/config"
	"github.com/ArtemGretsov/teletodo/internal/service/entities"
	"github.com/ArtemGretsov/teletodo/internal/service/usecases"

	"gopkg.in/telebot.v3"
)

type Telegram struct {
	bot      *telebot.Bot
	config   *config.Config
	usecases *usecases.Usecases
}

func NewTelegram(cfg *config.Config, usecases *usecases.Usecases) (*Telegram, error) {
	settings := telebot.Settings{
		Token:  cfg.TelegramBot.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Telegram{
		config:   cfg,
		bot:      bot,
		usecases: usecases,
	}, nil
}

func (t *Telegram) Start() error {
	if err := t.initRouter(); err != nil {
		return err
	}
	if err := t.initCommands(); err != nil {
		return err
	}

	t.bot.Start()

	return nil
}

func (t *Telegram) initCommands() error {
	commands := []telebot.Command{
		{
			Text:        StartRoute,
			Description: "Начать работу",
		},
		{
			Text:        TextListRoute,
			Description: "Получить текстовый список задач",
		},
		{
			Text:        ListRoute,
			Description: "Получить список задач",
		},
	}

	err := t.bot.SetCommands(commands)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *Telegram) initRouter() error {
	if t.config.TelegramBot.WhiteList != "" {
		list, err := t.config.TelegramBot.WhiteList.SplitInt64()
		if err != nil {
			return err
		}

		t.bot.Use(middleware.Whitelist(list...))
	}

	t.bot.Handle(StartRoute, t.HandleStart)
	t.bot.Handle(TextListRoute, t.HandleTextList)
	t.bot.Handle(ListRoute, t.HandleList)

	t.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		message := ctx.Message().Text
		if len(message) == 0 {
			return nil
		}

		if message == EmptyMessage {
			return nil
		}

		firstChar := string([]rune(message)[0])

		switch firstChar {
		case TodoLabel:
			return t.HandleDeleteTodo(ctx)
		case "/":
			return nil
		default:
			return t.HandleCreateTodo(ctx)
		}
	})

	return nil
}

func (t *Telegram) emit(userID int64, what interface{}, opts ...interface{}) error {
	_, err := t.bot.Send(&telebot.User{ID: userID}, what, opts...)

	if err != nil {
		return errors.Wrapf(err, "telegram message seinding error (id: %d)", userID)
	}

	return nil
}

func (t *Telegram) emitAll(senderID string, users entities.Users, what interface{}, opts ...interface{}) {
	for _, user := range users {
		if user.UserAuthentication.ExternalID != senderID {
			id, err := strconv.Atoi(user.UserAuthentication.ExternalID)
			if err != nil {
				log.Println(errors.WithStack(err))
				continue
			}

			err = t.emit(int64(id), what, opts...)

			if err != nil {
				log.Println(err)
			}
		}
	}
}
