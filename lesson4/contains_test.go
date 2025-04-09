package lesson4

import (
	"strings"
	"testing"
)

func Test_Contains(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		substr   string
		expected bool
	}{
		{
			name:     "single",
			str:      "it's a test",
			substr:   "test",
			expected: true,
		},
		{
			name:     "multiple",
			str:      "it's a test",
			substr:   "best",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := strings.Contains(tc.str, tc.substr)
			if tc.expected != result {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
