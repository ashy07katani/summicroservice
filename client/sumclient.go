package main

import (
	"context"
	pb "summicroservice/proto"
)

func doSum(c pb.SumServiceClient, req *pb.SumRequest) (*pb.SumResponse, error) {
	return c.DoSum(context.Background(), req)
}
