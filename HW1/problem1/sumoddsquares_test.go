package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOddSquares(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"n = 2", 2, 1},
		{"n = 10", 10, 165},
		{"n = 0", 0, 0},
		{"n = 4", 4, 10},
		{"n = -5", -5, 0},
		{"n = 1", 1, 0},
		{"n = 3", 3, 1},
		{"n = 5", 5, 10},
		{"n = 7", 7, 35},
		{"n = 9", 9, 84},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sumoddsquares(tc.input)
			assert.Equal(t, tc.expected, result, "For input %d", tc.input)
		})
	}
}