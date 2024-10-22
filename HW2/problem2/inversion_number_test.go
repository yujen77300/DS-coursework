package problem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInversionNumberIterative(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 6, 2, 5, 1, 3}, 10},
		{[]int{12, 11, 13, 5, 6, 7}, 10},
		{[]int{35, 11, 26, 13, 64, 21, 44, 6, 100, 57, 77, 82}, 19},
	}

	for _, tc := range testCases {
		totalInversion := inversionNumberIterative(tc.input)
		assert.Equal(t, tc.expected, totalInversion, "For input %v", tc.input)
	}
}

func TestInversionNumberRecursive(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 6, 2, 5, 1, 3}, 10},
		{[]int{12, 11, 13, 5, 6, 7}, 10},
		{[]int{35, 11, 26, 13, 64, 21, 44, 6, 100, 57, 77, 82}, 19},
	}

	for _, tc := range testCases {
		totalInversion := inversionNumberRecursive(tc.input)
		assert.Equal(t, tc.expected, totalInversion, "For input %v", tc.input)
	}
}

func TestInversionNumberTwoRecurs(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 6, 2, 5, 1, 3}, 10},
		{[]int{12, 11, 13, 5, 6, 7}, 10},
		{[]int{35, 11, 26, 13, 64, 21, 44, 6, 100, 57, 77, 82}, 19},
	}

	for _, tc := range testCases {
		totalInversion := inversionNumberTwoRecurs(tc.input)
		assert.Equal(t, tc.expected, totalInversion, "For input %v", tc.input)
	}
}
