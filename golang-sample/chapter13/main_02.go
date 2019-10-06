package main

import (
	"fmt"
	"os"
)

func main() {

	//os.newFile("adfadf", "text.txt")

	file, err := os.Open("text.txt")

	if err != nil {
		fmt.Println("오류를 처리")
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)
	fmt.Println(str)

}
