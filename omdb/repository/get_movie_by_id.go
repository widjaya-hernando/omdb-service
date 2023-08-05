package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/widjaya-hernando/omdb-service/models"
)

func GetMovieByID(ID string) (*models.GetMovieByIDResponse, error) {
	// URL to make the HTTP GET request to
	url := "https://www.omdbapi.com/?apikey=faf7e5bb&i=%v"

	// Perform the HTTP GET request
	resp, err := http.Get(fmt.Sprintf(url, ID))
	if err != nil {
		fmt.Println("Error making the request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", resp.Status)
		return nil, err
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var resps models.GetMovieByIDResponse

	err = json.Unmarshal(body, &resps)
	if err != nil {
		fmt.Println("Error Parsing JSON:", err)
		return nil, err
	}
	if resps.Response != "True" {
		fmt.Println("Error API:", resps.Error)
		return nil, fmt.Errorf(resps.Error)
	}

	return &resps, nil
}
