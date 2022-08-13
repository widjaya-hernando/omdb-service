package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	var response pb.UpdateProductResponse

	if req == nil || req.Product == nil {
		return nil, fmt.Errorf("product is not provided")
	}

	product := &models.Product{
		ID:          req.Product.Id,
		ProductName: req.Product.ProductName,
		Quantity:    req.Product.Quantity,
		Price:       req.Product.Price,
		UpdatedAt:   time.Now().Unix(),
	}

	err := s.repository.UpdateProduct(nil, product)

	if err != nil {
		log.Println("err-update-product: ", err)
		return nil, err
	}

	response.Product = models.ProductToGrpcMessage(product)

	return &response, nil
}
