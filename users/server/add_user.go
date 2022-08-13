package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	var response pb.AddUserResponse

	if req == nil || req.User == nil {
		return nil, fmt.Errorf("user is not provided")
	}

	user := &models.User{
		UserName:  req.User.UserName,
		Phone:     req.User.Phone,
		Email:     req.User.Email,
		Address:   req.User.Address,
		CreatedAt: time.Now().Unix(),
	}

	err := s.repository.AddUser(nil, user)

	if err != nil {
		log.Println("err-add-user: ", err)
		return nil, err
	}

	response.User = models.UserToGrpcMessage(user)

	return &response, nil
}
