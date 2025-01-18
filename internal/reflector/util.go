package reflector

import (
	"strings"
	"unicode"
)

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
