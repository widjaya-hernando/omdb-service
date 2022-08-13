package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	var response pb.AddProductResponse

	if req == nil || req.Product == nil {
		return nil, fmt.Errorf("product is not provided")
	}

	product := &models.Product{
		ProductName: req.Product.ProductName,
		Quantity:    req.Product.Quantity,
		Price:       req.Product.Price,
		CreatedAt:   time.Now().Unix(),
	}

	err := s.repository.AddProduct(nil, product)

	if err != nil {
		log.Println("err-add-product: ", err)
		return nil, err
	}

	response.Product = models.ProductToGrpcMessage(product)

	return &response, nil
}
