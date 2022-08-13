package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/models"
)

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var response pb.UpdateUserResponse

	if req == nil || req.User == nil {
		return nil, fmt.Errorf("user is not provided")
	}

	user := &models.User{
		ID:        req.User.Id,
		UserName:  req.User.UserName,
		Phone:     req.User.Phone,
		Email:     req.User.Email,
		Address:   req.User.Address,
		UpdatedAt: time.Now().Unix(),
	}

	err := s.repository.UpdateUser(nil, user)

	if err != nil {
		log.Println("err-update-user: ", err)
		return nil, err
	}

	response.User = models.UserToGrpcMessage(user)

	return &response, nil
}
