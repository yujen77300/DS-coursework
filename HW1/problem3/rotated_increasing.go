package problem3

func RotatedIncreasing(sequence []int) bool {
	length := len(sequence)
	if length <= 2 {
		return true
	}

	directionChanges := 0
	for i := 0; i < length; i++ {
		if sequence[i] > sequence[(i+1)%length] {
			directionChanges += 1
		}

		if directionChanges > 1 {
			return false
		}
	}

	return true

}
