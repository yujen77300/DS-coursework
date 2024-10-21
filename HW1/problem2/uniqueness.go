package problem2

func uniqueness(arr []int) bool {
	for i, v := range arr {
		for i = i + 1; i < len(arr); i++ {
			if v == arr[i] {
				return false
			}
		}
	}
	return true
}
