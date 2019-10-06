package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(){
	// 포트 대기
	ln, err := net.Listen("tcp",":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 연결 수락
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client(){
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "HelloWorld"

	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main()  {
	go server()
	//go client()

	var input string
	fmt.Scanln(&input)
}


