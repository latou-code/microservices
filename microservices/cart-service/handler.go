package handler

import (
	pb "cart-service/proto"
	"cart-service/repository"
	"context"
)

type CartServiceServer struct {
	pb.UnimplementedCartServiceServer
}

func (s *CartServiceServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	err := repository.AddProductToCart(req.UserId, req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &pb.AddToCartResponse{Message: "Product added to cart"}, nil
}

func (s *CartServiceServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	items, err := repository.GetCartItems(req.UserId)
	if err != nil {
		return nil, err
	}

	var cartItems []*pb.CartItem
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &pb.GetCartResponse{Items: cartItems}, nil
}
