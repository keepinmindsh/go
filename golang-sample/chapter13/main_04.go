package main

import (
	"os"
)

func main() {
	file, err := os.Create("text.txt")

	if err != nil {
		return
	}

	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString("sdfasdfasdfasdfasdfasdf")
	}

	//file.WriteString("testststatsfdsdf")
}
