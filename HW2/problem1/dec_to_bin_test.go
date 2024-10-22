package problem1

import (
	"strconv"
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

func BenchmarkDecToBinIte(b *testing.B) {
	testCases := []int{0, 1, 10, 100, 1000, 10000}

	for _, tc := range testCases {
		b.Run(strconv.Itoa(tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				decToBinIte(tc)
			}
		})
	}
}

func BenchmarkDecToBinRec(b *testing.B) {
	testCases := []int{0, 1, 10, 100, 1000, 10000}

	for _, tc := range testCases {
		b.Run(strconv.Itoa(tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				decToBinRec(tc)
			}
		})
	}
}

// strconv 內建函數
func decToBinBuiltin(n int) string {
	if n == 0 {
		return "0"
	}
	return strconv.FormatInt(int64(n), 2)
}

func BenchmarkDecToBinBuiltin(b *testing.B) {
	testCases := []int{0, 1, 10, 100, 1000, 10000}

	for _, tc := range testCases {
		b.Run(strconv.Itoa(tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				decToBinBuiltin(tc)
			}
		})
	}
}
