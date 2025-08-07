package main

import (
	"fmt"
	"net"
)

const message string = "Hello"

func main() {
	ln, err := net.Listen("tcp", "192.168.1.67:8080")
	log(err)
	fmt.Printf("Starting server on %s\n", ln.Addr())
	for {
		conn, connerr := ln.Accept()
		fmt.Printf("Connection from %s\n", conn.RemoteAddr())
	
		log(connerr)
		go func() {
		buffer := make([]byte, 1024)
		data, readErr := conn.Read(buffer)
		fmt.Printf("%s\n",string(buffer[:data]))
		log(readErr)
		defer conn.Close()
		conn.Write([]byte(message))
		}()
	}
	
		
}

func log(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}
