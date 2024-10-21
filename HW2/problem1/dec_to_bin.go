package problem1

import "strconv"

func decToBinIte(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		remainder := n % 2
		n = n / 2
		binary = strconv.Itoa(remainder) + binary
	}
	return binary
}

func decToBinRec(n int) string {
	if n == 0 {
		return "0"
	}

	if n == 1 {
		return "1"
	}

	return decToBinRec(n/2) + strconv.Itoa(n%2)
}
