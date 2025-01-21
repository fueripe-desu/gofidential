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
