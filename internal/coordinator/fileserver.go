package coordinator

import "net/rpc"

type FileServer struct {
	port    string
	rpcPort string
	host    string
	client  *rpc.Client
	id      int
}
