package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type TvSeriesListsService interface {
	GetAiringToday(ctx context.Context, queryParams ...queryParam) (*TvSeriesAiringTodayResponse, error)
	GetOnTheAir(ctx context.Context, queryParams ...queryParam) (*TvSeriesOnTheAirResponse, error)
	GetPopular(ctx context.Context, queryParams ...queryParam) (*TvSeriesPopularResponse, error)
	GetTopRated(ctx context.Context, queryParams ...queryParam) (*TvSeriesTopRatedResponse, error)
}

type TvSeriesListsClient struct {
	baseClient *Client
}

type TvSeriesResult struct {
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
}

type TvSeriesAiringTodayResponse struct {
	Page         int              `json:"page"`
	Results      []TvSeriesResult `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type TvSeriesOnTheAirResponse struct {
	Page         int              `json:"page"`
	Results      []TvSeriesResult `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type TvSeriesPopularResponse struct {
	Page         int              `json:"page"`
	Results      []TvSeriesResult `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type TvSeriesTopRatedResponse struct {
	Page         int              `json:"page"`
	Results      []TvSeriesResult `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

func (tc *TvSeriesListsClient) GetAiringToday(ctx context.Context, queryParams ...queryParam) (*TvSeriesAiringTodayResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, "/tv/airing_today", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesAiringTodayResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesListsClient) GetOnTheAir(ctx context.Context, queryParams ...queryParam) (*TvSeriesOnTheAirResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, "/tv/on_the_air", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesOnTheAirResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesListsClient) GetPopular(ctx context.Context, queryParams ...queryParam) (*TvSeriesPopularResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, "/tv/popular", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesPopularResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesListsClient) GetTopRated(ctx context.Context, queryParams ...queryParam) (*TvSeriesTopRatedResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, "/tv/top_rated", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesTopRatedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
