package server

import (
	"context"
	"log"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var response pb.ListUsersResponse

	users, err := s.repository.ListUsers(nil)
	if err != nil {
		log.Println("err-list-users: ", err)
		return nil, err
	}

	for _, user := range users {
		response.Users = append(response.Users, models.UserToGrpcMessage(&user))
	}

	return &response, nil
}
