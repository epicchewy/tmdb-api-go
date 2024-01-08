package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type MovieListsService interface {
	GetNowPlaying(ctx context.Context, queryParams ...queryParam) (*MoviesNowPlayingResponse, error)
	GetPopular(ctx context.Context, queryParams ...queryParam) (*MoviesPopularResponse, error)
	GetTopRated(ctx context.Context, queryParams ...queryParam) (*MoviesTopRatedResponse, error)
	GetUpcoming(ctx context.Context, queryParams ...queryParam) (*MoviesUpcomingResponse, error)
}

type MovieListsClient struct {
	baseClient *Client
}

type MovieListMovie struct {
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
}

type MoviesNowPlayingResponse struct {
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	Page         int              `json:"page"`
	Results      []MovieListMovie `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type MoviesPopularResponse struct {
	Page         int              `json:"page"`
	Results      []MovieListMovie `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type MoviesTopRatedResponse struct {
	Page         int              `json:"page"`
	Results      []MovieListMovie `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type MoviesUpcomingResponse struct {
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	Page         int              `json:"page"`
	Results      []MovieListMovie `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

func (mlc *MovieListsClient) GetNowPlaying(ctx context.Context, queryParams ...queryParam) (*MoviesNowPlayingResponse, error) {
	resp, err := mlc.baseClient.request(ctx, http.MethodGet, "/movie/now_playing", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var moviesNowPlayingResponse MoviesNowPlayingResponse
	if err := json.NewDecoder(resp.Body).Decode(&moviesNowPlayingResponse); err != nil {
		return nil, err
	}
	return &moviesNowPlayingResponse, nil
}

func (mlc *MovieListsClient) GetPopular(ctx context.Context, queryParams ...queryParam) (*MoviesPopularResponse, error) {
	resp, err := mlc.baseClient.request(ctx, http.MethodGet, "/movie/popular", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var moviesPopularResponse MoviesPopularResponse
	if err := json.NewDecoder(resp.Body).Decode(&moviesPopularResponse); err != nil {
		return nil, err
	}
	return &moviesPopularResponse, nil
}

func (mlc *MovieListsClient) GetTopRated(ctx context.Context, queryParams ...queryParam) (*MoviesTopRatedResponse, error) {
	resp, err := mlc.baseClient.request(ctx, http.MethodGet, "/movie/top_rated", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var moviesTopRatedResponse MoviesTopRatedResponse
	if err := json.NewDecoder(resp.Body).Decode(&moviesTopRatedResponse); err != nil {
		return nil, err
	}
	return &moviesTopRatedResponse, nil
}

func (mlc *MovieListsClient) GetUpcoming(ctx context.Context, queryParams ...queryParam) (*MoviesUpcomingResponse, error) {
	resp, err := mlc.baseClient.request(ctx, http.MethodGet, "/movie/upcoming", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var moviesUpcomingResponse MoviesUpcomingResponse
	if err := json.NewDecoder(resp.Body).Decode(&moviesUpcomingResponse); err != nil {
		return nil, err
	}
	return &moviesUpcomingResponse, nil
}
