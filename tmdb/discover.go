package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type DiscoverService interface {
	GetMovies(ctx context.Context, queryParams ...queryParam) (*DiscoverMoviesResponse, error)
	GetTVShows(ctx context.Context, queryParams ...queryParam) (*DiscoverTVShowsResponse, error)
}

type DiscoverClient struct {
	baseClient *Client
}

type DiscoverMoviesResponse struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Results      []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int32 `json:"genre_ids"`
		ID               int32   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int32   `json:"vote_count"`
	} `json:"results"`
}

type DiscoverTVShowsResponse struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Results      []struct {
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIds         []int32  `json:"genre_ids"`
		ID               int32    `json:"id"`
		Name             string   `json:"name"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int32    `json:"vote_count"`
	} `json:"results"`
}

func (dc *DiscoverClient) GetMovies(ctx context.Context, queryParams ...queryParam) (*DiscoverMoviesResponse, error) {
	reps, err := dc.baseClient.request(ctx, http.MethodGet, "/discover/movie", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var result DiscoverMoviesResponse
	if err := json.NewDecoder(reps.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (dc *DiscoverClient) GetTVShows(ctx context.Context, queryParams ...queryParam) (*DiscoverTVShowsResponse, error) {
	reps, err := dc.baseClient.request(ctx, http.MethodGet, "/discover/tv", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var result DiscoverTVShowsResponse
	if err := json.NewDecoder(reps.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
