package gptemplate

import (
	"fmt"
	"testing"
)

func TestReplaceDelimited(t *testing.T) {

	tests := []struct {
		name         string
		input        string
		delimiter    Delimiter
		replacements map[string]string
		expected     string
	}{
		{
			name:      "replace multiple values",
			input:     "/dev/shm/{line}/{hypothesis}/simulator.yaml",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line":       "momentum",
				"hypothesis": "price-displacement",
			},
			expected: "/dev/shm/momentum/price-displacement/simulator.yaml",
		},
		{
			name:      "replace repeated value",
			input:     "{line}/{line}/{line}",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "momentum/momentum/momentum",
		},
		{
			name:      "keep unknown key",
			input:     "{line}/{unknown}",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "momentum/{unknown}",
		},
		{
			name:      "no placeholders",
			input:     "/dev/shm/simulator.yaml",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "/dev/shm/simulator.yaml",
		},
		{
			name:      "empty replacement",
			input:     "/dev/shm/{line}/simulator.yaml",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "",
			},
			expected: "/dev/shm//simulator.yaml",
		},
		{
			name:      "different delimiter",
			input:     "/dev/shm/[line]/simulator.yaml",
			delimiter: DelimiterSquareBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "/dev/shm/momentum/simulator.yaml",
		},
		{
			name:      "missing closing delimiter",
			input:     "/dev/shm/{line/simulator.yaml",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "/dev/shm/{line/simulator.yaml",
		},
		{
			name:      "empty placeholder",
			input:     "/dev/shm/{}/simulator.yaml",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"": "momentum",
			},
			expected: "/dev/shm/momentum/simulator.yaml",
		},
		{
			name:      "empty input",
			input:     "",
			delimiter: DelimiterCurlyBrackets,
			replacements: map[string]string{
				"line": "momentum",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Replace(test.input, test.delimiter, test.replacements)

			fmt.Println(result)
			fmt.Println(test.expected)
			fmt.Println("---")

			if result != test.expected {
				t.Errorf("ReplaceDelimited() = %q, expected %q", result, test.expected)
			}
		})
	}
}
