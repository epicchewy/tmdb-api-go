package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TrendingService interface {
	GetAll(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingAllResponse, error)
	GetMovies(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingMoviesResponse, error)
	GetTvShows(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingTvShowsResponse, error)
	GetPeople(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingPeopleResponse, error)
}

type TrendingClient struct {
	baseClient *Client
}

type TrendingResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int32 `json:"genre_ids"`
	ID               int32   `json:"id"`
	MediaType        string  `json:"media_type"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int32   `json:"vote_count"`
}

type TrendingAllResponse struct {
	Page         int32            `json:"page"`
	TotalPages   int32            `json:"total_pages"`
	TotalResults int32            `json:"total_results"`
	Results      []TrendingResult `json:"results"`
}

type TrendingMoviesResponse struct {
	Page         int32            `json:"page"`
	TotalPages   int32            `json:"total_pages"`
	TotalResults int32            `json:"total_results"`
	Results      []TrendingResult `json:"results"`
}

type TrendingTvShowsResponse struct {
	Page         int32 `json:"page"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
	Results      []struct {
		FirstAirDate  string   `json:"first_air_date"`
		OriginCountry []string `json:"origin_country"`
		TrendingResult
	} `json:"results"`
}

type TrendingPeopleResponse struct {
	Page         int32 `json:"page"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
	Results      []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		Name               string  `json:"name"`
		MediaType          string  `json:"media_type"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		KnownForDepartment string  `json:"known_for_department"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int32 `json:"genre_ids"`
			ID               int32   `json:"id"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			Popularity       float32 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float32 `json:"vote_average"`
			VoteCount        int32   `json:"vote_count"`
		} `json:"known_for"`
	} `json:"results"`
}

func (t *TrendingClient) GetAll(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingAllResponse, error) {
	resp, err := t.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/trending/all/%s", timeWindow), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TrendingAllResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TrendingClient) GetMovies(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingMoviesResponse, error) {
	resp, err := t.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/trending/movie/%s", timeWindow), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TrendingMoviesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TrendingClient) GetTvShows(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingTvShowsResponse, error) {
	resp, err := t.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/trending/tv/%s", timeWindow), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TrendingTvShowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TrendingClient) GetPeople(ctx context.Context, timeWindow string, queryParams ...queryParam) (*TrendingPeopleResponse, error) {
	resp, err := t.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/trending/people/%s", timeWindow), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TrendingPeopleResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
