package str

import "strings"

// Normalize removes leading and trailing spaces from a string and converts it to lowercase.
func Normalize(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

// Reverse reverses a string.
func Reverse(value string) string {
	var runes = []rune(value)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
