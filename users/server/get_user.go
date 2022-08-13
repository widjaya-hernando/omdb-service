package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse

	store, err := s.repository.GetUser(nil, req.GetId())
	if err != nil {
		log.Println("err-get-store: ", err)
		return nil, err
	}

	response.User = models.UserToGrpcMessage(store)

	return &response, nil
}
