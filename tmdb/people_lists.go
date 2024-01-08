package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type PeopleListsService interface {
	GetPopular(ctx context.Context, queryParams ...queryParam) (*PeopleListPopularResponse, error)
}

type PeopleListsClient struct {
	baseClient *Client
}

type PeopleListPopularResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult    bool  `json:"adult"`
		Gender   int32 `json:"gender"`
		ID       int32 `json:"id"`
		KnownFor []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int32 `json:"genre_ids"`
			ID               int32   `json:"id"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int32   `json:"vote_count"`
		} `json:"known_for"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func (pc *PeopleListsClient) GetPopular(ctx context.Context, queryParams ...queryParam) (*PeopleListPopularResponse, error) {
	resp, err := pc.baseClient.request(ctx, http.MethodGet, "/person/popular", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result PeopleListPopularResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
