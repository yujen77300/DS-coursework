package problem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueness(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"All unique", []int{1, 2, 3, 4, 5, 6}, true},
		{"Contains duplicate", []int{1, 2, 3, 3, 5, 6}, false},
		{"Empty slice", []int{}, true},
		{"Single element", []int{1}, true},
		{"Two same elements", []int{1, 1}, false},
		{"Longer unique sequence", []int{1, 2, 3, 4, 5, 6, 7}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uniqueness(tc.input)
			assert.Equal(t, tc.expected, result, "For input %v", tc.input)
		})
	}
}
