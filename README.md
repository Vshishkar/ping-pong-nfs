# ping-pong-nfs

This repo contains ping pong network file system.

NFS looks like a ring of servers where every adjacent server have a connection and data is streamed in clockwise direction.

Client connects to a random server and tries to embed a new file to a stream.

When client wants to receive a file, they connect to a server and wait until the stream arrived on a server.

Information about cluster membership and bytes send will be stored in zookeeper.

TODO:

1. Create TCP File Server
    - Create a server which listens TCP connection
    - Add configuration using cla
    - Use stream buffer and create simple encoding format

2. Create a simple TCP File Server client to test File server

    - Create client and send random file over network to file server

3. Add logging

    - Add monitoring and logging (need to research opensource solutions)

4. Create Coordinator service

    - Add coordinator service
    - Make file servers register themself with coordinator (Add RPC method on Coordinator)
    - Add Heartbeat PRC on coordinator service to call file servers (Add PRC method on FileServer)
    - Store coordinator port at root in a txt file

5. Add Chain replication structure.
    - Add config for chain length
    - Add RPC at coordinator service to allow File servers to add themself on a chain

TODO: revisit this list


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