package problem1

import (
	"errors"
	"strconv"
)

type Queue struct {
	Elements []string
}

// Enqueue - add an element of type int to the end of our queue
func (q *Queue) Enqueue(item string) {
	q.Elements = append(q.Elements, item)
}

// Dequeue - return the first element from our queue
func (q *Queue) Dequeue() (string, error) {
	if len(q.Elements) == 0 {
		return "", errors.New("empty Queue")
	}
	var firstElement string
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

// Size - return the size of the queue
func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func generateBinaryNumbers(n int) []string {
	queue := &Queue{}
	result := []string{}

	queue.Enqueue("1")

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		// Convert binary string to integer
		num, _ := strconv.ParseInt(current, 2, 64)

		if int(num) <= n {
			result = append(result, current)

			append0 := current + "0"
			append1 := current + "1"

			num0, _ := strconv.ParseInt(append0, 2, 64)
			num1, _ := strconv.ParseInt(append1, 2, 64)

			if int(num0) <= n {
				queue.Enqueue(append0)
			}
			if int(num1) <= n {
				queue.Enqueue(append1)
			}
		}
	}
	return result
}
