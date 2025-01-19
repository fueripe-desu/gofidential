package parser

import (
	"bytes"
	"strings"
)

// The [parser.parser] struct encapsulates the state and functions required to parse
// a .env file's byte buffer into a map of key-value pairs.
//
// This struct maintains the parser's state, buffers, and position tracking during the
// parsing process. It uses a finite state machine (FSM) to handle the parsing logic.
//
// Fields:
//   - data ([]byte): A normalized copy of the byte buffer, with line endings unified.
//   - lineNumber (int): The current line number the parser is processing in the buffer.
//   - colmNumber (int): The current column number the parser pointer is positioned at.
//   - keyBuffer (strings.Builder): A buffer for constructing the key from the parsed bytes.
//   - valueBuffer (strings.Builder): A buffer for constructing the value from the parsed bytes.
//   - hasPrecedingSpace (bool): Indicates if the previous character was a space. Used for error validation.
//   - isEscape (bool): Indicates if the previous character was a backslash. Used for escape handling and error checking.
//   - state (parser.parserState): Represents the current state of the parser's finite state machine.
type parser struct {
	data []byte

	lineNumber int
	colmNumber int

	keyBuffer   strings.Builder
	valueBuffer strings.Builder

	hasPrecedingSpace bool
	isEscape          bool

	state parserState
}

// The [parser.parser.Parse] method processes the byte buffer initialized in the parser
// instance and converts it into a map of key-value pairs representing the contents of a .env file.
//
// Returns:
//   - map[string]string: A map containing the parsed key-value pairs from the .env file.
//   - error: An error indicating why the parsing process failed.
func (p *parser) Parse() (map[string]string, error) {
	parsed := map[string]string{}

	for i, b := range p.data {
		if p.state == BeforeKey {
			if b == ' ' {
				p.hasPrecedingSpace = true
			} else {
				if p.isEndOfLine(b) {
					p.nextLine()
					continue
				}

				if b == '=' {
					return nil, newMissingKeyError(p.lineNumber, p.colmNumber)
				}

				if p.hasPrecedingSpace {
					return nil, newLeadingSpaceError(p.lineNumber)
				}

				if b == '#' {
					p.state = InsideComment
				} else {
					p.state = ParsingKey
					if err := p.handleParsingKey(b, i); err != nil {
						return nil, err
					}
				}
			}
		} else if p.state == InsideComment {
			if p.isEndOfLine(b) {
				p.nextLine()
				continue
			}
		} else if p.state == ParsingKey {
			if err := p.handleParsingKey(b, i); err != nil {
				return nil, err
			}
		} else if p.state == FoundAssignment {
			if p.isEndOfLine(b) {
				return nil, newMissingValueError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}

			if b == ' ' {
				return nil, newSpacedSeparatorError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}

			if b == '\'' {
				return nil, newSingleQuotedValueError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}

			if b != '"' {
				return nil, newUnquotedValueError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}

			p.state = InsideQuotes
		} else if p.state == InsideQuotes {
			if !p.isEscape && p.isEndOfLine(b) {
				return nil, newUnterminatedQuotesError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}

			if p.isEscape {
				switch b {
				case 'n':
					p.valueBuffer.WriteByte('\n')
				case '"':
					p.valueBuffer.WriteByte('"')
				case '\\':
					p.valueBuffer.WriteByte('\\')
				case '\n':
					return nil, newMultilineValueError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
				default:
					return nil, newUnallowedEscapeError(p.lineNumber, p.colmNumber, p.keyBuffer.String(), string(b))
				}

				p.isEscape = false
			} else if b == '\\' {
				p.isEscape = true
			} else if b == '"' {
				p.state = AfterValue
			} else {
				p.valueBuffer.WriteByte(b)
			}
		} else if p.state == AfterValue {
			if p.isEndOfLine(b) {
				if p.valueBuffer.Len() == 0 {
					// Subtracts one from the column, so the column number in the error points exactly to the
					// closing quote instead of the new line escape.
					return nil, newEmptyValueError(p.lineNumber, p.colmNumber-1, p.keyBuffer.String())
				}

				key := p.keyBuffer.String()

				// The key ends with an underscore
				if key[len(key)-1] == '_' {
					return nil, newTrailingUnderscoreError(p.lineNumber, p.keyBuffer.Len(), p.keyBuffer.String())
				}

				parsed[key] = p.valueBuffer.String()

				p.nextLine()
				continue
			}

			if b == '#' {
				return nil, newInlineCommentError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			} else if b != ' ' {
				return nil, newUnescapedQuoteCharError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
			}
		}

		p.nextColumn()
	}

	return parsed, nil
}

// The [parser.parser.handleParsingKey] method is an internal parser function
// responsible for handling the "ParsingKey" state of the parser's finite
// state machine (FSM). It is implemented as a separate function to allow
// reuse across multiple parts of the system.
//
// Parameters:
//   - b (byte): The current byte being processed from the buffer.
//   - globalIndex (int): The zero-based index of the byte in the entire
//     buffer, independent of line and column.
//
// Returns:
//   - error: An error explaining why processing the current key byte failed.
func (p *parser) handleParsingKey(b byte, globalIndex int) error {
	if b == ' ' {
		p.hasPrecedingSpace = true
		return nil
	}

	if p.keyBuffer.Len() == 0 {
		if b == '_' {
			errBuffer := p.getKeyErrorBuffer(globalIndex)
			return newLeadingUnderscoreError(p.lineNumber, p.colmNumber, errBuffer.String())
		}
	}

	if b >= 'a' && b <= 'z' {
		errBuffer := p.getKeyErrorBuffer(globalIndex)
		return newLowercaseKeyError(p.lineNumber, p.colmNumber, errBuffer.String())
	}

	if b >= '0' && b <= '9' {
		errBuffer := p.getKeyErrorBuffer(globalIndex)
		return newNumericKeyCharsError(p.lineNumber, p.colmNumber, errBuffer.String())
	}

	if b == '"' {
		return newMissingAssignmentError(p.lineNumber, p.colmNumber, p.keyBuffer.String())
	}

	if b == '=' {
		if p.hasPrecedingSpace {
			// Subtracts one from the column, so the column number in the error points exactly to the
			// space before the assignment operator, without this subtraction, the column number will
			// point to the assignment operator, rather than the preceding space.
			return newSpacedSeparatorError(p.lineNumber, p.colmNumber-1, p.keyBuffer.String())
		}

		p.state = FoundAssignment
	} else {
		if p.hasPrecedingSpace {
			errBuffer := p.getKeyErrorBuffer(globalIndex)
			return newInvalidKeyCharsError(p.lineNumber, p.colmNumber, errBuffer.String())
		}

		if !(b >= 'A' && b <= 'Z' || b == '_') {
			errBuffer := p.getKeyErrorBuffer(globalIndex)
			return newInvalidKeyCharsError(p.lineNumber, p.colmNumber, errBuffer.String())
		}

		p.keyBuffer.WriteByte(b)
	}

	return nil
}

// The [parser.parser.getKeyErrorBuffer] method retrieves the remaining portion of a key
// when an error occurs during key parsing. Since the parser processes one byte at a time,
// an error interrupts parsing, leaving the pointer at its current position without advancing
// to the end of the key. This method ensures the full key can be extracted for use in error messages.
//
// Parameters:
//   - globalIndex (int): The zero-based index of the current byte in the buffer, independent
//     of line and column.
//
// Returns:
//   - strings.Builder: A buffer containing the full key string for error reporting purposes.
func (p *parser) getKeyErrorBuffer(globalIndex int) strings.Builder {
	var errBuffer strings.Builder

	errBuffer.WriteString(p.keyBuffer.String())

	for ei := globalIndex; ei < len(p.data); ei++ {
		b := p.data[ei]

		if b == '=' || b == '\n' || b == '"' {
			break
		}

		errBuffer.WriteByte(b)
	}

	return errBuffer
}

// The [parser.parser.isEndOfLine] method is an internal helper function that checks
// whether the given byte represents a newline character. This improves code readability
// by abstracting the newline check logic.
//
// Parameters:
//   - b (byte): The byte to evaluate.
//
// Returns:
//   - bool: True if the input byte is a newline character; false otherwise.
func (p *parser) isEndOfLine(b byte) bool {
	return b == '\n'
}

// The [parser.parser.nextColumn] method is an internal helper function that increments
// the [parser.parser.colmNumber] field. This field tracks the current column number
// in the parsing process for accurate error reporting and state management.
func (p *parser) nextColumn() {
	p.colmNumber++
}

// The [parser.parser.nextLine] method is an internal helper function that resets
// the parser state for a new line. It clears the key and value buffers, restarts
// the column counter, and increments the line number, preparing the parser for
// the next line of input.
func (p *parser) nextLine() {
	// Resets initial state.
	p.state = BeforeKey
	p.hasPrecedingSpace = false
	p.isEscape = false

	// Clear all buffers.
	p.keyBuffer.Reset()
	p.valueBuffer.Reset()

	// Sets index to the first character of the next line.
	p.lineNumber++
	p.colmNumber = 1
}

// The [parser.Parse] function parses the contents of a .env file from a byte buffer
// into a map of key-value pairs.
//
// Parameters:
//   - buffer (bytes.Buffer): The buffer containing the bytes of the .env file.
//
// Returns:
//   - map[string]string: A map containing the parsed key-value pairs from the .env file.
//   - error: An error indicating why the parsing process failed.
//
// Notes:
//   - This function normalizes new line characters in the buffer before storing them.
func Parse(buffer bytes.Buffer) (map[string]string, error) {
	// Converts Windows-style newline to only a newline escape character.
	byteData := bytes.Replace(buffer.Bytes(), []byte("\r\n"), []byte("\n"), -1)

	parser := &parser{
		data:              byteData,
		lineNumber:        1,
		colmNumber:        1,
		hasPrecedingSpace: false,
		state:             BeforeKey,
	}

	return parser.Parse()
}
