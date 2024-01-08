package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type GenresService interface {
	GetMovieGenres(ctx context.Context, queryParams ...queryParam) (*GenreList, error)
	GetTVGenres(ctx context.Context, queryParams ...queryParam) (*GenreList, error)
}

type GenreClient struct {
	baseClient *Client
}

type GenreList struct {
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

func (gc *GenreClient) GetMovieGenres(ctx context.Context, queryParams ...queryParam) (*GenreList, error) {
	resp, err := gc.baseClient.request(ctx, http.MethodGet, "/genre/movie/list", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GenreList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (gc *GenreClient) GetTVGenres(ctx context.Context, queryParams ...queryParam) (*GenreList, error) {
	resp, err := gc.baseClient.request(ctx, http.MethodGet, "/genre/tv/list", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GenreList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
