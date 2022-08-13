package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var response pb.ListProductsResponse

	products, err := s.repository.ListProducts(nil)
	if err != nil {
		log.Println("err-list-products: ", err)
		return nil, err
	}

	for _, product := range products {
		response.Products = append(response.Products, models.ProductToGrpcMessage(&product))
	}

	return &response, nil
}
