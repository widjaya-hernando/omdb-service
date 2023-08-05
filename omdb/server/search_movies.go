package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/widjaya-hernando/omdb-service/gen"
	"github.com/widjaya-hernando/omdb-service/models"
	"github.com/widjaya-hernando/omdb-service/omdb/repository"
)

func (s *Server) SearchMovies(ctx context.Context, req *pb.SearchMoviesRequest) (*pb.SearchMoviesResponse, error) {
	if req.Query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}
	if req.Type == "" {
		return nil, fmt.Errorf("type cannot be empty")
	}
	if req.Page <= 0 {
		return nil, fmt.Errorf("page must be bigger than 0")
	}

	searchMoviesRequest := models.SearchMoviesRequest{
		Query: req.Query,
		Type:  req.Type,
		Page:  req.Page,
	}

	resp, err := repository.SearchMovies(searchMoviesRequest)
	if err != nil {
		log.Println("err-search-movies: ", err)
		return nil, fmt.Errorf(err.Error())
	}

	response := *models.SearchMoviesResponseToGrpcMessage(resp)

	return &response, nil
}
