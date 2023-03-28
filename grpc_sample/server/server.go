package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	sample_grpc "github.com/Gavachas/grpc_sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort = "4041"

type server struct {
	sample_grpc.ItilServiceServer
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Itil grpc Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	sample_grpc.RegisterItilServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
