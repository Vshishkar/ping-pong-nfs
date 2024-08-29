.ONESHELL:

build:
	cd ./src/server
	rm -f server
	go build -o ../../bin server

run: build
	cd ./bin
	./server
