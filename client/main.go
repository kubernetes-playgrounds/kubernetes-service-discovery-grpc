package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/kubernetes-playgrounds/kubernetes-service-discovery-grpc/IDL"
	"google.golang.org/grpc"
)

const (
	port = ":8082"
)

func main() {
	// Set up a connection to the server.
	var url = "0.0.0.0"
	conn, err := grpc.Dial(url+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		// Contact the server and print out its response.
		name := "Jude"
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Greeting: ", r.GetMessage())

		time.Sleep(5 * time.Second)
	}
}
