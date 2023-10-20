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

	rInputStr := []rune(inputStr)

	if unicode.IsDigit(rInputStr[0]) {
		return "", ErrInvalidString
	}

	countRune := 1
	outStr := ""

	for i := len(rInputStr) - 1; i > -1; i-- {
		if unicode.IsDigit(rInputStr[i]) {
			if i != 0 && unicode.IsDigit(rInputStr[i-1]) {
				return "", ErrInvalidString
			}

			countRune = int(rInputStr[i] - '0')
		} else {
			outStr = strings.Repeat(string(rInputStr[i]), countRune) + outStr
			countRune = 1
		}
	}

	return outStr, nil
}
