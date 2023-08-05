package models

import (
	pb "github.com/widjaya-hernando/omdb-service/gen"
)

type GetMovieByIDRequest struct {
	ID string `json:"id"`
}

func GetMovieByIDRequestToGrpcMessage(product *GetMovieByIDRequest) *pb.GetMovieByIDRequest {
	return &pb.GetMovieByIDRequest{
		Id: product.ID,
	}
}
