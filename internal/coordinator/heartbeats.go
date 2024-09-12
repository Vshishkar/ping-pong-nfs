package coordinator

import (
	"log/slog"
	"nfs/internal/rpcmodels"
	"time"
)

func (c *Coordinator) startHeartbeats() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			{
				c.sendHeartbeats()
			}
		}
	}
}

func (c *Coordinator) sendHeartbeats() {
	c.mu.Lock()
	defer c.mu.Unlock()

	slog.Info("sending heartbeats")

	for id, cfg := range c.fileServers {
		slog.Info("calling server", "id", id)
		client := cfg.client

		args := &rpcmodels.HeartbeatArgs{
			Message: "Hi!",
		}

		reply := &rpcmodels.HeartbeatReply{}

		err := client.Call("Server.Heartbeat", args, reply)
		if err != nil {
			slog.Error("error", "err", err)
		}
	}
}
