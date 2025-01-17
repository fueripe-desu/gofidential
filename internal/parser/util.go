package parser

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isUnderscore(b byte) bool {
	return b == '_'
}

func isLowercase(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isUppercase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
