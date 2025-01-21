package parser

const (
	// The PARSER_MISSING_ASSIGNMENT error occurs when a key-value pair in the .env file is
	// missing the assignment operator (=), which separates the key from its value.
	MissingAssignment string = "PARSER_MISSING_ASSIGNMENT"

	// The PARSER_MISSING_KEY error occurs when a key-value pair in the .env file is missing
	// a key—the identifier that appears before the assignment operator (=).
	MissingKey string = "PARSER_MISSING_KEY"

	// The PARSER_MISSING_VALUE error occurs when a key-value pair in the .env file is missing
	// the value—the part that is assigned to the key and follows the assignment operator (=).
	MissingValue string = "PARSER_MISSING_VALUE"

	// The PARSER_EMPTY_VALUE error occurs when a key-value pair in the .env file has an empty
	// value enclosed in double quotes.
	EmptyValue string = "PARSER_EMPTY_VALUE"

	// The PARSER_SINGLE_QUOTED_VALUE error occurs when a key-value pair in the .env file has
	// its value enclosed in single quotes, which is not allowed.
	SingleQuotedValue string = "PARSER_SINGLE_QUOTED_VALUE"

	// The PARSER_UNQUOTED_VALUE error occurs when a key-value pair in the .env file does not
	// have the value enclosed in double quotes, which is not allowed.
	UnquotedValue string = "PARSER_UNQUOTED_VALUE"

	// The PARSER_UNTERMINATED_QUOTES error occurs when a key-value pair in the .env file starts
	// with an opening double quote but is missing the closing double quote for the value.
	UnterminatedQuotes string = "PARSER_UNTERMINATED_QUOTES"

	// The PARSER_SPACED_SEPARATOR error occurs when a key-value pair in the .env file contains
	// spaces around the assignment operator (=).
	SpacedSeparator string = "PARSER_SPACED_SEPARATOR"

	// The PARSER_LEADING_SPACE error occurs when a key-value pair in the .env file contains
	// leading spaces before the key.
	LeadingSpace string = "PARSER_LEADING_SPACE"

	// The PARSER_UNALLOWED_ESCAPE error occurs when a key-value pair in the .env file contains
	// an escape sequence in the value that is not permitted. The parser strictly allows only the
	// following escape characters:
	//
	//	- \n for a newline,
	//	- \" for a double quote,
	//	- \\ for a backslash.
	//
	// All other escape sequences are considered invalid.
	UnallowedEscape string = "PARSER_UNALLOWED_ESCAPE"

	// The PARSER_UNESCAPED_QUOTE_CHAR error occurs when a key-value pair in the .env file contains
	// multiple unescaped double quotes within the value. This typically happens when the user forgets
	// to escape the double quote character (\") that appears as part of the content.
	UnescapedQuoteChar string = "PARSER_UNESCAPED_QUOTE"

	// The PARSER_LOWERCASE_KEY error occurs when a key in the .env file contains one or more lowercase
	// characters. The parser enforces strict rules requiring all keys to be written entirely in uppercase.
	LowercaseKey string = "PARSER_LOWERCASE_KEY"

	// The PARSER_NUMERIC_KEY_CHARS error occurs when a key in the .env file contains one or more numeric
	// characters. The parser enforces strict rules requiring all keys to be written only using uppercase
	// letters and underscores, without any numeric characters.
	NumericKeyChars string = "PARSER_NUMERIC_KEY_CHARS"

	// The PARSER_LEADING_UNDERSCORE error occurs when a key in the .env file starts with a leading underscore
	// (_).
	LeadingUnderscore string = "PARSER_KEY_STARTS_WITH_UNDERSCORE"

	// The PARSER_TRAILING_UNDERSCORE error occurs when a key in the .env file ends with a trailing underscore
	// (_).
	TrailingUnderscore string = "PARSER_TRAILING_UNDERSCORE"

	// The PARSER_INVALID_KEY_CHARS error occurs when a key in the .env file contains characters other than
	// uppercase letters and underscores. The parser strictly enforces these rules, rejecting keys with numbers,
	// lowercase letters, or special characters.
	InvalidKeyChars string = "PARSER_INVALID_KEY_CHARS"

	// The PARSER_INLINE_COMMENT error occurs when a key-value pair in the .env file includes an inline comment
	// on the same line as the definition. The parser enforces a strict rule prohibiting inline comments.
	InlineComment string = "PARSER_INLINE_COMMENT"

	// The PARSER_MULTILINE_VALUE error occurs when a key-value pair in the .env file includes a value that spans
	// multiple lines. The parser strictly prohibits multiline values to maintain consistency and simplicity in
	// the .env file format.
	//
	// If you need to represent a line break within a value, use the \n escape character instead.
	MultilineValue string = "PARSER_MULTILINE_VALUE"
)
