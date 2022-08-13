package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (s *Server) GetUserByTransactionId(ctx context.Context, req *pb.GetUserByTransactionIdRequest) (*pb.GetUserByTransactionIdResponse, error) {
	var response pb.GetUserByTransactionIdResponse

	ser, err := s.repository.GetUserByTransactionId(nil, req.GetId())
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-ser: ", err)
		return nil, err
	}
	if ser != nil {
		response.User = models.UserToGrpcMessage(ser)
	}

	return &response, nil
}
