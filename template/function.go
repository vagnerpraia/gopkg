package gptemplate

import (
	"strings"
)

func Replace(input string, delimiter Delimiter, replacements map[string]string) string {

	var result strings.Builder

	for i := 0; i < len(input); {
		if input[i] != delimiter.Open() {
			result.WriteByte(input[i])
			i++

			continue
		}

		end := strings.IndexByte(input[i+1:], delimiter.Close())
		if end == -1 {
			result.WriteByte(input[i])
			i++

			continue
		}

		end += i + 1

		key := input[i+1 : end]

		if value, ok := replacements[key]; ok {
			result.WriteString(value)
		} else {
			result.WriteString(input[i : end+1])
		}

		i = end + 1
	}

	return result.String()
}
