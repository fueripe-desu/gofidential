package parser

type parserState int

const (
	BeforeKey parserState = iota
	ParsingKey
	FoundAssignment
	ParsingValue
	InsideQuotes
	AfterValue
	InsideComment
)
