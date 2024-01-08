package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TvEpisodesService interface {
	GetDetails(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesDetailsResponse, error)
	GetAccountStates(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesAccountStatesResponse, error)
	GetChanges(ctx context.Context, episodeID int32) (*TvEpisodesChangesResponse, error)
	GetCredits(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesCreditsResponse, error)
	GetExternalIDs(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32) (*TvEpisodesExternalIDsResponse, error)
	GetImages(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesImagesResponse, error)
	GetTranslations(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32) (*TvEpisodesTranslationsResponse, error)
	GetVideos(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesVideosResponse, error)
}

type TvEpisodesClient struct {
	baseClient *Client
}

// TvEpisodesDetailsResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-details
type TvEpisodesDetailsResponse struct {
	AirDate       string `json:"air_date"`
	EpisodeNumber int32  `json:"episode_number"`
	Crew          []struct {
		ID                 int32   `json:"id"`
		CreditID           string  `json:"credit_id"`
		Name               string  `json:"name"`
		Department         string  `json:"department"`
		Job                string  `json:"job"`
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		OrignalName        string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
	GuestStars []struct {
		ID                 int32   `json:"id"`
		CreditID           string  `json:"credit_id"`
		Name               string  `json:"name"`
		Character          string  `json:"character"`
		Order              int32   `json:"order"`
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		OrignalName        string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"guest_stars"`
	ID             int32   `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int32   `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int32   `json:"vote_count"`
	Runtime        int32   `json:"runtime"`
}

// TvEpisodesAccountStatesResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-account-states
type TvEpisodesAccountStatesResponse struct {
	ID       int32 `json:"id"`
	Favorite bool  `json:"favorite"`
	Rated    struct {
		Value float32 `json:"value"`
	} `json:"rated"`
	Watchlist bool `json:"watchlist"`
}

// TvEpisodesChangesResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-changes-by-id
type TvEpisodesChangesResponse struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  string `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// TvEpisodesCreditsResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-credits
type TvEpisodesCreditsResponse struct {
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
	GuestStars []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditID           string  `json:"credit_id"`
		Character          string  `json:"character"`
		Order              int32   `json:"order"`
	} `json:"guest_stars"`
}

// TvEpisodesExternalIDsResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-external-ids
type TvEpisodesExternalIDsResponse struct {
	ID          int32  `json:"id"`
	IMDBID      string `json:"imdb_id"`
	TVDBID      int32  `json:"tvdb_id"`
	TVRageID    int32  `json:"tvrage_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	WikipediaID int32  `json:"wikipedia_id"`
}

// TvEpisodesImagesResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-images
type TvEpisodesImagesResponse struct {
	ID     int32 `json:"id"`
	Stills []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int32   `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int32   `json:"vote_count"`
		Width       int32   `json:"width"`
	} `json:"stills"`
}

// TvEpisodesTranslationsResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-translations
type TvEpisodesTranslationsResponse struct {
	ID           int32 `json:"id"`
	Translations []struct {
		Iso_639_1   string `json:"iso_639_1"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Name     string `json:"name"`
			Overview string `json:"overview"`
		} `json:"data"`
	} `json:"translations"`
}

// TvEpisodesVideosResponse struct is based off of https://developer.themoviedb.org/reference/tv-episode-videos
type TvEpisodesVideosResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID          string `json:"id"`
		Iso_639_1   string `json:"iso_639_1"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Key         string `json:"key"`
		Official    bool   `json:"official"`
		PublishedAt string `json:"published_at"`
		Name        string `json:"name"`
		Site        string `json:"site"`
		Size        int32  `json:"size"`
		Type        string `json:"type"`
	} `json:"results"`
}

func (tc *TvEpisodesClient) GetDetails(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesDetailsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d", seriesID, seasonNumber, episodeNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetAccountStates(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesAccountStatesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/account_states", seriesID, seasonNumber, episodeNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesAccountStatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetChanges(ctx context.Context, episodeID int32) (*TvEpisodesChangesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/episode/%d/changes", episodeID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesChangesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetCredits(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesCreditsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/credits", seriesID, seasonNumber, episodeNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetExternalIDs(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32) (*TvEpisodesExternalIDsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/external_ids", seriesID, seasonNumber, episodeNumber))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesExternalIDsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetImages(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesImagesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/images", seriesID, seasonNumber, episodeNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetTranslations(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32) (*TvEpisodesTranslationsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/translations", seriesID, seasonNumber, episodeNumber))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesTranslationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvEpisodesClient) GetVideos(ctx context.Context, seriesID int32, seasonNumber int32, episodeNumber int32, queryParams ...queryParam) (*TvEpisodesVideosResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/season/%d/episode/%d/videos", seriesID, seasonNumber, episodeNumber), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvEpisodesVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
