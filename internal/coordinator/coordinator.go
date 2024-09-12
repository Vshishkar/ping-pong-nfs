package coordinator

import (
	"log/slog"
	"net"
	"net/rpc"
	"sync"
)

type Coordinator struct {
	mu              sync.RWMutex
	serverIdCounter int
	fileServers     map[int]FileServer
	quitCh          chan struct{}
}

func MakeCoordinator() *Coordinator {
	return &Coordinator{
		fileServers: make(map[int]FileServer),
		quitCh:      make(chan struct{}),
	}
}

func (c *Coordinator) Start() error {
	go c.acceptRPCs()
	go c.startHeartbeats()
	<-c.quitCh
	return nil
}

func (c *Coordinator) acceptRPCs() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		slog.Error("failed to connect to 8080", "err", err)
		return
	}

	rpc.Register(c)
	rpc.Accept(ln)
}
