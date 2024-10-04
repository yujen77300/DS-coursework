package problem1

func sumoddsquares(n int) int {
	total := 0
	for i := 1; i < n; i = i + 2 {
		total += i * i
	}
	return total
}
