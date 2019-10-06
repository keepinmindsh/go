package main

import "fmt"

func main() {

	type Circle struct {
		x float64
		y float64
		r float64
	}

	//c := new(Circle)

	c := Circle{x: 0, y: 100, r: 20}

	fmt.Println(c.x, c.y, c.r)

	fmt.Println(c)

	fmt.Println("vim-go")

	fmt.Println(circleArea(c))

	fmt.Println(c.area())
}

func (c *Circle) area() float64 {
	return c.r * c.r
}

func circleArea(c Circle) float64 {
	return c.r * c.r
}
