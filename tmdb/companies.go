package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CompaniesService interface {
	GetDetails(ctx context.Context, companyId int32) (*Company, error)
	GetAlternativeNames(ctx context.Context, companyId int32) (*CompanyAlternativeNames, error)
	GetImages(ctx context.Context, companyId int32) (*CompanyImages, error)
}

type CompaniesClient struct {
	baseClient *Client
}

type Company struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
	ParentCompany string `json:"parent_company"`
}

type CompanyAlternativeNames struct {
	ID      int `json:"id"`
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

type CompanyImages struct {
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

func (cc *CompaniesClient) GetDetails(ctx context.Context, companyId int32) (*Company, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/company/%d", companyId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Company
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *CompaniesClient) GetAlternativeNames(ctx context.Context, companyId int32) (*CompanyAlternativeNames, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/company/%d/alternative_names", companyId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CompanyAlternativeNames
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *CompaniesClient) GetImages(ctx context.Context, companyId int32) (*CompanyImages, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/company/%d/images", companyId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CompanyImages
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
