package main

import (
	"fmt"
	"net"
	"strings"
)



var routes = map[string]string{
	"/index": "hello",
}

type HTMLREQ struct {
	requestType string
	requested string
	route string
}

type clientConn struct {
	net.Conn
}

const message string = "Hello"

func main() {
	ln, err := net.Listen("tcp", "192.168.1.67:8080")
	log(err)
	fmt.Printf("Starting server on %s\n", ln.Addr())
	for {
		conn, connerr := ln.Accept()
		fmt.Printf("Connection from %s\n", conn.RemoteAddr())
	
		log(connerr)
		clientConn := &clientConn{conn}
		go func() {
			clientConn.handleClient()
			
		}()
	}
	
		
}

func (conn *clientConn) handleClient(){
	var(
		method string
		uri string
	)
	buffer := make([]byte, 1024)
	data, readErr := conn.Read(buffer)
	stringData := string(buffer[:data])
	fmt.Printf("%s\n",stringData)
	log(readErr)
	
	//quoted := strconv.Quote(stringData)
	//var slash string = `/`
	request := strings.Split(stringData, "\r")[0]
	method = strings.Split(request, "/")[0]
	uri = strings.Split(stringData, "\r")[1]
	fmt.Printf("Request: %v\n", request)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Uri: %s\n", uri)
	if (method == "GET"){
		fmt.Printf("Get request logged: %s\n", uri)
	}
	//fmt.Printf("[+] Raw string: %s\n", quoted)
	defer conn.Close()
	//conn.Write([]byte(message))
}



func log(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}
