package telegram

import (
	"strconv"
	"strings"
)

func EncodeQueryString(str string) (*int, *int) {
	split := strings.Split(str, "-")

	messageId, err := strconv.Atoi(split[0])

	if err != nil {
		return nil, nil
	}
	index, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, nil
	}
	return &messageId, &index
}
