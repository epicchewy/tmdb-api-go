package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreditsService interface {
	GetDetails(ctx context.Context, creditId int32) (*CreditDetailsResponse, error)
}

type CreditsClient struct {
	baseClient *Client
}

type CreditDetailsResponse struct {
	ID         int32  `json:"id"`
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	MediaType  string `json:"media_type"`
	Media      struct {
		Adult        bool   `json:"adult"`
		BackdropPath string `json:"backdrop_path"`
		Character    string `json:"character"`
		// Episodes array of ?
		FirstAirDate     string  `json:"first_air_date"`
		GenreIds         []int32 `json:"genre_ids"`
		ID               int32   `json:"id"`
		MediaType        string  `json:"media_type"`
		Name             string  `json:"name"`
		OriginalLanguage string  `json:"original_language"`
		OriginalName     string  `json:"original_name"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int32   `json:"vote_count"`
		Seasons          []struct {
			AirDate      string `json:"air_date"`
			EpisodeCount int32  `json:"episode_count"`
			ID           int32  `json:"id"`
			Name         string `json:"name"`
			Overview     string `json:"overview"`
			PosterPath   string `json:"poster_path"`
			SeasonNumber int32  `json:"season_number"`
			ShowID       int32  `json:"show_id"`
		} `json:"seasons"`
	} `json:"media"`
	Person struct {
		Adult              bool    `json:"adult"`
		ID                 int32   `json:"id"`
		Name               string  `json:"name"`
		Gender             int32   `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		MediaType          string  `json:"media_type"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"person"`
}

func (cc *CreditsClient) GetDetails(ctx context.Context, creditId int32) (*CreditDetailsResponse, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/credit/%d", creditId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CreditDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
