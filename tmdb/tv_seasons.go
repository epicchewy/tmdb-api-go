package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TvSeasonsService interface {
	GetDetails(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsDetailsResponse, error)
	GetAccountStates(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsAccountStatesResponse, error)
	GetAggregateCredits(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsAggregateCreditsResponse, error)
	GetChanges(ctx context.Context, seasonID int32, queryParams ...queryParam) (*TvSeasonsChangesResponse, error)
	GetCredits(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsCreditsResponse, error)
	GetExternalIds(ctx context.Context, seriesID int32, seasonNumber int32) (*TvSeasonsExternalIdsResponse, error)
	GetImages(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsImagesResponse, error)
	GetTranslations(ctx context.Context, seriesID int32, seasonNumber int32) (*TvSeasonsTranslationsResponse, error)
	GetVideos(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsVideosResponse, error)
	// GetWatchProviders(ctx context.Context, seasonID int32, queryParams ...queryParam) (*TvSeasonsWatchProvidersResponse, error)
}

type TvSeasonsClient struct {
	baseClient *Client
}

// TvSeasonsDetailsResponse struct is based off of https://developer.themoviedb.org/reference/tv-season-details
type TvSeasonsDetailsResponse struct {
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate       string `json:"air_date"`
		EpisodeNumber int32  `json:"episode_number"`
		Crew          []struct {
			Department         string  `json:"department"`
			Job                string  `json:"job"`
			CreditID           string  `json:"credit_id"`
			Adult              bool    `json:"adult"`
			Gender             int32   `json:"gender"`
			ID                 int32   `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float32 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
		} `json:"crew"`
		GuestStars []struct {
			CreditID           string  `json:"credit_id"`
			Adult              bool    `json:"adult"`
			Gender             int32   `json:"gender"`
			ID                 int32   `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float32 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			Character          string  `json:"character"`
			Order              int32   `json:"order"`
		} `json:"guest_stars"`
		ID             int32   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		PosterPath     string  `json:"poster_path"`
		Runtime        int32   `json:"runtime"`
		SeasonNumber   int32   `json:"season_number"`
		ShowID         int32   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int32   `json:"vote_count"`
	} `json:"episodes"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	ID           int32   `json:"id"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int32   `json:"season_number"`
	VoteAverage  float32 `json:"vote_average"`
}

// TvSeasonsAccountStatesResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-account-states
type TvSeasonsAccountStatesResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID            string `json:"id"`
		EpisodeNumber int32  `json:"episode_number"`
		Rated         struct {
			Value float32 `json:"value"`
		} `json:"rated"`
	} `json:"results"`
}

// TvSeasonsAggregateCreditsResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-aggregate-credits
type TvSeasonsAggregateCreditsResponse struct {
	ID   int32 `json:"id"`
	Cast []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Roles              []struct {
			CreditID     string `json:"credit_id"`
			Character    string `json:"character"`
			EpisodeCount int32  `json:"episode_count"`
		} `json:"roles"`
		TotalEpisodeCount int32 `json:"total_episode_count"`
		Order             int32 `json:"order"`
	} `json:"cast"`
	Crew []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Jobs               []struct {
			CreditID     string `json:"credit_id"`
			Job          string `json:"job"`
			EpisodeCount int32  `json:"episode_count"`
		} `json:"jobs"`
		Department        string `json:"department"`
		TotalEpisodeCount int32  `json:"total_episode_count"`
	} `json:"crew"`
}

// TvSeasonsChangesResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-changes
type TvSeasonsChangesResponse struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  struct {
				EpisodeID     int32 `json:"episode_id"`
				EpisodeNumber int32 `json:"episode_number"`
			} `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// TvSeasonsCreditsResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-credits
type TvSeasonsCreditsResponse struct {
	ID   int32 `json:"id"`
	Cast []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Order              int32   `json:"order"`
	} `json:"cast"`
	Crew []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditID           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Job                string  `json:"job"`
	} `json:"crew"`
}

// TvSeasonsExternalIdsResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-external-ids
type TvSeasonsExternalIdsResponse struct {
	ID          int32  `json:"id"`
	FreebaseID  string `json:"freebase_id"`
	FreebaseMid string `json:"freebase_mid"`
	TvdbID      int32  `json:"tvdb_id"`
	TvrageID    string `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
}

// TvSeasonsImagesResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-images
type TvSeasonsImagesResponse struct {
	ID      int32 `json:"id"`
	Posters []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int32   `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int32   `json:"vote_count"`
		Width       int32   `json:"width"`
	} `json:"posters"`
}

// TvSeasonsTranslationsResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-translations
type TvSeasonsTranslationsResponse struct {
	ID           int32 `json:"id"`
	Translations []struct {
		Iso_3166_1  string `json:"iso_3166_1"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        []struct {
			Name     string `json:"name"`
			Overview string `json:"overview"`
		} `json:"data"`
	} `json:"translations"`
}

// TvSeasonsVideosResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/tv-season-videos
type TvSeasonsVideosResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID         string `json:"id"`
		Iso_639_1  string `json:"iso_639_1"`
		Iso_3166_1 string `json:"iso_3166_1"`
		Key        string `json:"key"`
		Name       string `json:"name"`
		Site       string `json:"site"`
		Size       int32  `json:"size"`
		Type       string `json:"type"`
		Official   bool   `json:"official"`
		PublishAt  string `json:"publish_at"`
	} `json:"results"`
}

// TvSeasonsWatchProvidersResponse struct is based off of https://developers.themoviedb.org/3/tv-seasons/get-tv-season-watch-providers
// TODO: figure this out

func (tc *TvSeasonsClient) GetDetails(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsDetailsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d", seriesID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetAccountStates(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsAccountStatesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/account_states", seriesID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsAccountStatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetAggregateCredits(ctx context.Context, seasonID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsAggregateCreditsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/aggregate_credits", seasonID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsAggregateCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetChanges(ctx context.Context, seasonID int32, queryParams ...queryParam) (*TvSeasonsChangesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/changes", seasonID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsChangesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetCredits(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsCreditsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/credits", seriesID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetExternalIds(ctx context.Context, seriesID int32, seasonNumber int32) (*TvSeasonsExternalIdsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/external_ids", seriesID, seasonNumber))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsExternalIdsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetImages(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsImagesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/images", seriesID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetTranslations(ctx context.Context, seriesID int32, seasonNumber int32) (*TvSeasonsTranslationsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/translations", seriesID, seasonNumber))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsTranslationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeasonsClient) GetVideos(ctx context.Context, seriesID int32, seasonNumber int32, queryParams ...queryParam) (*TvSeasonsVideosResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/videos", seriesID, seasonNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeasonsVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
