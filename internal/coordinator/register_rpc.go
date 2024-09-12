package coordinator

import (
	"log/slog"
	"net"
	"net/rpc"
	"nfs/internal/rpcmodels"
)

func (c *Coordinator) RegisterServer(args *rpcmodels.RegisterFileServerArgs, reply *rpcmodels.RegisterFileServerReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	slog.Info("Got request to register server ", "port", args.Port, "host", args.Host)
	c.serverIdCounter++
	reply.Id = c.serverIdCounter

	c.fileServers[c.serverIdCounter] = FileServer{
		port:    args.Port,
		host:    args.Host,
		id:      c.serverIdCounter,
		rpcPort: args.RpcPort,
		client:  nil,
	}

	go c.connectToFS(c.serverIdCounter)

	return nil
}

func (c *Coordinator) connectToFS(serverId int) {
	c.mu.Lock()
	fsConfig := c.fileServers[serverId]
	c.mu.Unlock()

	for {
		conn, err := net.Dial("tcp", fsConfig.host+":"+fsConfig.rpcPort)
		if err != nil {
			slog.Info("error ", "err", err)
			continue
		}

		client := rpc.NewClient(conn)
		fsConfig.client = client

		c.mu.Lock()
		c.fileServers[serverId] = fsConfig
		c.mu.Unlock()
		return
	}
}
