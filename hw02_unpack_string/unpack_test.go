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
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		// additional
		{input: "", expected: ""},
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "a", expected: "a"},
		{input: "aaa0b", expected: "aab"},
		{input: "日本語4", expected: "日本語語語語"},
		{input: "日本4語", expected: "日本本本本語"},
		{input: ",4,.m", expected: ",,,,,.m"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "\u65e54\u8a9e", expected: "\u65e5\u65e5\u65e5\u65e5\u8a9e"},
		{input: "\U000065e5\U00008a9e3", expected: "\U000065e5\U00008a9e\U00008a9e\U00008a9e"},
		{input: `\3`, expected: "3"},
		{input: `\\`, expected: `\`},
		{input: "a0", expected: ""},
		{input: "a0b0c0d0", expected: ""},
		{input: "a0b0c0d0e", expected: "e"},
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
	invalidStrings := []string{"3abc", "45", "aaa10b", `qwe\`, "a45a", "a45", `qw\ne`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
