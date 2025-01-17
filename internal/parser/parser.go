package parser

import (
	"bytes"
	"strings"
)

func Parse(buffer bytes.Buffer) (map[string]string, error) {
	// Converts Windows-style newline to only a newline escape character.
	byteData := bytes.Replace(buffer.Bytes(), []byte("\r\n"), []byte("\n"), -1)

	// Counters for the line numbers and column number when reading the .env file.
	lineNumber := 1
	ColmNumber := 1

	// Buffers to store the read bytes so later they can be turned into a string
	// and then cleared.
	var keyBuffer strings.Builder
	var valueBuffer strings.Builder

	// Set to true when the assignment operator '=' is found in the line.
	foundAssignment := false

	// Set to true when a backslash '\' is found in the line.
	foundEscape := false

	// Set to true when a closing double quote for the value is found.
	isAfterValue := false

	// Set to true when the first character of the key has not been found yet.
	isBeforeKey := true

	// Set to true when the first non-space character found in the line is '#'.
	isComment := false

	// Set to true when the reading the content inside the value's double quotes.
	insideQuotes := false

	// Set to true when the last character is a space and is not inside the value's double quotes.
	lastSpace := false

	// The final map to be returned with all key-value pairs.
	parsed := map[string]string{}

	for _, b := range byteData {
		if !insideQuotes && b == ' ' {
			lastSpace = true
		} else if isBeforeKey && b == '#' {
			if lastSpace {
				return nil, newLeadingSpaceError(lineNumber)
			}

			isComment = true
		} else if isBeforeKey && b == '=' {
			return nil, newMissingKeyError(lineNumber)
		} else if b == '\n' {
			if isComment {
				lineNumber++
				ColmNumber = 1

				foundAssignment = false
				foundEscape = false
				isAfterValue = false
				insideQuotes = false
				lastSpace = false
				isBeforeKey = true
				isComment = false

				continue
			}

			if isBeforeKey {
				lineNumber++
				ColmNumber = 1

				foundAssignment = false
				foundEscape = false
				isAfterValue = false
				insideQuotes = false
				lastSpace = false
				isBeforeKey = true
				isComment = false

				continue
			}

			if foundEscape {
				return nil, newMultilineValueError(lineNumber)
			}

			if insideQuotes {
				return nil, newUnterminatedQuoteError(lineNumber)
			}

			if !isBeforeKey && foundAssignment && !isAfterValue {
				return nil, newMissingValueError(lineNumber)
			}

			foundAssignment = false
			foundEscape = false
			isAfterValue = false
			insideQuotes = false
			lastSpace = false
			isBeforeKey = true
			isComment = false

			if valueBuffer.Len() == 0 {
				return nil, newEmptyValueError(lineNumber)
			}

			parsed[keyBuffer.String()] = valueBuffer.String()

			keyBuffer.Reset()
			valueBuffer.Reset()

			lineNumber++
			ColmNumber = 1
			continue
		} else if isComment {
			continue
		} else if isAfterValue {
			if b == ' ' {
				continue
			} else if b == '#' {
				return nil, newInlineCommentError(lineNumber)
			} else {
				return nil, newUnescapedQuoteError(lineNumber)
			}
		} else if foundEscape {
			switch b {
			case 'n':
				valueBuffer.WriteByte('\n')
			case '"':
				valueBuffer.WriteByte('"')
			case '\\':
				valueBuffer.WriteByte('\\')
			default:
				return nil, newUnallowedEscapeError(lineNumber, string(b))
			}
			foundEscape = false
		} else if insideQuotes {
			if b == '\\' {
				foundEscape = true
				continue
			}

			if b == '"' {
				insideQuotes = false
				isAfterValue = true
				continue
			}

			valueBuffer.WriteByte(b)
		} else if foundAssignment {
			if lastSpace {
				return nil, newSpacedSeparatorError(lineNumber)
			}

			if b == '\'' {
				return nil, newSingleQuotesError(lineNumber)
			}

			if b != '"' {
				return nil, newUnquotedValueError(lineNumber)
			}

			insideQuotes = true
		} else if b == '=' && !foundAssignment {
			if lastSpace {
				return nil, newSpacedSeparatorError(lineNumber)
			}

			foundAssignment = true
		} else if !foundAssignment {
			if lastSpace {
				if keyBuffer.Len() > 0 {
					return nil, newInvalidKeyCharsError(lineNumber)
				}

				return nil, newLeadingSpaceError(lineNumber)
			}

			if keyBuffer.Len() == 0 && isUnderscore(b) {
				return nil, newKeyStartsWithUnderscoreError(lineNumber)
			}

			if keyBuffer.Len() == 0 && isNumber(b) {
				return nil, newKeyStartsWithNumError(lineNumber)
			}

			if isLowercase(b) {
				return nil, newLowercaseKeyError(lineNumber)
			}

			if b == '"' {
				return nil, newMissingAssignmentError(lineNumber)
			}

			if !isUppercase(b) && !isUnderscore(b) {
				return nil, newInvalidKeyCharsError(lineNumber)
			}

			isBeforeKey = false
			keyBuffer.WriteByte(b)
		} else {
		}

		ColmNumber++
	}

	return parsed, nil
}
