package main

import "fmt"

func main() {
	//fmt.Println("vim-go")
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(xPtr *int) {
	*xPtr = 0
}
