package main

import "fmt"

func main() {
	//fmt.Println("vim-go")
	go f(0)
	go f(10)
	var input string
	fmt.Scanf("%d", &input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
