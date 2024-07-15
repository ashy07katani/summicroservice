package main

import (
	"log"
	"net"
	pb "summicroservice/proto"

	"google.golang.org/grpc"
)

func main() {
	const address = "localhost:50051"
	//create a listener
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error in establishing the connection: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &serveSumImpl{})
	log.Printf("Server started at port %v", address)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
