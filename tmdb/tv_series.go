package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TvSeriesService interface {
	GetDetails(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesDetailsResponse, error)
	GetAccountStates(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesAccountStatesResponse, error)
	GetAggregateCredits(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesAggregateCreditsResponse, error)
	GetAlternativeTitles(ctx context.Context, seriesID int32) (*TvSeriesAlternativeTitlesResponse, error)
	GetChanges(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesChangesResponse, error)
	GetContentRatings(ctx context.Context, seriesID int32) (*TvSeriesContentRatingsResponse, error)
	GetCredits(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesCreditsResponse, error)
	GetEpisodeGroups(ctx context.Context, seriesID int32) (*TvSeriesEpisodeGroupsResponse, error)
	GetExternalIds(ctx context.Context, seriesID int32) (*TvSeriesExternalIdsResponse, error)
	GetImages(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesImagesResponse, error)
	GetKeywords(ctx context.Context, seriesID int32) (*TvSeriesKeywordsResponse, error)
	GetLatest(ctx context.Context) (*TvSeriesLatestResponse, error)
	GetRecommendations(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesRecommendationsResponse, error)
	GetReviews(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesReviewsResponse, error)
	GetScreenedTheatrically(ctx context.Context, seriesID int32) (*TvSeriesScreenedTheatricallyResponse, error)
	GetSimilar(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesSimilarResponse, error)
	GetTranslations(ctx context.Context, seriesID int32) (*TvSeriesTranslationsResponse, error)
	GetVideos(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesVideosResponse, error)
	// GetWatchProviders(ctx context.Context, seriesID int32) (*TvSeriesWatchProvidersResponse, error)
}

type TvSeriesClient struct {
	baseClient *Client
}

type TvSeriesDetailsResponse struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int32  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int32  `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime []int32 `json:"episode_run_time"`
	FirstAirDate   string  `json:"first_air_date"`
	Genres         []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int32    `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
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
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int32   `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int32  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int32    `json:"number_of_episodes"`
	NumberOfSeasons     int32    `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float32  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int32  `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Name       string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string  `json:"air_date"`
		EpisodeCount int32   `json:"episode_count"`
		ID           int32   `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		SeasonNumber int32   `json:"season_number"`
		VoteAverage  float32 `json:"vote_average"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int32   `json:"vote_count"`
}

type TvSeriesAccountStatesResponse struct {
	ID       int32 `json:"id"`
	Favorite bool  `json:"favorite"`
	Rated    struct {
		Value float32 `json:"value"`
	} `json:"rated"`
	Watchlist bool `json:"watchlist"`
}

type TvSeriesAggregateCreditsResponse struct {
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
		TotalEpisodeCount int32  `json:"total_episode_count"`
		Department        string `json:"department"`
	} `json:"crew"`
}

type TvSeriesAlternativeTitlesResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Title      string `json:"title"`
		Type       string `json:"type"`
	} `json:"results"`
}

type TvSeriesChangesResponse struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID         string `json:"id"`
			Action     string `json:"action"`
			Time       string `json:"time"`
			Iso_639_1  string `json:"iso_639_1"`
			Iso_3166_1 string `json:"iso_3166_1"`
			Value      struct {
				Poster struct {
					Filepath  string `json:"file_path"`
					Iso_639_1 string `json:"iso_639_1"`
				} `json:"poster"`
			} `json:"value"`
			OriginalValue struct {
				Poster struct {
					Filepath  string `json:"file_path"`
					Iso_639_1 string `json:"iso_639_1"`
				} `json:"poster"`
			} `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

type TvSeriesContentRatingsResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		Descriptors []struct{} `json:"descriptors"`
		Iso_3166_1  string     `json:"iso_3166_1"`
		Rating      string     `json:"rating"`
	} `json:"results"`
}

type TvSeriesCreditsResponse struct {
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

type TvSeriesEpisodeGroupsResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		Description  string `json:"description"`
		EpisodeCount int32  `json:"episode_count"`
		GroupCount   int32  `json:"group_count"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Network      struct {
			ID            int32  `json:"id"`
			LogoPath      string `json:"logo_path"`
			Name          string `json:"name"`
			OriginCountry string `json:"origin_country"`
		} `json:"network"`
		Type int32 `json:"type"`
	} `json:"results"`
}

type TvSeriesExternalIdsResponse struct {
	ID          int32  `json:"id"`
	IMDBID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int32  `json:"tvdb_id"`
	TvrageID    int32  `json:"tvrage_id"`
	WikidataId  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
}

type TvSeriesImagesResponse struct {
	ID        int32 `json:"id"`
	Backdrops []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int32   `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int32   `json:"vote_count"`
		Width       int32   `json:"width"`
	} `json:"backdrops"`
	Logos []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int32   `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int32   `json:"vote_count"`
		Width       int32   `json:"width"`
	} `json:"logos"`
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

type TvSeriesKeywordsResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
}

type TvSeriesLatestResponse struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	// CreatedBy    []struct {}
	// EpisodeRunTime []int32 `json:"episode_run_time"`
	FirstAirDate string `json:"first_air_date"`
	Genres       []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int32    `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
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
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int32   `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int32  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes int32    `json:"number_of_episodes"`
	NumberOfSeasons  int32    `json:"number_of_seasons"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float32  `json:"popularity"`
	// PosterPath          string   `json:"poster_path"`
	// ProductionCompanies []struct {}
	// ProductionCountries []struct {}
	Seasons []struct {
		AirDate      string  `json:"air_date"`
		EpisodeCount int32   `json:"episode_count"`
		ID           int32   `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		SeasonNumber int32   `json:"season_number"`
		VoteAverage  float32 `json:"vote_average"`
	} `json:"seasons"`
	// SpokenLanguages []struct {}
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int32   `json:"vote_count"`
}

type TvSeriesRecommendationsResponse struct {
	Page         int32 `json:"page"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
	Results      []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIds         []int32  `json:"genre_ids"`
		ID               int32    `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		VoteAverage      float32  `json:"vote_average"`
		VoteCount        int32    `json:"vote_count"`
	} `json:"results"`
}

type TvSeriesReviewsResponse struct {
	ID      int32 `json:"id"`
	Page    int32 `json:"page"`
	Results []struct {
		Author        string `json:"author"`
		AuthorDetails struct {
			Name       string `json:"name"`
			Username   string `json:"username"`
			AvatarPath string `json:"avatar_path"`
			Rating     int32  `json:"rating"`
		} `json:"author_details"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
		ID        string `json:"id"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type TvSeriesScreenedTheatricallyResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID            string `json:"id"`
		EpisodeNumber int32  `json:"episode_number"`
		SeasonNumber  int32  `json:"season_number"`
	} `json:"results"`
}

type TvSeriesSimilarResponse struct {
	Page         int32 `json:"page"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
	Results      []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIds         []int32  `json:"genre_ids"`
		ID               int32    `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		VoteAverage      float32  `json:"vote_average"`
		VoteCount        int32    `json:"vote_count"`
	} `json:"results"`
}

type TvSeriesTranslationsResponse struct {
	ID            int32 `json:"id"`
	Transalations []struct {
		Iso_3166_1  string `json:"iso_3166_1"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Name     string `json:"name"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
			Tagline  string `json:"tagline"`
		} `json:"data"`
	} `json:"translations"`
}

type TvSeriesVideosResponse struct {
	ID      int32 `json:"id"`
	Results []struct {
		ID          string `json:"id"`
		Iso_639_1   string `json:"iso_639_1"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Site        string `json:"site"`
		Size        int32  `json:"size"`
		Type        string `json:"type"`
		Official    bool   `json:"official"`
		PublishedAt string `json:"published_at"`
	} `json:"results"`
}

// type TvSeriesWatchProvidersResponse struct {}

func (tc *TvSeriesClient) GetDetails(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesDetailsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetAccountStates(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesAccountStatesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/account_states", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesAccountStatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetAggregateCredits(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesAggregateCreditsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/aggregate_credits", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesAggregateCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetAlternativeTitles(ctx context.Context, seriesID int32) (*TvSeriesAlternativeTitlesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/alternative_titles", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesAlternativeTitlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetChanges(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesChangesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/changes", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesChangesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetContentRatings(ctx context.Context, seriesID int32) (*TvSeriesContentRatingsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/content_ratings", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesContentRatingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetCredits(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesCreditsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/credits", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetEpisodeGroups(ctx context.Context, seriesID int32) (*TvSeriesEpisodeGroupsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/episode_groups", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesEpisodeGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetExternalIds(ctx context.Context, seriesID int32) (*TvSeriesExternalIdsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/external_ids", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesExternalIdsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetImages(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesImagesResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/images", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetKeywords(ctx context.Context, seriesID int32) (*TvSeriesKeywordsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/keywords", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesKeywordsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetLatest(ctx context.Context) (*TvSeriesLatestResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, "/tv/latest")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesLatestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetRecommendations(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesRecommendationsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/recommendations", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesRecommendationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetReviews(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesReviewsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/reviews", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesReviewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetScreenedTheatrically(ctx context.Context, seriesID int32) (*TvSeriesScreenedTheatricallyResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/screened_theatrically", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesScreenedTheatricallyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetSimilar(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesSimilarResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/similar", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesSimilarResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetTranslations(ctx context.Context, seriesID int32) (*TvSeriesTranslationsResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/translations", seriesID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesTranslationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (tc *TvSeriesClient) GetVideos(ctx context.Context, seriesID int32, queryParams ...queryParam) (*TvSeriesVideosResponse, error) {
	resp, err := tc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/tv/%d/videos", seriesID), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TvSeriesVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
