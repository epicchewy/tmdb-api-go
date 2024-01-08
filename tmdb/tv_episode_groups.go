package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TvEpisodeGroupsService interface {
	GetDetails(ctx context.Context, tvEpisodeGroupId string) (*TvEpisodeGroupDetailsResponse, error)
}

type TvEpisodeGroupsClient struct {
	baseClient *Client
}

type TvEpisodeGroupDetailsResponse struct {
	Description  string `json:"description"`
	EpisodeCount int32  `json:"episode_count"`
	GroupCount   int32  `json:"group_count"`
	Groups       []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Order    int32  `json:"order"`
		Episodes []struct {
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
		} `json:"episodes"`
		Locked bool `json:"locked"`
	} `json:"groups"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Network struct {
		ID            int32  `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"network"`
	Type int32 `json:"type"`
}

func (tc *TvEpisodeGroupsClient) GetDetails(ctx context.Context, tvEpisodeGroupId string) (*TvEpisodeGroupDetailsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/episode_group/%s", tvEpisodeGroupId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodeGroupDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
