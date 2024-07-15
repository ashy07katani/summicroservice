package main

import (
	"context"
	pb "summicroservice/proto"
)

type serveSumImpl struct {
	pb.UnimplementedSumServiceServer
}

func (s *serveSumImpl) DoSum(ctx context.Context, r *pb.SumRequest) (res *pb.SumResponse, err error) {
	res = &pb.SumResponse{
		Result: r.Num1 + r.Num2,
	}

	return
}
