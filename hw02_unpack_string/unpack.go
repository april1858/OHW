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
				if unicode.IsLetter(buff) {
					str := RepeatLetter(r, buff)
					sb.WriteString(str)
					buff = 0
				} else {
					str := RepeatControl(r)
					sb.WriteString(str)
					buff = 0
				}
			case unicode.IsLetter(r):
				if unicode.IsLetter(buff) {
					sb.WriteString(string(buff))
					buff = r
				} else {
					sb.WriteString("\n")
					buff = r
				}
			case unicode.IsControl(r):
				if unicode.IsLetter(buff) {
					sb.WriteString(string(buff))
					buff = r
				} else {
					sb.WriteString("\n")
					buff = r
				}
			}
		}
	}
	if buff != 0 {
		sb.WriteString(string(buff))
	}
	return sb.String(), nil
}

func RepeatControl(r rune) string {
	var b strings.Builder
	for i := 0; i < int(r-'0'); i++ {
		b.WriteString("\n")
	}
	return b.String()
}

func RepeatLetter(r, buff rune) string {
	str := strings.Repeat(string(buff), int(r-'0'))
	return str
}
