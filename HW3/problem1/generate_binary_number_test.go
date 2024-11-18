package problem1

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBinaryNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "when n is 10",
			input:    10,
			expected: []string{"1", "10", "11", "100", "101", "110", "111", "1000", "1001", "1010"},
		},
		{
			name:     "when n is 13",
			input:    13,
			expected: []string{"1", "10", "11", "100", "101", "110", "111", "1000", "1001", "1010", "1011", "1100", "1101"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateBinaryNumbers(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}