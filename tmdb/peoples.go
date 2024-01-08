package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PeoplesService interface {
	GetDetails(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleResponse, error)
	GetChanges(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleChangesResponse, error)
	GetCombinedCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleCombinedCreditsResponse, error)
	GetExternalIds(ctx context.Context, personId int32) (*PeopleExternalIdsResponse, error)
	GetImages(ctx context.Context, personId int32) (*PeopleImagesResponse, error)
	GetLatest(ctx context.Context) (*PeopleLatestResponse, error)
	GetMovieCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleMovieCreditsResponse, error)
	GetTVCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleTVCreditsResponse, error)
	GetTranslations(ctx context.Context, personId int32) (*PeopleTranslationsResponse, error)
}

type PeopleClient struct {
	baseClient *Client
}

type PeopleResponse struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           string   `json:"birthday"`
	Deathday           string   `json:"deathday"`
	Gender             int32    `json:"gender"`
	Homepage           string   `json:"homepage"`
	ID                 int32    `json:"id"`
	ImdbID             string   `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	Popularity         float32  `json:"popularity"`
	ProfilePath        string   `json:"profile_path"`
}

type PeopleChangesResponse struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID         string `json:"id"`
			Action     string `json:"action"`
			Time       string `json:"time"`
			Iso_639_1  string `json:"iso_639_1"`
			Iso_3166_1 string `json:"iso_3166_1"`
			Value      string `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

type Cast struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Character        string  `json:"character"`
	CreditID         string  `json:"credit_id"`
	GenreIDs         []int32 `json:"genre_ids"`
	ID               int32   `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Order            int32   `json:"order"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int32   `json:"vote_count"`
	MediaType        string  `json:"media_type"`
}

type Crew struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	CreditID         string  `json:"credit_id"`
	Department       string  `json:"department"`
	GenreIDs         []int32 `json:"genre_ids"`
	ID               int32   `json:"id"`
	Job              string  `json:"job"`
	MediaType        string  `json:"media_type"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int32   `json:"vote_count"`
}

type PeopleCombinedCreditsResponse struct {
	ID   int32  `json:"id"`
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type PeopleExternalIdsResponse struct {
	ID          int32  `json:"id"`
	FreebaseID  string `json:"freebase_id"`
	FreebaseMID string `json:"freebase_mid"`
	IMDBID      string `json:"imdb_id"`
	TvrageID    string `json:"tvrage_id"`
	WikidataId  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TiktokID    string `json:"tiktok_id"`
	TwitterID   string `json:"twitter_id"`
	YoutubeID   string `json:"youtube_id"`
}

// TODO: figure this out
type PeopleImagesResponse struct{}

type PeopleLatestResponse struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           string   `json:"birthday"`
	Deathday           string   `json:"deathday"`
	Gender             string   `json:"gender"`
	Homepage           string   `json:"homepage"`
	ID                 int32    `json:"id"`
	ImdbID             string   `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	Popularity         float32  `json:"popularity"`
	ProfilePath        string   `json:"profile_path"`
}

type PeopleMovieCreditsResponse struct {
	ID   int32 `json:"id"`
	Cast Cast  `json:"cast"`
	Crew Crew  `json:"crew"`
}

type PeopleTVCreditsResponse struct {
	ID   int32 `json:"id"`
	Cast []struct {
		EpisodeCount int32 `json:"episode_count"`
		Cast
	} `json:"cast"`
	Crew []struct {
		EpisodeCount int32 `json:"episode_count"`
		Crew
	} `json:"crew"`
}

type PeopleTranslationsResponse struct {
	ID           int32 `json:"id"`
	Translations []struct {
		Iso_639_1   string `json:"iso_639_1"`
		Iso_3166_1  string `json:"iso_3166_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Biography string `json:"biography"`
		} `json:"data"`
	} `json:"translations"`
}

func (c *PeopleClient) GetDetails(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d", personId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetChanges(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleChangesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/changes", personId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleChangesResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetCombinedCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleCombinedCreditsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/combined_credits", personId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleCombinedCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetExternalIds(ctx context.Context, personId int32) (*PeopleExternalIdsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/external_ids", personId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleExternalIdsResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetImages(ctx context.Context, personId int32) (*PeopleImagesResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/images", personId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetLatest(ctx context.Context) (*PeopleLatestResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, "/person/latest")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleLatestResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetMovieCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleMovieCreditsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/movie_credits", personId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleMovieCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetTVCredits(ctx context.Context, personId int32, queryParams ...queryParam) (*PeopleTVCreditsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/tv_credits", personId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleTVCreditsResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}

func (c *PeopleClient) GetTranslations(ctx context.Context, personId int32) (*PeopleTranslationsResponse, error) {
	resp, err := c.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/person/%d/translations", personId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var people PeopleTranslationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}
	return &people, nil
}
