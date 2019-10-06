package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 5)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string, 10)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
