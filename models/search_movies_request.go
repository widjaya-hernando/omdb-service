package models

import (
	pb "github.com/widjaya-hernando/omdb-service/gen"
)

type SearchMoviesRequest struct {
	Query string `json:"query"`
	Type  string `json:"type"`
	Page  uint64 `json:"page"`
}

func SearchMoviesRequestToGrpcMessage(searchMoviesRequest *SearchMoviesRequest) *pb.SearchMoviesRequest {
	return &pb.SearchMoviesRequest{
		Query: searchMoviesRequest.Query,
		Type:  searchMoviesRequest.Type,
		Page:  searchMoviesRequest.Page,
	}
}
