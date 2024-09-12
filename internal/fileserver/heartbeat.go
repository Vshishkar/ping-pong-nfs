package fileserver

import (
	"fmt"
	"log/slog"
	"nfs/internal/rpcmodels"
)

func (s *Server) Heartbeat(args *rpcmodels.HeartbeatArgs, reply *rpcmodels.HeartbeatReply) error {
	slog.Info("got heartbeat from coordinator")

	reply.Message = fmt.Sprintf("server %v is healthy", s.id)
	return nil
}
