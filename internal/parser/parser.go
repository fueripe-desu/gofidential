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

	for _, b := range p.data {
		if p.state == BeforeKey {
			if b == ' ' {
				p.hasPrecedingSpace = true
			} else {
				if p.isEndOfLine(b) {
					p.nextLine()
					continue
				}

				if b == '=' {
					return nil, newMissingKeyError(p.lineNumber)
				}

				if p.hasPrecedingSpace {
					return nil, newLeadingSpaceError(p.lineNumber)
				}

				if b == '#' {
					p.state = InsideComment
				} else {
					p.state = ParsingKey
					if err := p.handleParsingKey(b); err != nil {
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
			if err := p.handleParsingKey(b); err != nil {
				return nil, err
			}
		} else if p.state == FoundAssignment {
			if p.isEndOfLine(b) {
				return nil, newMissingValueError(p.lineNumber)
			}

			if b == ' ' {
				return nil, newSpacedSeparatorError(p.lineNumber)
			}

			if b == '\'' {
				return nil, newSingleQuotesError(p.lineNumber)
			}

			if b != '"' {
				return nil, newUnquotedValueError(p.lineNumber)
			}

			p.state = InsideQuotes
		} else if p.state == InsideQuotes {
			if !p.isEscape && p.isEndOfLine(b) {
				return nil, newUnterminatedQuoteError(p.lineNumber)
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
					return nil, newMultilineValueError(p.lineNumber)
				default:
					return nil, newUnallowedEscapeError(p.lineNumber, string(b))
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
					return nil, newEmptyValueError(p.lineNumber)
				}

				parsed[p.keyBuffer.String()] = p.valueBuffer.String()

				p.nextLine()
				continue
			}

			if b == '#' {
				return nil, newInlineCommentError(p.lineNumber)
			} else if b != ' ' {
				return nil, newUnescapedQuoteError(p.lineNumber)
			}
		}

		p.nextColumn()
	}

	return parsed, nil
}

func (p *parser) handleParsingKey(b byte) error {
	if b == ' ' {
		p.hasPrecedingSpace = true
		p.nextColumn()
		return nil
	}

	if p.keyBuffer.Len() == 0 {
		if b == '_' {
			return newKeyStartsWithUnderscoreError(p.lineNumber)
		}
	}

	if b >= 'a' && b <= 'z' {
		return newLowercaseKeyError(p.lineNumber)
	}

	if b >= '0' && b <= '9' {
		return newNumberKeyError(p.lineNumber)
	}

	if b == '"' {
		return newMissingAssignmentError(p.lineNumber)
	}

	if b == '=' {
		if p.hasPrecedingSpace {
			return newSpacedSeparatorError(p.lineNumber)
		}

		p.state = FoundAssignment
	} else {
		if p.hasPrecedingSpace {
			return newInvalidKeyCharsError(p.lineNumber)
		}

		if !(b >= 'A' && b <= 'Z' || b == '_') {
			return newInvalidKeyCharsError(p.lineNumber)
		}

		p.keyBuffer.WriteByte(b)
	}

	return nil
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
