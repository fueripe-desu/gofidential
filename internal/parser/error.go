package parser

import (
	"fmt"

	"github.com/fueripe-desu/gofidential/errors"
	errorCode "github.com/fueripe-desu/gofidential/errors/parser"
)

// The [parser.newMissingAssignmentError] function creates an error with the
// code "MissingAssignment".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newMissingAssignmentError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MissingAssignment,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' is missing the assignment operator.", line, col, key),
		Hint:    "Add an assignment operator between the key and the value.",
	}
}

// The [parser.newMissingKeyError] function creates an error with the
// code "MissingKey".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newMissingKeyError(line int, col int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MissingKey,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: Key is missing.", line, col),
		Hint:    "Ensure that the key is placed right before the assignment operator (=).",
	}
}

// The [parser.newMissingValueError] function creates an error with the
// code "MissingValue".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newMissingValueError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MissingValue,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' is missing its value.", line, col, key),
		Hint:    "Ensure that the value is placed right after the assignment operator (=).",
	}
}

// The [parser.newEmptyValueError] function creates an error with the
// code "EmptyValue".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newEmptyValueError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.EmptyValue,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' has an empty value.", line, col, key),
		Hint:    "If the key isn't needed, consider removing it.",
	}
}

// The [parser.newSingleQuotedValueError] function creates an error with the
// code "SingleQuotedValue".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newSingleQuotedValueError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.SingleQuotedValue,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The value of the key '%s' is single quoted.", line, col, key),
		Hint:    "Use double quotes around the value instead of single quotes.",
	}
}

// The [parser.newUnquotedValueError] function creates an error with the
// code "UnquotedValue".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newUnquotedValueError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnquotedValue,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The value of the key '%s' is unquoted.", line, col, key),
		Hint:    "Use double quotes around the value instead of leaving it unquoted.",
	}
}

// The [parser.newUnterminatedQuotesError] function creates an error with the
// code "UnterminatedQuotes".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newUnterminatedQuotesError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnterminatedQuotes,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The value of the key '%s' is missing the closing double quotes.", line, col, key),
		Hint:    "Make sure to add a closing double quote at the end of the value.",
	}
}

// The [parser.newSpacedSeparatorError] function creates an error with the
// code "SpacedSeparator".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newSpacedSeparatorError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.SpacedSeparator,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The assignment operator of the key '%s' is spaced.", line, col, key),
		Hint:    "Remove any spaces before or after the assignment operator.",
	}
}

// The [parser.newLeadingSpaceError] function creates an error with the
// code "LeadingSpace".
//
// Parameters:
//   - line (int): The line where the error occured.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newLeadingSpaceError(line int) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.LeadingSpace,
		Message: fmt.Sprintf("[Line: %d, Column: 0]: Leading spaces before the key are not allowed.", line),
		Hint:    "Remove any leading spaces before the key.",
	}
}

// The [parser.newUnallowedEscapeError] function creates an error with the
// code "UnallowedEscape".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//   - escape (string): The unallowed escape that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newUnallowedEscapeError(line int, col int, key string, escape string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnallowedEscape,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains the invalid escape character '%s' in its value.", line, col, key, escape),
		Hint:    "Consider removing the invalid escape character or using an allowed one.",
	}
}

// The [parser.newLowercaseKeyError] function creates an error with the
// code "LowercaseKey".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newLowercaseKeyError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.LowercaseKey,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains lowercase characters, which are not allowed.", line, col, key),
		Hint:    "Please use only uppercase characters in the key.",
	}
}

// The [parser.newNumericKeyCharsError] function creates an error with the
// code "NumericKeyChars".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newNumericKeyCharsError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.NumericKeyChars,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains numeric characters, which are not allowed.", line, col, key),
		Hint:    "Remove all numbers from the key.",
	}
}

// The [parser.newLeadingUnderscoreError] function creates an error with the
// code "LeadingUnderscore".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newLeadingUnderscoreError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.LeadingUnderscore,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' has leading underscores, which are not allowed.", line, col, key),
		Hint:    "Remove underscores in the beginning of the key.",
	}
}

// The [parser.newTrailingUnderscoreError] function creates an error with the
// code "TrailingUnderscore".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newTrailingUnderscoreError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.TrailingUnderscore,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' has trailing underscores, which are not allowed.", line, col, key),
		Hint:    "Remove underscores in the end of the key.",
	}
}

// The [parser.newInvalidKeyCharsError] function creates an error with the
// code "InvalidKeyChars".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newInvalidKeyCharsError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidKeyChars,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains invalid characters.", line, col, key),
		Hint:    "Rename the key to contain only uppercase letters (A-Z) and underscores (_).",
	}
}

// The [parser.newInlineCommentError] function creates an error with the
// code "InlineComment".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newInlineCommentError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InlineComment,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains an inline comment.", line, col, key),
		Hint:    "Move comments to their own lines, instead of placing them in the same line as the key-value pair.",
	}
}

// The [parser.newUnescapedQuoteCharError] function creates an error with the
// code "UnescapedQuoteChar".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newUnescapedQuoteCharError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnescapedQuoteChar,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The value of the key '%s' contains unescaped double quotes.", line, col, key),
		Hint:    "Maybe you forgot to escape the double quotes character?",
	}
}

// The [parser.newMultilineValueError] function creates an error with the
// code "MultilineValue".
//
// Parameters:
//   - line (int): The line where the error occured.
//   - col (int): The column wherer the error occured.
//   - key (string): The key that caused the error.
//
// Returns:
//   - *errors.GofidentialError: The created error with the relevant details.
func newMultilineValueError(line int, col int, key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MultilineValue,
		Message: fmt.Sprintf("[Line: %d, Column: %d]: The key '%s' contains a multi-line value", line, col, key),
		Hint:    "Replace multi-line values by a single-line value, and consider using the \\n escape character for line breaks.",
	}
}
