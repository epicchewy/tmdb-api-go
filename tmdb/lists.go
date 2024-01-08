package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListsService interface {
	CheckItemStatus(ctx context.Context, listID string, queryParams ...queryParam) (*ListItemStatusResponse, error)
	GetDetails(ctx context.Context, listID string, queryParams ...queryParam) (*ListDetailsResponse, error)
}

type ListsClient struct {
	baseClient *Client
}

type ListItemStatusResponse struct {
	ID          int  `json:"id"`
	ItemPresent bool `json:"item_present"`
}

type ListDetailsResponse struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count"`
	ID            int    `json:"id"`
	Items         []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int32 `json:"genre_ids"`
		ID               int32   `json:"id"`
		MediaType        string  `json:"media_type"`
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
	} `json:"items"`
	ItemCount  int    `json:"item_count"`
	Iso_639_1  string `json:"iso_639_1"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

func (lc *ListsClient) CheckItemStatus(ctx context.Context, listID string, queryParams ...queryParam) (*ListItemStatusResponse, error) {
	reps, err := lc.baseClient.request(ctx, http.MethodGet, "/list/"+listID+"/item_status", queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var listItemStatusResponse ListItemStatusResponse
	if err := json.NewDecoder(reps.Body).Decode(&listItemStatusResponse); err != nil {
		return nil, err
	}
	return &listItemStatusResponse, nil
}

func (lc *ListsClient) GetDetails(ctx context.Context, listID string, queryParams ...queryParam) (*ListDetailsResponse, error) {
	reps, err := lc.baseClient.request(ctx, http.MethodGet, "/list/"+listID, queryParams...)
	if err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	var listDetailsResponse ListDetailsResponse
	if err := json.NewDecoder(reps.Body).Decode(&listDetailsResponse); err != nil {
		return nil, err
	}
	return &listDetailsResponse, nil
}
