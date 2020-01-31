FROM golang:1.9

WORKDIR /go/src/github.com/guanw/kubernetes-service-discovery/server
COPY . .
RUN go get google.golang.org/grpc
RUN go get github.com/kubernetes-playgrounds/kubernetes-service-discovery-grpc/IDL
RUN make build-linux
EXPOSE 8082
CMD ["/go/src/github.com/guanw/kubernetes-service-discovery/server/server_binary_unix"]