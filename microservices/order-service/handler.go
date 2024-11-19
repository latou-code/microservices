package handler

import (
	"context"
	pb "order-service/proto"
	"order-service/repository"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderID, err := repository.CreateOrder(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: orderID, Message: "Order created successfully"}, nil
}
