package problem2

import (
	"math/rand"
	"strconv"
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

func generateSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(n)
	}
	return slice
}

func BenchmarkInversionNumber(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		testData := generateSlice(size)

		b.Run("Iterative/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inversionNumberIterative(testData)
			}
		})

		b.Run("Recursive/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inversionNumberRecursive(testData)
			}
		})
		b.Run("MergeSort/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inversionNumberTwoRecurs(testData)
			}
		})
	}
}
