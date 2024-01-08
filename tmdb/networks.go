package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworksService interface {
	GetDetails(ctx context.Context, networkID int) (*NetworkDetailsResponse, error)
	GetAlternativeNames(ctx context.Context, networkID int) (*NetworkAlternativeNamesResponse, error)
	GetImages(ctx context.Context, networkID int) (*NetworkImagesResponse, error)
}

type NetworksClient struct {
	baseClient *Client
}

type NetworkDetailsResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type NetworkAlternativeNamesResponse struct {
	ID      int `json:"id"`
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

type NetworkImagesResponse struct {
	ID    int `json:"id"`
	Logos []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		FileType    string  `json:"file_type"`
		Height      int     `json:"height"`
		ID          string  `json:"id"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"logos"`
}

func (nc *NetworksClient) GetDetails(ctx context.Context, networkID int) (*NetworkDetailsResponse, error) {
	resp, err := nc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/network/%d", networkID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result NetworkDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (nc *NetworksClient) GetAlternativeNames(ctx context.Context, networkID int) (*NetworkAlternativeNamesResponse, error) {
	resp, err := nc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/network/%d/alternative_names", networkID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result NetworkAlternativeNamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (nc *NetworksClient) GetImages(ctx context.Context, networkID int) (*NetworkImagesResponse, error) {
	resp, err := nc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/network/%d/images", networkID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result NetworkImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
