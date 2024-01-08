package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountService interface {
	GetDetails(ctx context.Context, accountId string, queryParams ...queryParam) (*AccountDetails, error)
	GetFavoriteMovies(ctx context.Context, accountId string, queryParams ...queryParam) (*FavoriteMoviesList, error)
	GetFavoriteTVShows(ctx context.Context, accountId string, queryParams ...queryParam) (*FavoriteTVShowsList, error)
	GetLists(ctx context.Context, accountId string, queryParams ...queryParam) (*AccountLists, error)
	GetRatedMovies(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedMovieList, error)
	GetRatedTVShows(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedTVShowsList, error)
	GetRatedTVEpisodes(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedTVShowEpisodesList, error)
	GetMovieWatchlist(ctx context.Context, accountId string, queryParams ...queryParam) (*MovieWatchlist, error)
	GetTVShowWatchlist(ctx context.Context, accountId string, queryParams ...queryParam) (*TVShowWatchlist, error)
}

type AccountClient struct {
	baseClient *Client
}

type AccountDetails struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
		Tmdb struct {
			AvatarPath string `json:"avatar_path"`
		} `json:"tmdb"`
	} `json:"avatar"`
	ID           int32  `json:"id"`
	Iso_639_1    string `json:"iso_639_1"`
	Iso_3166_1   string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

type FavoriteMoviesList struct {
	Page    int32 `json:"page"`
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
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type FavoriteTVShowsList struct {
	Page    int32 `json:"page"`
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
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type AccountLists struct {
	Page    int32 `json:"page"`
	Results []struct {
		Description   string `json:"description"`
		FavoriteCount int32  `json:"favorite_count"`
		ID            string `json:"id"`
		ItemCount     int32  `json:"item_count"`
		Iso_639_1     string `json:"iso_639_1"`
		ListType      string `json:"list_type"`
		Name          string `json:"name"`
		PosterPath    string `json:"poster_path"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type RatedMovieList struct {
	Page    int32 `json:"page"`
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
		Rating           float64 `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type RatedTVShowsList struct {
	Page    int32 `json:"page"`
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
		Rating           float64  `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type RatedTVShowEpisodesList struct {
	Page    int32 `json:"page"`
	Results []struct {
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
		Rating         float64 `json:"rating"`
	} `json:"results"`
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type MovieWatchlist struct {
	Page    int32 `json:"page"`
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
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

type TVShowWatchlist struct {
	Page    int32 `json:"page"`
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
	TotalPages   int32 `json:"total_pages"`
	TotalResults int32 `json:"total_results"`
}

func (ac *AccountClient) GetDetails(ctx context.Context, accountId string, queryParams ...queryParam) (*AccountDetails, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AccountDetails
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetFavoriteMovies(ctx context.Context, accountId string, queryParams ...queryParam) (*FavoriteMoviesList, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/favorite/movies", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FavoriteMoviesList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetFavoriteTVShows(ctx context.Context, accountId string, queryParams ...queryParam) (*FavoriteTVShowsList, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/favorite/tv", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FavoriteTVShowsList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetLists(ctx context.Context, accountId string, queryParams ...queryParam) (*AccountLists, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/lists", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AccountLists
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetRatedMovies(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedMovieList, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/rated/movies", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedMovieList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetRatedTVShows(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedTVShowsList, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/rated/tv", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedTVShowsList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetRatedTVEpisodes(ctx context.Context, accountId string, queryParams ...queryParam) (*RatedTVShowEpisodesList, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/rated/tv/episodes", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RatedTVShowEpisodesList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetMovieWatchlist(ctx context.Context, accountId string, queryParams ...queryParam) (*MovieWatchlist, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/watchlist/movies", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result MovieWatchlist
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AccountClient) GetTVShowWatchlist(ctx context.Context, accountId string, queryParams ...queryParam) (*TVShowWatchlist, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/account/%s/watchlist/tv", accountId), queryParams...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TVShowWatchlist
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
