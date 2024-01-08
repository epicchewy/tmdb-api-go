package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CollectionsService interface {
	GetDetails(ctx context.Context, collectionId int32, queryParams ...queryParam) (*Collection, error)
	GetImages(ctx context.Context, collectionId int32, queryParams ...queryParam) (*CollectionImages, error)
	GetTranslations(ctx context.Context, collectionId int32) (*CollectionTranslations, error)
}

type CollectionsClient struct {
	baseClient *Client
}

type Collection struct {
	BackdropPath string `json:"backdrop_path"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	Overview     string `json:"overview"`
	Parts        []Part `json:"parts"`
}

type Part struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       string  `json:"poster_path"`
	Popularity       float64 `json:"popularity"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type CollectionImages struct {
	ID        int `json:"id"`
	Backdrops []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"backdrops"`
	Posters []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso_639_1   string  `json:"iso_639_1"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
}

type CollectionTranslations struct {
	ID           int `json:"id"`
	Translations []struct {
		Iso_3166_1  string `json:"iso_3166_1"`
		Iso_639_1   string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title        string `json:"title"`
			Overview     string `json:"overview"`
			Homepage     string `json:"homepage"`
			Tagline      string `json:"tagline"`
			PosterPath   string `json:"poster_path"`
			BackdropPath string `json:"backdrop_path"`
		} `json:"data"`
	} `json:"translations"`
}

func (cc *CollectionsClient) GetDetails(ctx context.Context, collectionId int32, queryParams ...queryParam) (*Collection, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/collection/%d", collectionId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Collection
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *CollectionsClient) GetImages(ctx context.Context, collectionId int32, queryParams ...queryParam) (*CollectionImages, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/collection/%d/images", collectionId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CollectionImages
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *CollectionsClient) GetTranslations(ctx context.Context, collectionId int32) (*CollectionTranslations, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/collection/%d/translations", collectionId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CollectionTranslations
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
