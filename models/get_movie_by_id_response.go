package models

import (
	"strings"

	pb "github.com/widjaya-hernando/omdb-service/gen"
)

type GetMovieByIDResponse struct {
	Response  string `json:"response"`
	Error     string `json:"error"`
	ID        string `json:"imdbid"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	Rated     string `json:"rated"`
	Genre     string `json:"genre"`
	Plot      string `json:"plot"`
	Actors    string `json:"actors"`
	Language  string `json:"language"`
	Country   string `json:"country"`
	Type      string `json:"type"`
	PosterUrl string `json:"poster"`
}

func GetMovieByIDResponseToGrpcMessage(getMovieByIDResponse *GetMovieByIDResponse) *pb.GetMovieByIDResponse {
	return &pb.GetMovieByIDResponse{
		Id:        getMovieByIDResponse.ID,
		Title:     getMovieByIDResponse.Title,
		Year:      getMovieByIDResponse.Year,
		Rated:     getMovieByIDResponse.Rated,
		Genre:     getMovieByIDResponse.Genre,
		Plot:      getMovieByIDResponse.Plot,
		Actors:    strings.Split(getMovieByIDResponse.Actors, ", "),
		Language:  getMovieByIDResponse.Language,
		Country:   getMovieByIDResponse.Country,
		Type:      getMovieByIDResponse.Type,
		PosterUrl: getMovieByIDResponse.PosterUrl,
	}
}
