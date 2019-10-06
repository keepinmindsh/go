package main

import "fmt"

func main() {
	xs := []float64{91, 22, 34, 23, 34}

	fmt.Println(average(xs))

	x, y := f()

	fmt.Println(x, y)

	result := add(10, 1, 2, 12, 3, 4)

	fmt.Println(result)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 10, 20
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
