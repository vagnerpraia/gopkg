package gplowercamelcase

import (
	"strings"
	"unicode"
)

func FromWords(value string) string {

	words := strings.FieldsFunc(value, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	var result strings.Builder

	for i, word := range words {
		word = strings.ToLower(word)

		if i > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			word = string(runes)
		}

		result.WriteString(word)
	}

	return result.String()
}

func ToWords(value string) string {

	var result strings.Builder

	for i, r := range value {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteByte(' ')
		}

		if i == 0 {
			result.WriteRune(unicode.ToUpper(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

func FromKebabCase(value string) string {

	var result strings.Builder

	upperNext := false

	for _, r := range value {
		if r == '-' {
			upperNext = true
			continue
		}

		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
			continue
		}

		result.WriteRune(r)
	}

	return result.String()
}
