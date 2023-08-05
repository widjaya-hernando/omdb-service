package models

import (
	"strconv"

	pb "github.com/widjaya-hernando/omdb-service/gen"
)

type MovieResult struct {
	ID        string `json:"imdbid"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	Type      string `json:"type"`
	PosterUrl string `json:"poster"`
}

type SearchMoviesResponse struct {
	Response     string        `json:"response"`
	Error        string        `json:"error"`
	Movies       []MovieResult `json:"search"`
	TotalResults string        `json:"totalResults"`
}

func SearchMoviesResponseToGrpcMessage(searchMoviesResponse *SearchMoviesResponse) *pb.SearchMoviesResponse {
	var movies []*pb.MovieResult

	for _, movie := range searchMoviesResponse.Movies {
		movies = append(movies, &pb.MovieResult{
			Id:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Type:      movie.Type,
			PosterUrl: movie.PosterUrl,
		})
	}

	totalResults, _ := strconv.ParseUint(searchMoviesResponse.TotalResults, 10, 64)

	return &pb.SearchMoviesResponse{
		Movies:       movies,
		TotalResults: totalResults,
	}
}
