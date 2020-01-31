package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/kubernetes-playgrounds/kubernetes-service-discovery-grpc/IDL"
	"google.golang.org/grpc"
)

const (
	port = ":8082"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err.Error())
	}
}
