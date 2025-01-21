package reflector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_upperToPascal(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic conversion",
			input:    "THIS_IS_SNAKE_CASE",
			expected: "ThisIsSnakeCase",
		},
		{
			name:     "single word",
			input:    "SINGLE",
			expected: "Single",
		},
		{
			name:     "two words with a single underscore",
			input:    "ONE_WORD",
			expected: "OneWord",
		},
		{
			name:     "leading underscore",
			input:    "_LEADING_UNDERSCORE",
			expected: "LeadingUnderscore",
		},
		{
			name:     "trailing underscore",
			input:    "TRAILING_UNDERSCORE_",
			expected: "TrailingUnderscore",
		},
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
		{
			name:     "multiple underscores",
			input:    "MULTIPLE_______UNDERSCORES",
			expected: "MultipleUnderscores",
		},
		{
			name:     "letters and numbers",
			input:    "KEY_123_EXAMPLE",
			expected: "Key123Example",
		},
		{
			name:     "single letter words",
			input:    "A_B_C",
			expected: "ABC",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			assert := assert.New(t)

			// Act
			result := upperToPascal(tc.input)

			// Assert
			assert.Equal(tc.expected, result)
		})
	}
}
