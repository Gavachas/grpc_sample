package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sample_grpc "github.com/Gavachas/grpc_sample/grpc_s"
	"google.golang.org/grpc"
)

const defaultPort = "4041"

func main() {

	fmt.Println("Client")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:4041", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Maybe this should be in a separate function and the error handled?

	c := sample_grpc.NewItilServiceClient(cc)

	// read Region
	fmt.Println("Reading the region")
	readRegionReq := &sample_grpc.GetUserRequest{Id: 1}
	readRegionRes, readRegionErr := c.GetUserRegion(context.Background(), readRegionReq)
	if readRegionErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readRegionErr)
	}

	fmt.Printf("Region was read: %v \n", readRegionRes)

}
