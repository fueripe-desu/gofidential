package errors

const (
	// Occurs when the assignment operator is missing.
	MissingAssignmentCode string = "MISSING_ASSIGNMENT"

	// Occurs when the key is missing in a key-value pair in a .env file.
	MissingKeyCode string = "MISSING_KEY"

	// Occurs when the value is missing in a key-value pair in a .env file.
	MissingValueCode string = "MISSING_VALUE"

	// Occurs when the value inside the double quotes is empty..
	EmptyValueCode string = "EMPTY_VALUE"

	// Occurs when a value is quoted using single quotes in a .env file.
	SingleQuotesCode string = "SINGLE_QUOTES"

	// Occurs when a value is unquoted in a .env file.
	UnquotedValueCode string = "UNQUOTED_VALUE"

	// Occurs when the quotes around a value are not terminated correctly in a .env file.
	UnterminatedQuoteCode string = "UNTERMINATED_QUOTE"

	// Occurs when the assignment operator (=), as in KEY=VALUE, has spaces around it.
	SpacedSeparatorCode string = "SPACED_SEPARATOR"

	// Occurs when a .env file has a leading space before the key in one of its lines.
	LeadingSpaceCode string = "LEADING_SPACE"

	// Occurs when a .env file has an unallowed escape character.
	UnallowedEscapeCode string = "UNALLOWED_ESCAPE"

	// Occurs when a quote is escaped incorrectly resulting in traling text after the value.
	UnescapedQuoteCode string = "UNESCAPED_QUOTE"

	// Occurs when a key contains lowercase characters in a .env file.
	LowercaseKeyCode string = "LOWERCASE_KEY"

	// Occurs when a key starts with a number in a .env file.
	KeyStartsWithNumCode string = "KEY_STARTS_WITH_NUM"

	// Occurs when a key starts with an underscore in a .env file.
	KeyStartsWithUnderscoreCode string = "KEY_STARTS_WITH_UNDERSCORE"

	// Occurs when a key contains invalid characters in a .env file.
	InvalidKeyCharsCode string = "INVALID_KEY_CHARS"

	// Occurs when a .env file has an inline comment (which is a comment in the same line as a key).
	InlineCommentCode string = "INLINE_COMMENT"

	// Occurs when multiline values (which are not allowed) are used in a .env file.
	MultilineValueCode string = "MULTILINE_VALUE"
)
