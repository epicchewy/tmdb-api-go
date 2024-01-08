package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type FindService interface {
	FindByID(ctx context.Context, externalId string, externalSource queryParam, queryParams ...queryParam) (*FindResponse, error)
}

type FindClient struct {
	baseClient *Client
}

type FindResponse struct {
	MovieResults []struct {
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
	} `json:"movie_results"`

	// TODO: figure out what these are
	PersonResults    []interface{} `json:"person_results"`
	TvResults        []interface{} `json:"tv_results"`
	TvEpisodeResults []interface{} `json:"tv_episode_results"`
	TvSeasonResults  []interface{} `json:"tv_season_results"`
}

func (fc *FindClient) FindByID(ctx context.Context, externalId string, externalSource queryParam, queryParams ...queryParam) (*FindResponse, error) {
	resp, err := fc.baseClient.request(ctx, http.MethodGet, "/find/"+externalId, append(queryParams, externalSource)...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var findResponse FindResponse
	if err := json.NewDecoder(resp.Body).Decode(&findResponse); err != nil {
		return nil, err
	}
	return &findResponse, nil
}
