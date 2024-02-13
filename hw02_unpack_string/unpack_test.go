package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		r    rune
		buff rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "first", args: args{r: 51, buff: 97}, want: "aaa"},
		{name: "second", args: args{r: 48, buff: 97}, want: ""},
		{name: "third", args: args{r: 53, buff: 92}, want: `\\\\\`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.r, tt.args.buff); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffE(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{name: "first", args: args{r: 49}, want: 0, wantErr: true},
		{name: "second", args: args{r: 97}, want: 97, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buffE(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("buffE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buffE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_receivedD(t *testing.T) {
	type args struct {
		r    rune
		buff rune
		sl   int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 rune
	}{
		{name: "first", args: args{r: 51, buff: 97, sl: 0}, want: "aaa", want1: 0},
		{name: "second", args: args{r: 51, buff: 92, sl: 1}, want: "", want1: 51},
		{name: "third", args: args{r: 51, buff: 92, sl: 2}, want: `\\\`, want1: 0},
		{name: "fourth", args: args{r: 51, buff: 92, sl: 3}, want: `\`, want1: 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := receivedD(tt.args.r, tt.args.buff, tt.args.sl)
			if got != tt.want {
				t.Errorf("receivedD() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("receivedD() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
