package config

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type StringEnumerable string

func (t StringEnumerable) Split() []string {
	return strings.Split(string(t), ",")
}

func (t StringEnumerable) SplitInt64() (arr []int64, err error) {
	list := t.Split()

	for index, item := range list {
		numberItem, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing enumerable error (index: %d)", index)
		}

		arr = append(arr, int64(numberItem))
	}

	return arr, nil
}
