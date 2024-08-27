# ping-pong-nfs

This repo contains ping pong network file system.

NFS looks like a ring of servers where every adjacent server have a connection and data is streamed in clockwise direction.

Client connects to a random server and tries to embed a new file to a stream.

When client wants to receive a file, they connect to a server and wait until the stream arrived on a server.

Information about cluster membership and bytes send will be stored in zookeeper.

TODO:

1. Create a Golang server with 2 APIs:

    - POST: sendFile(path, byte[])
    - GET: getFile(path): byte[]

    store file in memory. create a simple dictionary.

2. Add zookeeper for coordination between servers

    - Run zookeeper instance and make server register itself with zookeeper
    - Run multiple servers

3. Connect servers to each other

    - Find a way to organize a ring structure in zookeeper
    - Create tmp heartbeat endpoint for servers to talk to each other
    - Make servers send heartbeats to each other

4. Add http streaming between servers

    - add constant http streaming between servers
    - start moving files through the stream