package parser

// The [parser.parserState] type represents the various states in the
// parser's finite state machine (FSM) during the parsing of a .env file.
type parserState int

const (
	// [parser.parserState.BeforeKey] is the initial state, responsible
	// for handling the bytes before the first character of the key.
	BeforeKey parserState = iota

	// [parser.parserState.ParsingKey] is the state where the parser has
	// found the first character of the key and continues to handle all
	// subsequent key bytes until the assignment operator is encountered.
	ParsingKey

	// [parser.parserState.FoundAssignment] is the state reached when the
	// assignment operator is found. It handles the bytes between the
	// assignment operator and the opening double quote of the value.
	FoundAssignment

	// [parser.parserState.InsideQuotes] is the state where the parser is
	// inside the value portion, handling the bytes between the opening and
	// closing double quotes that make up the value.
	InsideQuotes

	// [parser.parserState.AfterValue] is the state after the closing double
	// quote is found. It handles the bytes following the value, such as
	// trailing spaces, new line characters, or comments, and manages
	// any invalid content after the value.
	AfterValue

	// [parser.parserState.InsideComment] is a special state that occurs when
	// a line begins with a comment mark (#). This state indicates that the
	// entire line should be ignored during parsing.
	InsideComment
)
