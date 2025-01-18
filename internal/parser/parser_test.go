package parser

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	testcases := []struct {
		name           string
		fixtureName    string
		expectedParsed map[string]string
		expectedErr    *errors.GofidentialError
	}{
		{
			name:        "should ignore trailing spaces",
			fixtureName: "trailing_spaces.env",
			expectedParsed: map[string]string{
				"FIRST_EXAMPLE":  "some random value",
				"SECOND_EXAMPLE": "second value",
				"THIRD_EXAMPLE":  "third value",
			},
		},
		{
			name:        "should ignore empty lines",
			fixtureName: "empty_lines.env",
			expectedParsed: map[string]string{
				"FIRST_KEY":  "first value",
				"SECOND_KEY": "second value",
				"THIRD_KEY":  "third value",
			},
		},
		{
			name:        "should ignore comments",
			fixtureName: "comments.env",
			expectedParsed: map[string]string{
				"EXAMPLE_KEY": "example value",
				"ANOTHER_KEY": "another value",
			},
		},
		{
			name:           "should return an empty map if the file is empty",
			fixtureName:    "empty_file.env",
			expectedParsed: map[string]string{},
		},
		{
			name:        "should return an error if assignment operator is missing",
			fixtureName: "missing_assignment.env",
			expectedErr: newMissingAssignmentError(3),
		},
		{
			name:        "should return an error if key is missing",
			fixtureName: "missing_key.env",
			expectedErr: newMissingKeyError(2),
		},
		{
			name:        "should return an error if there are leading spaces before key",
			fixtureName: "leading_spaces.env",
			expectedErr: newLeadingSpaceError(2),
		},
		{
			name:        "should return an error if there are spaces before the assignment operator",
			fixtureName: "space_before_assignment.env",
			expectedErr: newSpacedSeparatorError(2),
		},
		{
			name:        "should return an error if there are spaces after the assignment operator",
			fixtureName: "space_after_assignment.env",
			expectedErr: newSpacedSeparatorError(2),
		},
		{
			name:        "should return an error if there are spaces around the assignment operator",
			fixtureName: "space_around_assignment.env",
			expectedErr: newSpacedSeparatorError(2),
		},
		{
			name:        "should return an error if the value is missing",
			fixtureName: "missing_value.env",
			expectedErr: newMissingValueError(2),
		},
		{
			name:        "should return an error if the value open quotes is a single quote",
			fixtureName: "open_single_quoted_value.env",
			expectedErr: newSingleQuotesError(2),
		},
		{
			name:        "should return an error if the value close quotes is a single quote",
			fixtureName: "close_single_quoted_value.env",
			expectedErr: newUnterminatedQuoteError(2),
		},
		{
			name:        "should return an error if the value is partially unquoted",
			fixtureName: "partially_unquoted_value.env",
			expectedErr: newUnquotedValueError(2),
		},
		{
			name:        "should return an error if the value is fully unquoted",
			fixtureName: "fully_unquoted_value.env",
			expectedErr: newUnquotedValueError(2),
		},
		{
			name:        "should return an error if the value has unterminated quotes",
			fixtureName: "unterminated_quotes.env",
			expectedErr: newUnterminatedQuoteError(2),
		},
		{
			name:        "should return an error if the value is empty",
			fixtureName: "empty_value.env",
			expectedErr: newEmptyValueError(2),
		},
		{
			name:        "should return an error if the key is lowercase",
			fixtureName: "lowercase_key.env",
			expectedErr: newLowercaseKeyError(2),
		},
		{
			name:        "should return an error if the key is mixed case",
			fixtureName: "mixed_case_key.env",
			expectedErr: newLowercaseKeyError(2),
		},
		{
			name:        "should return an error if the key is not alphanumeric",
			fixtureName: "non_alphanumeric_key.env",
			expectedErr: newInvalidKeyCharsError(2),
		},
		{
			name:        "should return an error if the key starts with a number",
			fixtureName: "key_starts_with_num.env",
			expectedErr: newNumberKeyError(2),
		},
		{
			name:        "should return an error if the key contains any numbers",
			fixtureName: "key_contains_numbers.env",
			expectedErr: newNumberKeyError(2),
		},
		{
			name:        "should return an error if the key starts with an underscore",
			fixtureName: "key_starts_with_underscore.env",
			expectedErr: newKeyStartsWithUnderscoreError(2),
		},
		{
			name:        "should return an error if the key has a trailing underscore",
			fixtureName: "trailing_underscore.env",
			expectedErr: newTrailingUnderscoreError(2),
		},
		{
			name:        "should return an error if line has an inline comment",
			fixtureName: "inline_comments.env",
			expectedErr: newInlineCommentError(2),
		},
		{
			name:        "should return an error if there are multiline values",
			fixtureName: "multiline_value.env",
			expectedErr: newMultilineValueError(2),
		},
		{
			name:        "should parse correctly if value contains a new line escape",
			fixtureName: "new_line.env",
			expectedParsed: map[string]string{
				"FIRST_KEY":  "some\nrandom\nvalue",
				"SECOND_KEY": "normal value",
			},
		},
		{
			name:        "should parse correctly if value contains a backslash escape",
			fixtureName: "slash_escape.env",
			expectedParsed: map[string]string{
				"FIRST_KEY":  `this is a \ character`,
				"SECOND_KEY": "normal value",
			},
		},
		{
			name:        "should parse correctly if value contains a double quotes escape",
			fixtureName: "double_quotes_escape.env",
			expectedParsed: map[string]string{
				"FIRST_KEY":  `He said: "I'll be back by tomorrow".`,
				"SECOND_KEY": "normal value",
			},
		},
		{
			name:        "should parse correctly if value contains single quotes",
			fixtureName: "single_quotes.env",
			expectedParsed: map[string]string{
				"FIRST_KEY": `you need to pass the 'verbose' argument`,
			},
		},
		{
			name:        "should parse correctly a valid env file",
			fixtureName: "example.env",
			expectedParsed: map[string]string{
				"APP_NAME":          "MyApp",
				"VERSION":           "1.0.0",
				"DEBUG":             "true",
				"ENVIRONMENT":       "production",
				"DATABASE_HOST":     "localhost",
				"DATABASE_PORT":     "5432",
				"DATABASE_NAME":     "my_database",
				"DATABASE_USER":     "db_user",
				"DATABASE_PASSWORD": "secure_password",
				"LOG_FILE":          "app_log.txt",
				"LOG_LEVEL":         "info",
				"API_URL":           "https://api.example.com",
				"API_KEY":           "your_api_key_here",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			var buffer bytes.Buffer

			fp := filepath.Join("./../../testdata/parser", tc.fixtureName)

			file, err := os.Open(fp)
			require.NoError(err, "Failed to open fixture.")

			_, err = io.Copy(&buffer, file)
			require.NoError(err, "Failed to copy fixture data into the buffer.")

			err = file.Close()
			assert.NoError(err, "Could not close expected buffer file.")

			parsed, err := Parse(buffer)

			if tc.expectedErr == nil {
				require.NoError(err, "An unexpected error ocurred")
				assert.Equal(parsed, tc.expectedParsed)
				assert.Nil(err)
			} else {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

				assert.Nil(parsed)
				assert.True(
					castErr.Equal(tc.expectedErr),
					"The actual error does not match the expected one. Actual: %v, Expected: %v",
					castErr,
					tc.expectedErr,
				)
			}
		})
	}
}
