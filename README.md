# ping-pong-nfs

This repo contains ping pong network file system. 

NFS looks like a ring of servers where every adjacent server have a connection and data is streamed in clockwise direction. 

Client connects to a random server and tries to embed a new file to a stream. 

When client wants to receive a file, they connect to a server and wait until the stream arrived on a server. 

Information about cluster membership and bytes send will be stored in zookeeper 


