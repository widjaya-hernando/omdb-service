package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var response pb.GetProductResponse

	product, err := s.repository.GetProduct(nil, req.GetId())
	if err != nil {
		log.Println("err-get-product: ", err)
		return nil, err
	}

	response.Product = models.ProductToGrpcMessage(product)

	return &response, nil
}
