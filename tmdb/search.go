package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type SearchService interface {
	GetCollection(ctx context.Context, queryParams ...queryParam) (*SearchCollectionResponse, error)
	GetCompany(ctx context.Context, queryParams ...queryParam) (*SearchCompanyResponse, error)
	GetKeyword(ctx context.Context, queryParams ...queryParam) (*SearchKeywordResponse, error)
	GetMovie(ctx context.Context, queryParams ...queryParam) (*SearchMovieResponse, error)
	GetMulti(ctx context.Context, queryParams ...queryParam) (*SearchMultiResponse, error)
	GetPerson(ctx context.Context, queryParams ...queryParam) (*SearchPersonResponse, error)
	GetTv(ctx context.Context, queryParams ...queryParam) (*SearchTvResponse, error)
}

type SearchClient struct {
	baseClient *Client
}

type SearchCollectionResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool   `json:"adult"`
		BackdropPath     string `json:"backdrop_path"`
		ID               int    `json:"id"`
		Name             string `json:"name"`
		OriginalLanguage string `json:"original_language"`
		OriginalName     string `json:"original_name"`
		Overview         string `json:"overview"`
		PosterPath       string `json:"poster_path"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchCompanyResponse struct {
	Page    int `json:"page"`
	Results []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchKeywordResponse struct {
	Page    int `json:"page"`
	Results []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchMovieResponse struct {
	Page    int `json:"page"`
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
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchMultiResponse struct {
	Page    int `json:"page"`
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
		MediaType        string  `json:"media_type"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchPersonResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult              bool    `json:"adult"`
		Gender             int32   `json:"gender"`
		ID                 int32   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int32 `json:"genre_ids"`
			ID               int32   `json:"id"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int32   `json:"vote_count"`
		} `json:"known_for"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchTvResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIds         []int32  `json:"genre_ids"`
		ID               int32    `json:"id"`
		Name             string   `json:"name"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int32    `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func (sc *SearchClient) GetCollection(ctx context.Context, queryParams ...queryParam) (*SearchCollectionResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/collection", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetCompany(ctx context.Context, queryParams ...queryParam) (*SearchCompanyResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/company", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchCompanyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetKeyword(ctx context.Context, queryParams ...queryParam) (*SearchKeywordResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/keyword", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchKeywordResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetMovie(ctx context.Context, queryParams ...queryParam) (*SearchMovieResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/movie", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchMovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetMulti(ctx context.Context, queryParams ...queryParam) (*SearchMultiResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/multi", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchMultiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetPerson(ctx context.Context, queryParams ...queryParam) (*SearchPersonResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/person", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchPersonResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (sc *SearchClient) GetTv(ctx context.Context, queryParams ...queryParam) (*SearchTvResponse, error) {
	resp, err := sc.baseClient.request(ctx, http.MethodGet, "/search/tv", queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchTvResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
