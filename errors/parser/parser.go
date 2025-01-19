package parser

const (
	// Occurs when the assignment operator is missing.
	MissingAssignment string = "MISSING_ASSIGNMENT"

	// Occurs when the key is missing in a key-value pair in a .env file.
	MissingKey string = "MISSING_KEY"

	// Occurs when the value is missing in a key-value pair in a .env file.
	MissingValue string = "MISSING_VALUE"

	// Occurs when the value inside the double quotes is empty..
	EmptyValue string = "EMPTY_VALUE"

	// Occurs when a value is quoted using single quotes in a .env file.
	SingleQuotedValue string = "SINGLE_QUOTES"

	// Occurs when a value is unquoted in a .env file.
	UnquotedValue string = "UNQUOTED_VALUE"

	// Occurs when the quotes around a value are not terminated correctly in a .env file.
	UnterminatedQuotes string = "UNTERMINATED_QUOTE"

	// Occurs when the assignment operator (=), as in KEY=VALUE, has spaces around it.
	SpacedSeparator string = "SPACED_SEPARATOR"

	// Occurs when a .env file has a leading space before the key in one of its lines.
	LeadingSpace string = "LEADING_SPACE"

	// Occurs when a .env file has an unallowed escape character.
	UnallowedEscape string = "UNALLOWED_ESCAPE"

	// Occurs when a quote is escaped incorrectly resulting in traling text after the value.
	UnescapedQuoteChar string = "UNESCAPED_QUOTE"

	// Occurs when a key contains lowercase characters in a .env file.
	LowercaseKey string = "LOWERCASE_KEY"

	// Occurs when a key contains number characters (0-9) in a .env file.
	NumericKeyChars string = "NUMERIC_KEY"

	// Occurs when a key starts with an underscore in a .env file.
	LeadingUnderscore string = "KEY_STARTS_WITH_UNDERSCORE"

	// Occurs when a key has a trailing underscore in a .env file.
	TrailingUnderscore string = "TRAILING_UNDERSCORE"

	// Occurs when a key contains invalid characters in a .env file.
	InvalidKeyChars string = "INVALID_KEY_CHARS"

	// Occurs when a .env file has an inline comment (which is a comment in the same line as a key).
	InlineComment string = "INLINE_COMMENT"

	// Occurs when multiline values (which are not allowed) are used in a .env file.
	MultilineValue string = "MULTILINE_VALUE"
)
