tidy:
	go mod tidy

deps:
	go mod vendor

build: build-server build-node-image
	go build -o bin/main main.go

build-node-image:
	docker build -t kuber-barn-node .

build-server:
	go build -o bin/server api/http/server.go