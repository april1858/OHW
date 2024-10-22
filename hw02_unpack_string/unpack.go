package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	BACKSLASH byte = 92
	ZERO      byte = 48
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	stringLen := len(s)
	var b strings.Builder

	if stringLen == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	i := 0

	for i < stringLen {
		switch {
		case !unicode.IsDigit(rune(s[i])) && s[i] != BACKSLASH:

			if i < stringLen-1 && s[i+1] == ZERO {
				i += 2
			} else {
				b.WriteString(string(s[i]))
				i++
			}

		case unicode.IsDigit(rune(s[i])):

			if i < stringLen-1 && unicode.IsDigit(rune(s[i+1])) {
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(string(s[i]))
			if err != nil {
				fmt.Println("errir - ", err)
			} else {
				b.WriteString(strings.Repeat(string(s[i-1]), n-1))
				i++
			}

		default:

			switch {
			case i == stringLen-1:
				return "", ErrInvalidString

			case unicode.IsDigit(rune(s[i+1])):
				b.WriteString(string(s[i+1]))
				i += 2

			case s[i+1] == BACKSLASH:
				b.WriteString(string(s[i+1]))
				i += 2
			case s[i+1] != BACKSLASH:
				return "", ErrInvalidString
			}
		}
	}
	return b.String(), nil
}
