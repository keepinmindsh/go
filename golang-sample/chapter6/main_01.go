package main

import "fmt"

func main() {
	
	arr := make([]float64, 5 )
	
	fmt.Println(arr)

	arrList1 := []int{ 12,123,123}

        arrList2 := make([]int, 1, 10 )

	arrList4 := arrList1[1:2]

	arrList3 := make([]int, 4 )

	arrList5 := append(arrList1, 23,234,234)

	fmt.Println(arr)
 	fmt.Println(arrList1)
	fmt.Println(arrList2)
	fmt.Println(arrList3)
	fmt.Println(arrList4)
	fmt.Println(arrList5)
	fmt.Println(arrList1)


}
