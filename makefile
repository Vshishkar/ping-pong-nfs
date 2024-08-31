.ONESHELL:

run-server1: build-server
	cd ./bin
	./fileserver 3333

run-server2: build-server
	cd ./bin
	./fileserver 3334

run-client: build-client
	cd ./bin
	./client

run-coordinator: build-coordinator
	cd ./bin
	./coordinator


build-server:
	go build -o ../../bin/ nfs/cmd/fileserver

build-coordinator:
	go build -o ../../bin/ nfs/cmd/coordinator


