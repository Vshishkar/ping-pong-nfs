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
	cd ./src/fileserver
	go build -o ../../bin/ fileserver

build-client:
	cd ./src/client
	go build -o ../../bin/ client


build-coordinator:
	cd ./src/coordinator
	go build -o ../../bin/ coordinator


