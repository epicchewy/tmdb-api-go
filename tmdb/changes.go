package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ChangesService interface {
	GetMovieChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error)
	GetTVChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error)
	GetPersonChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error)
}

type ChangesClient struct {
	baseClient *Client
}

type Changes struct {
	Changes []struct {
		ID    string `json:"id"`
		Adult bool   `json:"adult"`
	} `json:"changes"`
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func (cc *ChangesClient) GetMovieChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error) {
	reps, err := cc.baseClient.request(ctx, http.MethodGet, "/movie/changes", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var result Changes
	if err := json.NewDecoder(reps.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ChangesClient) GetTVChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error) {
	reps, err := cc.baseClient.request(ctx, http.MethodGet, "/tv/changes", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var result Changes
	if err := json.NewDecoder(reps.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ChangesClient) GetPersonChanges(ctx context.Context, queryParams ...queryParam) (*Changes, error) {
	reps, err := cc.baseClient.request(ctx, http.MethodGet, "/person/changes", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var result Changes
	if err := json.NewDecoder(reps.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
