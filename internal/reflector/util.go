package reflector

import (
	"strings"
	"unicode"
)

// The [reflector.upperToPascal] helper function converts a string from
// uppercase snake case (used for .env keys) to Pascal case, which is
// the convention for exported struct fields.
//
// Parameters:
//   - v (string): The uppercase snake case string to be converted.
//
// Returns:
//   - string: The converted string in Pascal case.
func upperToPascal(v string) string {
	var result strings.Builder

	foundUnderscore := false

	for i, r := range v {
		if r == '_' {
			foundUnderscore = true
			continue
		}

		if i == 0 || foundUnderscore {
			result.WriteRune(r)
			foundUnderscore = false
		} else {
			lower := unicode.ToLower(r)
			result.WriteRune(lower)
		}
	}

	return result.String()
}

// The [reflector.pascalToUpper] helper function converts a string from
// Pascal case (used for exported struct fields) to uppercase snake case,
// which is the convention for .env keys.
//
// Parameters:
//   - v (string): The Pascal case string to be converted.
//
// Returns:
//   - string: The converted string in uppercase snake case.
func pascalToUpper(v string) string {
	var result strings.Builder
	// Flag to track if we need to insert an underscore before digits
	inDigitSequence := false

	for i, r := range v {
		// Check if the character is a digit
		if unicode.IsDigit(r) {
			// If it's the first digit in the sequence, add an underscore before it
			if !inDigitSequence && i > 0 {
				result.WriteRune('_')
				inDigitSequence = true
			}
			result.WriteRune(unicode.ToUpper(r))
			continue
		}

		// Insert an underscore before uppercase letters (except the first character)
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}

		// Write the uppercase version of the character
		result.WriteRune(unicode.ToUpper(r))

		// Reset the digit sequence flag after a non-digit
		inDigitSequence = false
	}

	return result.String()
}
