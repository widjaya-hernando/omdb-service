package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (s *Server) GetProductByTransactionId(ctx context.Context, req *pb.GetProductByTransactionIdRequest) (*pb.GetProductByTransactionIdResponse, error) {
	var response pb.GetProductByTransactionIdResponse

	product, err := s.repository.GetProductByTransactionId(nil, req.GetId())
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-product: ", err)
		return nil, err
	}
	if product != nil {
		response.Product = models.ProductToGrpcMessage(product)
	}

	return &response, nil
}
