package handler

import (
	"context"
	"fmt"
	"log"

	pb "inventory-service/proto"
	"inventory-service/repository"
)

// InventoryRepository - интерфейс для репозитория.
type InventoryRepository interface {
	GetAvailableStock(productID string) (int64, error)
	DecreaseStock(productID string, quantity int64) (bool, error)
}

// InventoryServiceServer - сервер для управления запасами.
type InventoryServiceServer struct {
	pb.UnimplementedInventoryServiceServer
	Repo InventoryRepository
}

// NewInventoryServiceServer - конструктор для InventoryServiceServer.
func NewInventoryServiceServer(repo InventoryRepository) *InventoryServiceServer {
	return &InventoryServiceServer{Repo: repo}
}

// CheckStock - проверяет наличие товара на складе.
func (s *InventoryServiceServer) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	// Валидация входных данных
	if req.ProductId == "" {
		log.Println("CheckStock: ProductId is empty")
		return nil, fmt.Errorf("product ID cannot be empty")
	}

	// Получение доступного количества
	availableQuantity, err := s.Repo.GetAvailableStock(req.ProductId)
	if err != nil {
		log.Printf("CheckStock: error getting stock for ProductId=%s: %v\n", req.ProductId, err)
		return nil, fmt.Errorf("failed to check stock: %w", err)
	}

	// Формирование ответа
	inStock := availableQuantity > 0
	log.Printf("CheckStock: ProductId=%s, AvailableQuantity=%d, InStock=%v\n", req.ProductId, availableQuantity, inStock)
	return &pb.CheckStockResponse{
		InStock:          inStock,
		AvailableQuantity: availableQuantity,
	}, nil
}

// DecreaseStock - уменьшает количество товара на складе.
func (s *InventoryServiceServer) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	// Валидация входных данных
	if req.ProductId == "" {
		log.Println("DecreaseStock: ProductId is empty")
		return nil, fmt.Errorf("product ID cannot be empty")
	}
	if req.Quantity <= 0 {
		log.Printf("DecreaseStock: invalid quantity %d for ProductId=%s\n", req.Quantity, req.ProductId)
		return nil, fmt.Errorf("quantity must be greater than zero")
	}

	// Уменьшение количества на складе
	success, err := s.Repo.DecreaseStock(req.ProductId, req.Quantity)
	if err != nil {
		log.Printf("DecreaseStock: error decreasing stock for ProductId=%s: %v\n", req.ProductId, err)
		return nil, fmt.Errorf("failed to decrease stock: %w", err)
	}

	// Формирование ответа
	if success {
		log.Printf("DecreaseStock: stock decreased successfully for ProductId=%s, Quantity=%d\n", req.ProductId, req.Quantity)
		return &pb.DecreaseStockResponse{
			Success: true,
			Message: "Stock decreased successfully",
		}, nil
	}

	log.Printf("DecreaseStock: insufficient stock for ProductId=%s, Quantity=%d\n", req.ProductId, req.Quantity)
	return &pb.DecreaseStockResponse{
		Success: false,
		Message: "Insufficient stock",
	}, nil
}
