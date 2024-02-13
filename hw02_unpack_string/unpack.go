package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	var buff, bf rune
	var letter string
	var err error
	slash := 0
	for _, r := range s {
		if r == 92 {
			slash++
		}
		switch {
		case buff == 0:
			bf, err := buffE(r)
			if err != nil {
				return "", err
			}
			buff = bf
		case unicode.IsDigit(r):
			letter, bf = receivedD(r, buff, slash)
			sb.WriteString(letter)
			buff = bf
			slash = 0
		default:
			letter, bf, err = receivedC(r, buff)
			if err != nil {
				return "", err
			}
			sb.WriteString(letter)
			buff = bf
		}
	}
	if buff != 0 {
		sb.WriteString(string(buff))
	}
	return sb.String(), nil
}

func buffE(r rune) (rune, error) {
	if unicode.IsDigit(r) {
		return 0, ErrInvalidString
	}
	return r, nil
}

func receivedD(r, buff rune, sl int) (string, rune) {
	var bf rune
	var letter string
	switch {
	case buff == 92:
		switch {
		case sl == 3:
			return string(buff), r
		case sl == 2:
			letter = Repeat(r, buff)
			bf = 0
		default:
			bf = r
		}

	default:
		letter = Repeat(r, buff)
		bf = 0
	}
	return letter, bf
}

func receivedC(r, buff rune) (string, rune, error) {
	var letter string
	switch {
	case r == 110 && buff == 92:
		return "", 0, ErrInvalidString
	case r == 92 && buff == 92:
		return "", buff, nil
	default:
		letter = string(buff)
	}
	return letter, r, nil
}

func Repeat(r, buff rune) string {
	str := strings.Repeat(string(buff), int(r-'0'))
	return str
}
