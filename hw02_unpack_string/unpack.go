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

			if i < stringLen-1 && s[i+1] == ZERO { // aaa0b => aab
				i += 2
			} else { // write the symbol as it is
				b.WriteString(string(s[i]))
				i++
			}

		case unicode.IsDigit(rune(s[i])):

			if i < stringLen-1 && unicode.IsDigit(rune(s[i+1])) { // two digits in a row "45"
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(string(s[i]))
			if err != nil {
				fmt.Println("error strconv - ", err)
			} else {
				b.WriteString(strings.Repeat(string(s[i-1]), n-1)) // "a4bc2d5e" => "aaaabccddddde". a + aaa
				i++
			}

		default:

			switch {
			case i == stringLen-1:
				return "", ErrInvalidString // `qwe\`

			case unicode.IsDigit(rune(s[i+1])): // all with slashes
				b.WriteString(string(s[i+1]))
				i += 2

			case s[i+1] == BACKSLASH: // more than one slash
				b.WriteString(string(s[i+1]))
				i += 2
			case s[i+1] != BACKSLASH: // `qwe\t`
				return "", ErrInvalidString
			}
		}
	}
	return b.String(), nil
}
