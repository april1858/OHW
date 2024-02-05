package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	var buff rune
	for _, r := range s {
		if buff == 0 {
			if unicode.IsDigit(r) {
				return "", ErrInvalidString
			}
			buff = r
		} else {
			switch {
			case unicode.IsDigit(r):
				str := Repeat(r, buff)
				sb.WriteString(str)
				buff = 0
			default:
				sb.WriteString(string(buff))
				buff = r
			}
		}
	}
	if buff != 0 {
		sb.WriteString(string(buff))
	}
	return sb.String(), nil
}

func Repeat(r, buff rune) string {
	str := strings.Repeat(string(buff), int(r-'0'))
	return str
}
