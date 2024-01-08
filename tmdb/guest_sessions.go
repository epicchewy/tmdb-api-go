package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GuestSessionsService interface {
	GetRatedMovies(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedMoviesResponse, error)
	GetRatedTVShows(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedTvShowsResponse, error)
	GetRatedTVEpisodes(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedTvShowEpisodesResponse, error)
}

type GuestSessionsClient struct {
	baseClient *Client
}

type RatedMoviesResponse struct {
	Page    int32 `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int32 `json:"genre_ids"`
		ID               int32   `json:"id"`
		OriginalCountry  string  `json:"original_country"`
		OriginalLanguage string  `json:"original_language"`
		OriginalName     string  `json:"original_name"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		FirstAirDate     string  `json:"first_air_date"`
		Name             string  `json:"name"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int32   `json:"vote_count"`
		Rating           float64 `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type RatedTvShowsResponse struct {
	Page    int32 `json:"page"`
	Results []struct {
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
		Rating           float64 `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type RatedTvShowEpisodesResponse struct {
	Page    int32 `json:"page"`
	Results []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int32   `json:"episode_number"`
		ID             int32   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		Runtime        int32   `json:"runtime"`
		SeasonNumber   int32   `json:"season_number"`
		ShowID         int32   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int32   `json:"vote_count"`
		Rating         float64 `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

func (gc *GuestSessionsClient) GetRatedMovies(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedMoviesResponse, error) {
	resp, err := gc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/guest_session/%d/rated/movies", guestSessionId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedMoviesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (gc *GuestSessionsClient) GetRatedTVShows(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedTvShowsResponse, error) {
	resp, err := gc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/guest_session/%d/rated/tv", guestSessionId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedTvShowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (gc *GuestSessionsClient) GetRatedTVEpisodes(ctx context.Context, guestSessionId int32, queryParams ...queryParam) (*RatedTvShowEpisodesResponse, error) {
	resp, err := gc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/guest_session/%d/rated/tv/episodes", guestSessionId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedTvShowEpisodesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
