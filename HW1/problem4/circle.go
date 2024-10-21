package main

import "fmt"

type Circle struct {
	radius float64
}

func NewCircle(r float64) *Circle {
	return &Circle{radius: r}
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle with radius %.2f", c.radius)
}

func main() {
	var r float64
	fmt.Print("Please give the radius of the circle: ")
	_, err := fmt.Scanf("%f", &r)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	circle := NewCircle(r)

	fmt.Println(circle)
	fmt.Printf("The area and perimeter of the input circle are %.2f and %.2f, respectively.\n",
		circle.Area(), circle.Perimeter())

}
