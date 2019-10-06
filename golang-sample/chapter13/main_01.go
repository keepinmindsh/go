package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
	)

	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
	str2 := []byte{'t', 'e', 's', 't'}
	fmt.Println(str2)
}
