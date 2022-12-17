package grpc_transport

import (
	"context"
	pb "github.com/Jungle20m/electricity/internal/module/grab/grpc-transport/proto"
)

func (t *Transportation) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return &pb.GetOrderResponse{Id: int32(200)}, nil
}

func (t *Transportation) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.GetServiceResponse, error) {
	return &pb.GetServiceResponse{Id: int32(200)}, nil
}
