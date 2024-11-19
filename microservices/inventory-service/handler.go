package handler

import (
	"context"
	pb "inventory-service/proto"
	"inventory-service/repository"
)

type InventoryServiceServer struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *InventoryServiceServer) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	availableQuantity, err := repository.GetAvailableStock(req.ProductId)
	if err != nil {
		return nil, err
	}

	inStock := availableQuantity > 0
	return &pb.CheckStockResponse{InStock: inStock, AvailableQuantity: availableQuantity}, nil
}

func (s *InventoryServiceServer) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	success, err := repository.DecreaseStock(req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}

	if success {
		return &pb.DecreaseStockResponse{Success: true, Message: "Stock decreased successfully"}, nil
	}
	return &pb.DecreaseStockResponse{Success: false, Message: "Insufficient stock"}, nil
}
