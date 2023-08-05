package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/widjaya-hernando/omdb-service/gen"
	"github.com/widjaya-hernando/omdb-service/models"
	"github.com/widjaya-hernando/omdb-service/omdb/repository"
)

func (s *Server) GetMovieByID(ctx context.Context, req *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	resp, err := repository.GetMovieByID(req.Id)
	if err != nil {
		log.Println("err-get-movie-by-id: ", err)
		return nil, err
	}

	response := *models.GetMovieByIDResponseToGrpcMessage(resp)

	return &response, nil
}
