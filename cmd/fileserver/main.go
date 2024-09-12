package main

import (
	"log"
	"nfs/internal/fileserver"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	PRC_PORT  = "4444"
	CONN_TYPE = "tcp"
)

func main() {
	port := CONN_PORT
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}

	rpcPort := PRC_PORT
	if len(os.Args) >= 3 {
		rpcPort = os.Args[2]
	}

	s := fileserver.MakeServer(fileserver.ServerConfig{
		ConnType: CONN_TYPE,
		ConnHost: CONN_HOST,
		ConnPort: port,
		RpcPort:  rpcPort,
	})

	log.Fatal(s.Start())
}
