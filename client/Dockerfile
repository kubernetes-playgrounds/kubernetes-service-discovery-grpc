FROM golang:1.9

WORKDIR /go/src/github.com/guanw/kubernetes-service-discovery/client
COPY . .
RUN go get google.golang.org/grpc
RUN go get github.com/kubernetes-playgrounds/kubernetes-service-discovery-grpc/IDL
RUN make build-linux
CMD ["/go/src/github.com/guanw/kubernetes-service-discovery/client/client_binary_unix"]