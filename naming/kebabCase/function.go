package gpkebabcase

import (
	"strings"
	"unicode"
)

func FromWords(value string) string {

	words := strings.FieldsFunc(value, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
}

func ToWords(value string) string {

	words := strings.Split(value, "-")

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	if len(words) > 0 && words[0] != "" {
		runes := []rune(words[0])
		runes[0] = unicode.ToUpper(runes[0])
		words[0] = string(runes)
	}

	return strings.Join(words, " ")
}

func FromLowerCamelCase(value string) string {

	var result strings.Builder

	for i, r := range value {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('-')
			}

			result.WriteRune(unicode.ToLower(r))
			continue
		}

		result.WriteRune(r)
	}

	return result.String()
}
