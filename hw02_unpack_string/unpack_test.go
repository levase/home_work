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
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
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
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackExtended(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "один повторяющийся символ",
			input:    "a5",
			expected: "aaaaa",
		},
		{
			name:     "один символ без повторения",
			input:    "a",
			expected: "a",
		},
		{
			name:     "несколько разных символов с повторениями",
			input:    "a2b3c1",
			expected: "aabbbc",
		},
		{
			name:     "нулевое повторение в начале",
			input:    "a0bc",
			expected: "bc",
		},
		{
			name:     "нулевое повторение в середине",
			input:    "ab0c",
			expected: "ac",
		},
		{
			name:     "нулевое повторение в конце",
			input:    "abc0",
			expected: "ab",
		},
		{
			name:     "несколько нулевых повторений",
			input:    "a0b0c0",
			expected: "",
		},
		{
			name:     "unicode символы с повторениями",
			input:    "🙂2世3界1",
			expected: "🙂🙂世世世界",
		},
		{
			name:     "экранированная цифра",
			input:    `a\2b`,
			expected: "a2b",
		},
		{
			name:     "экранированный обратный слеш",
			input:    `a\\b`,
			expected: `a\b`,
		},
		{
			name:     "экранированный слеш с повторением",
			input:    `a\\3b`,
			expected: `a\\\b`,
		},
		{
			name:     "несколько экранированных символов",
			input:    `a\2\3\4b`,
			expected: "a234b",
		},
		{
			name:     "экранированный символ с повторением",
			input:    `a\23`,
			expected: "a222",
		},
		{
			name:     "сложная escape последовательность",
			input:    `a\\\2b`,
			expected: `a\2b`,
		},
		{
			name:     "смешанные экранированные и обычные цифры",
			input:    `a\2b3c\4`,
			expected: "a2bbbc4",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackBoundaryCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		isError  bool
	}{
		{
			name:     "несколько последовательных escape символов",
			input:    `a\\\\b`,
			expected: `a\\b`,
		},
		{
			name:     "экранирование с нулевым повторением",
			input:    `a\0b`,
			expected: "a0b",
		},
		{
			name:    "одиночный слеш в конце",
			input:   `\`,
			isError: true,
		},
		{
			name:    "невалидный escape (экранирование буквы)",
			input:   `\qwe`,
			isError: true,
		},
		{
			name:     "перенос строки как символ",
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if tc.isError {
				require.Error(t, err)
				require.True(t, errors.Is(err, ErrInvalidString))
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}
