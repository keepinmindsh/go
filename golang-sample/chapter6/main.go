package main

import "fmt"

func main(){
	var x [5]float64
	x[4] = 100
	fmt.Println(x)

	var total float64 = 0

	for i := 0; i < len(x) ; i ++ {
		total += x[i]
	}

        fmt.Println(total/float64(len(x)))

	
	x = [5]float64{ 91,12,12,31}

	fmt.Println(x)

        var total1 float64 = 0
        for _, value := range x {
		total1 += value

	}

 	fmt.Println(total1/float64(len(x)))

}
