package fileserver

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"log/slog"
	"net"
	"net/rpc"
	"nfs/internal/rpcmodels"
	"os"
)

type Server struct {
	ln       net.Listener
	quitCh   chan struct{}
	connType string
	connHost string
	connPort string
	rpcPort  string
	msgCh    chan []byte
	id       int
}

type ServerConfig struct {
	ConnType string
	ConnHost string
	ConnPort string
	RpcPort  string
}

func MakeServer(cfg ServerConfig) *Server {
	return &Server{
		quitCh:   make(chan struct{}),
		msgCh:    make(chan []byte),
		connType: cfg.ConnType,
		connHost: cfg.ConnHost,
		connPort: cfg.ConnPort,
		rpcPort:  cfg.RpcPort,
	}
}

func (s *Server) Start() error {
	s.callRegisterToCoordinator()

	ln, err := net.Listen(s.connType, s.connHost+":"+s.connPort)
	if err != nil {
		slog.Info("Error listening:", "err", err)
		os.Exit(1)
	}

	s.ln = ln
	defer s.ln.Close()

	go s.acceptLoop()
	go s.acceptRPCs()
	go s.handleMsg()

	<-s.quitCh
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			slog.Info("Error accepting", "err", err)
		}

		slog.Info("Accepted conn from", "addr", conn.RemoteAddr())

		go s.handleRequest(conn)
	}
}

func (s *Server) acceptRPCs() {
	slog.Info("making server accept rpcs", "id", s.id)
	ln, err := net.Listen(s.connType, s.connHost+":"+s.rpcPort)
	if err != nil {
		slog.Error("failed to connect to", "port", s.rpcPort, "err", err)
		return
	}

	err = rpc.Register(s)
	if err != nil {
		slog.Error("Error registering server", "err", err)
	}

	slog.Info("registered server", "id", s.id)
	rpc.Accept(ln)
}

func (s *Server) handleMsg() {
	for {
		select {
		case msg := <-s.msgCh:
			slog.Info(string(msg))
		}
	}
}

func (s *Server) callRegisterToCoordinator() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error ", err)
	}

	client := rpc.NewClient(conn)
	args := &rpcmodels.RegisterFileServerArgs{
		Port:    s.connPort,
		Host:    s.connHost,
		RpcPort: s.rpcPort,
	}
	reply := &rpcmodels.RegisterFileServerReply{}

	err = client.Call("Coordinator.RegisterServer", args, reply)
	if err != nil {
		slog.Info("error after register server", "err", err)
		return
	}

	slog.Info("reply from coordinator", "reply", reply.Id)
	s.id = reply.Id
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := new(bytes.Buffer)
	var size int64
	binary.Read(conn, binary.LittleEndian, &size)

	n, err := io.CopyN(buf, conn, size)
	if err != nil {
		slog.Error("failed:", "err", err)
		return
	}
	slog.Info("read %d bytes", "n", n)
	s.msgCh <- buf.Bytes()
}
