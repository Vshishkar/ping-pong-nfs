package server

import (
	"bytes"
	"coordinator"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"os"
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
			fmt.Println(msg)
		}
	}
}

func (s *Server) callRegisterToCoordinator() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error ", err)
	}

	client := rpc.NewClient(conn)

	args := &coordinator.RegisterFileServerArgs{}

}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := new(bytes.Buffer)
	var size int64
	binary.Read(conn, binary.LittleEndian, &size)

	n, err := io.CopyN(buf, conn, size)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Printf("read %d bytes", n)
	s.msgCh <- buf.Bytes()
}
