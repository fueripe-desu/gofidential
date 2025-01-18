package parser

import (
	"runtime/debug"
	"strconv"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
)

const moduleIssuer string = "GoFidential/Parser"

func newMissingAssignmentError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.MissingAssignmentCode,
		Message:    "Missing assignment operator.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if all lines have an assignment operator.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newMissingKeyError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.MissingKeyCode,
		Message:    "Key is missing.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the key is in the same line as the value.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newMissingValueError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.MissingValueCode,
		Message:    "Value is missing.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Add a valid double-quoted value for the key.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newEmptyValueError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.EmptyValueCode,
		Message:    "Value must not be empty.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Try adding content inside the value double quotes.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newSingleQuotesError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.SingleQuotesCode,
		Message:    "Single quotes are not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Replace single quotes for double quotes.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newUnquotedValueError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.UnquotedValueCode,
		Message:    "Value is not double quoted.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Place double quotes around the unquoted value.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newUnterminatedQuoteError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.UnterminatedQuoteCode,
		Message:    "Unterminated quote found.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if all values are opened and closed with double quotes correctly, or if there is an unescaped double quotes in their values.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newSpacedSeparatorError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.SpacedSeparatorCode,
		Message:    "Assignment operator must not have spaces around it.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove any trailing or leading spaces between the assignment operator and the key or value it operates on.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newLeadingSpaceError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.LeadingSpaceCode,
		Message:    "Leading spaces before key are not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove leading spaces before the keys in the .env file.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newUnallowedEscapeError(lineNumber int, invalidEscape string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.UnallowedEscapeCode,
		Message:    "Escape character is not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove the unallowed escape character.",
		Details: map[string]string{
			"escape_char": invalidEscape,
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newLowercaseKeyError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.LowercaseKeyCode,
		Message:    "Key must not be lowercase.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Replace lowercase characters by uppercase characters.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newNumberKeyError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.NumberKeyCode,
		Message:    "Key must not contain numbers.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove any numbers from the key.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newLeadingUnderscoreError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.LeadingUnderscoreCode,
		Message:    "Key must not start with a underscore.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove underscores in the beginning of the key name.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newTrailingUnderscoreError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.TrailingUnderscoreCode,
		Message:    "Key must not end with an underscore.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove underscores in the end of the key name.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newInvalidKeyCharsError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.InvalidKeyCharsCode,
		Message:    "Key is not alphanumeric.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Rename the key so it has only numbers (0-9) and uppercase letters (A-Z).",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newInlineCommentError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.InlineCommentCode,
		Message:    "Inline comments are not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Move inline comments to their own lines, instead of placing them in the same line as the key-value pair, in the .env file.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newUnescapedQuoteError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.UnescapedQuoteCode,
		Message:    "Unescaped quote found.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Maybe you forgot to escape the double quotes character?",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}

func newMultilineValueError(lineNumber int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.MultilineValueCode,
		Message:    "Multiline values are not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Try replacing multiline values by a single-line value, or use the \\n escape character to indicate a line break.",
		Details: map[string]string{
			"line_number": strconv.Itoa(lineNumber),
		},
	}
}
