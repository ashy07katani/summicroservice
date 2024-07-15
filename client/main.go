package main

import (
	"log"
	pb "summicroservice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const address = "localhost:50051"
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	client := pb.NewSumServiceClient(conn)
	sumRequest := &pb.SumRequest{
		Num1: 10,
		Num2: 21,
	}
	response, err := doSum(client, sumRequest)
	if err != nil {
		log.Fatalf("Error in calculating the sum of %v and %v: %v", sumRequest.Num1, sumRequest.Num2, err)
	}
	log.Printf("The sum of %v and %v is : %v", sumRequest.Num1, sumRequest.Num2, response.Result)
}
