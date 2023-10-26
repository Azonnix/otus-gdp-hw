package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) {
	if inputStr == "" {
		return "", nil
	}

	inputStr = strings.TrimSpace(inputStr)
	rInputStr := []rune(inputStr)

	if unicode.IsDigit(rInputStr[0]) {
		return "", ErrInvalidString
	}

	stringBuilder := strings.Builder{}
	lastIndex := len(rInputStr) - 1

	for i := 0; i <= lastIndex; i++ {
		if unicode.IsDigit(rInputStr[i]) {
			if i != lastIndex && unicode.IsDigit(rInputStr[i+1]) {
				return "", ErrInvalidString
			}
		} else {
			if i != lastIndex && unicode.IsDigit(rInputStr[i+1]) {
				stringBuilder.WriteString(strings.Repeat(string(rInputStr[i]), int(rInputStr[i+1]-'0')))
			} else {
				stringBuilder.WriteString(string(rInputStr[i]))
			}
		}
	}

	return stringBuilder.String(), nil
}
