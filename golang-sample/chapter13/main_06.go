package main

import  ( "errors" ; "fmt" )

func main(){
	// https://www.bugsnag.com/blog/go-errors
	// 다음시간에 해보기
	err := errors.New("Err Messages")

	fmt.Println(err)

}


