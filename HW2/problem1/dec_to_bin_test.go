package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecToBinIte(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{10, "1010"},
		{19, "10011"},
		{63, "111111"},
	}

	for _, tc := range testCases {
		binary := decToBinIte(tc.input)
		assert.Equal(t, tc.expected, binary, "For input %d", tc.input)
	}
}

func TestDecToBinRec(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{10, "1010"},
		{19, "10011"},
		{63, "111111"},
	}

	for _, tc := range testCases {
		binary := decToBinRec(tc.input)
		assert.Equal(t, tc.expected, binary, "For input %d", tc.input)
	}
}
