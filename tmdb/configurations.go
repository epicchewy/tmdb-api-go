package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ConfigurationsService interface {
	GetDetails(ctx context.Context) (*ConfigurationDetails, error)
	GetCountries(ctx context.Context, queryParams ...queryParam) (*ConfigurationCountries, error)
	GetJobs(ctx context.Context) (*ConfigurationJobs, error)
	GetLanguages(ctx context.Context) (*ConfigurationLanguages, error)
	GetPrimaryTranslations(ctx context.Context) (*ConfigurationPrimaryTranslations, error)
	GetTimezones(ctx context.Context) (*ConfigurationTimezones, error)
}

type ConfigurationsClient struct {
	baseClient *Client
}

type ConfigurationDetails struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

type ConfigurationCountries struct {
	Countries []struct {
		Iso_3166_1  string `json:"iso_3166_1"`
		EnglishName string `json:"english_name"`
		NativeName  string `json:"native_name"`
	} `json:"countries"`
}

type Job struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

type ConfigurationJobs []Job

type Language struct {
	Iso_639_1   string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

type ConfigurationLanguages []Language

type ConfigurationPrimaryTranslations []string

type Timezone struct {
	Iso_3166_1 string   `json:"iso_3166_1"`
	Zones      []string `json:"zones"`
}

type ConfigurationTimezones []Timezone

func (cc *ConfigurationsClient) GetDetails(ctx context.Context) (*ConfigurationDetails, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationDetails
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ConfigurationsClient) GetCountries(ctx context.Context, queryParams ...queryParam) (*ConfigurationCountries, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration/countries", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationCountries
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ConfigurationsClient) GetJobs(ctx context.Context) (*ConfigurationJobs, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration/jobs")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationJobs
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ConfigurationsClient) GetLanguages(ctx context.Context) (*ConfigurationLanguages, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration/languages")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationLanguages
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ConfigurationsClient) GetPrimaryTranslations(ctx context.Context) (*ConfigurationPrimaryTranslations, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration/primary_translations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationPrimaryTranslations
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *ConfigurationsClient) GetTimezones(ctx context.Context) (*ConfigurationTimezones, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/configuration/timezones")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ConfigurationTimezones
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
