# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=server_binary
BINARY_UNIX=$(BINARY_NAME)_unix

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t guanw0826/kubernetes-service-discovery-grpc-server .

docker-push:
	docker push guanw0826/kubernetes-service-discovery-grpc-server

docker-run:
	docker run -d -p 8082:8082 guanw0826/kubernetes-service-discovery-grpc-server