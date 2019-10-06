package main

import (
	"crypto/sha256"
	"fmt"

)

func main(){
	h := sha256.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})

	s := string(bs)

	fmt.Println(s)

	fmt.Println(bs)
}