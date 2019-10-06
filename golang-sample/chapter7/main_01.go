package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	// 클로저 테스트
	x := 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(factorial(3))
}

func factorial(x uint) uint {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}
