package problem3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotatedIncreasing(t *testing.T) {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 0, 4}, false},
		{[]int{8, 9, 11, 12, 5}, true},
		{[]int{7, 9, 11, 12, 14}, true},
		{[]int{15, 9, 11, 12, 13}, true},
		{[]int{4, 5, 6, 1, 2, 3}, true},
		{[]int{4, 6, 1, 5, 3}, false},
	}

	for _, tc := range testCases {
		result := RotatedIncreasing(tc.input)
		assert.Equal(t, tc.expected, result, "For input %v", tc.input)
	}

}
