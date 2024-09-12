package rpcmodels

type RegisterFileServerArgs struct {
	Port    string
	Host    string
	RpcPort string
}

type RegisterFileServerReply struct {
	Id int
}
