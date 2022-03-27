package config

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type TelegramWhiteList string

func (t TelegramWhiteList) Split() []string {
	return strings.Split(string(t), ",")
}

func (t TelegramWhiteList) SplitInt64() (arr []int64, err error) {
	list := t.Split()

	for index, item := range list {
		numberItem, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing white list error (index: %d)", index)
		}

		arr = append(arr, int64(numberItem))
	}

	return arr, nil
}
