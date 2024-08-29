package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type Server struct {
	ln       net.Listener
	quitCh   chan struct{}
	bitrate  int
	connType string
	connHost string
	connPort string
	msgCh    chan []byte
}

type ServerConfig struct {
	BitRate  int
	ConnType string
	ConnHost string
	ConnPort string
}

func MakeServer(cfg ServerConfig) *Server {
	return &Server{
		quitCh:   make(chan struct{}),
		msgCh:    make(chan []byte),
		bitrate:  cfg.BitRate,
		connType: cfg.ConnType,
		connHost: cfg.ConnHost,
		connPort: cfg.ConnPort,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen(s.connType, s.connHost+":"+s.connPort)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	s.ln = ln
	defer s.ln.Close()

	go s.acceptLoop()
	go s.handleMsg()

	<-s.quitCh
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}

		fmt.Println("Accepted conn from ", conn.RemoteAddr())

		go s.handleRequest(conn)
	}
}

func (s *Server) handleMsg() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Println(string(msg))
		}
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, s.bitrate)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("failed:", err)
			return
		}

		msg := make([]byte, n)
		copy(msg, buf)
		s.msgCh <- msg
	}
}

func main() {
	s := MakeServer(ServerConfig{
		ConnType: CONN_TYPE,
		BitRate:  1024,
		ConnHost: CONN_HOST,
		ConnPort: CONN_PORT,
	})

	log.Fatal(s.Start())
}
