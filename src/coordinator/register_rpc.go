package main

import "fmt"

type RegisterFileServerArgs struct {
	Port string
	Name string
}

type RegisterFileServerReply struct {
	Message string
}

func (c *Coordinator) RegisterServer(args *RegisterFileServerArgs, reply *RegisterFileServerReply) error {
	fmt.Println("Got request to register server ", args.Port, args.Name)

	return nil
}
