package parser

import (
	"bytes"
	"strings"
)

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

func (p *parser) isEndOfLine(b byte) bool {
	return b == '\n'
}

func (p *parser) nextColumn() {
	p.colmNumber++
}

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

func newParser(buffer bytes.Buffer) *parser {
	// Converts Windows-style newline to only a newline escape character.
	byteData := bytes.Replace(buffer.Bytes(), []byte("\r\n"), []byte("\n"), -1)

	return &parser{
		data:              byteData,
		lineNumber:        1,
		colmNumber:        1,
		hasPrecedingSpace: false,
		state:             BeforeKey,
	}
}

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
