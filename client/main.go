package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kubernetes-playgrounds/kubernetes-service-discovery-grpc/IDL"
	"google.golang.org/grpc"
)

const (
	port = ":8082"
)

func main() {
	// Set up a connection to the server.
	var url = "localhost"
	conn, err := grpc.Dial(url+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())

		time.Sleep(5 * time.Second)
	}
}
