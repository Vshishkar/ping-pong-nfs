package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"nfs/internal/coordinator"
)

func main() {
	fmt.Println("Hello from coordinator!")

	c := &coordinator.Coordinator{}

	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to connect to 8080", err)
	}

	rpc.Register(c)
	rpc.Accept(ln)
}
