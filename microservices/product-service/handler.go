package handler

import (
	"context"
	pb "product-service/proto"
	"product-service/repository"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	id, err := repository.CreateProduct(req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{Id: id}, nil
}

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := repository.GetProductByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductServiceServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := repository.ListProducts()
	if err != nil {
		return nil, err
	}

	var productList []*pb.Product
	for _, p := range products {
		productList = append(productList, &pb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return &pb.ListProductsResponse{Products: productList}, nil
}
