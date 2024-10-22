package problem2

func inversionNumberIterative(P []int) int {
	inversionNumber := 0
	for i := 0; i < len(P); i++ {
		for j := i + 1; j < len(P); j++ {
			if P[i] > P[j] {
				inversionNumber++
			}
		}
	}
	return inversionNumber
}

func countInversion(p []int, index int) int {
	if index == len(p)-1 {
		return 0
	}

	inversionNumber := 0
	for i := index + 1; i < len(p); i++ {
		if p[index] > p[i] {
			inversionNumber++
		}
	}

	return inversionNumber + countInversion(p, index+1)
}

func inversionNumberRecursive(p []int) int {
	return countInversion(p, 0)
}

func mergeAndCount(left, right []int) ([]int, int) {
	result := make([]int, 0, len(left)+len(right))
	inversionCount := 0
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			inversionCount += len(left) - i
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result, inversionCount
}

func countInversions(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}

	mid := len(arr) / 2
	leftHalf := make([]int, mid)
	rightHalf := make([]int, len(arr)-mid)

	copy(leftHalf, arr[:mid])
	copy(rightHalf, arr[mid:])

	leftHalf, leftInversions := countInversions(leftHalf)
	rightHalf, rightInversions := countInversions(rightHalf)
	mergedArray, splitInversions := mergeAndCount(leftHalf, rightHalf)

	totalInversions := leftInversions + rightInversions + splitInversions
	return mergedArray, totalInversions
}

func inversionNumberTwoRecurs(P []int) int {
	_, totalInversions := countInversions(P)
	return totalInversions
}
