package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	fmt.Println("Hello from coordinator!")

	c := &Coordinator{}

	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to connect to 8080", err)
	}

	rpc.Register(c)
	rpc.Accept(ln)

	// not sure this is needed
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("got error while accepting connection", err)
		}

		fmt.Println("accepted connection from", conn.RemoteAddr())
	}
}
