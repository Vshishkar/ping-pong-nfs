package main

import (
	"fileserver/server"
	"log"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	port := CONN_PORT
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}

	s := server.MakeServer(server.ServerConfig{
		ConnType: CONN_TYPE,
		BitRate:  1024,
		ConnHost: CONN_HOST,
		ConnPort: port,
	})

	log.Fatal(s.Start())
}
