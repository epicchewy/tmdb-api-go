package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

// this package is the entry point for the tmdb package
// it has the main http client

const (
	defaultApiUrl     = "https://api.themoviedb.org"
	apiVersion        = "3"
	defaultMaxRetries = 3
)

type ClientType string

const (
	ApiKey     ClientType = "ApiKey"
	BearerAuth ClientType = "BearerAuth"
)

type ClientOption func(*Client)

type Client struct {
	client     *http.Client
	baseUrl    *url.URL
	clientType ClientType
	token      string
	logger     *slog.Logger

	maxRetries int

	// Services used for talking to different parts of the tmdb API
	Accounts        AccountService
	Authentication  AuthenticationService
	Certifications  CertificationService
	Changes         ChangesService
	Collections     CollectionsService
	Companies       CompaniesService
	Configuration   ConfigurationsService
	Credits         CreditsService
	Discover        DiscoverService
	Find            FindService
	Genres          GenresService
	GuestSessions   GuestSessionsService
	Keywords        KeywordsService
	Lists           ListsService
	MovieLists      MovieListsService
	Movies          MoviesService
	Networks        NetworksService
	PeopleLists     PeopleListsService
	People          PeoplesService
	Reviews         ReviewsService
	Search          SearchService
	Trending        TrendingService
	TvEpisodeGroups TvEpisodeGroupsService
	TvEpisodes      TvEpisodesService
	TvSeasons       TvSeasonsService
	TvSeriesLists   TvSeriesListsService
	TvSeries        TvSeriesService
	WatchProviders  WatchProvidersService
}

// NewClient returns a new tmdb client that adds a bearer auth header to each request
func NewClientWithBearerAuth(bearerToken string, opts ...ClientOption) (*Client, error) {
	if bearerToken == "" {
		return nil, ErrBearerTokenMissing
	}

	baseUrl, err := url.Parse(defaultApiUrl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:     http.DefaultClient,
		baseUrl:    baseUrl,
		maxRetries: defaultMaxRetries,
		clientType: BearerAuth,
		token:      bearerToken,
		logger:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Accounts = &AccountClient{baseClient: c}
	c.Authentication = &AuthenticationClient{baseClient: c}
	c.Certifications = &CertificationClient{baseClient: c}
	c.Changes = &ChangesClient{baseClient: c}
	c.Collections = &CollectionsClient{baseClient: c}
	c.Companies = &CompaniesClient{baseClient: c}
	c.Configuration = &ConfigurationsClient{baseClient: c}
	c.Credits = &CreditsClient{baseClient: c}
	c.Discover = &DiscoverClient{baseClient: c}
	c.Find = &FindClient{baseClient: c}
	c.Genres = &GenreClient{baseClient: c}
	c.GuestSessions = &GuestSessionsClient{baseClient: c}
	c.Keywords = &KeywordsClient{baseClient: c}
	c.Lists = &ListsClient{baseClient: c}
	c.MovieLists = &MovieListsClient{baseClient: c}
	c.Movies = &MoviesClient{baseClient: c}
	c.Networks = &NetworksClient{baseClient: c}
	c.PeopleLists = &PeopleListsClient{baseClient: c}
	c.People = &PeopleClient{baseClient: c}
	c.Reviews = &ReviewsClient{baseClient: c}
	c.Search = &SearchClient{baseClient: c}
	c.Trending = &TrendingClient{baseClient: c}
	c.TvEpisodeGroups = &TvEpisodeGroupsClient{baseClient: c}
	c.TvEpisodes = &TvEpisodesClient{baseClient: c}
	c.TvSeasons = &TvSeasonsClient{baseClient: c}
	c.TvSeriesLists = &TvSeriesListsClient{baseClient: c}
	c.TvSeries = &TvSeriesClient{baseClient: c}
	c.WatchProviders = &WatchProvidersClient{baseClient: c}

	return c, nil
}

func NewClientWithApiKey(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, ErrApiKeyMissing
	}

	baseUrl, err := url.Parse(defaultApiUrl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:     http.DefaultClient,
		baseUrl:    baseUrl,
		clientType: ApiKey,
		token:      apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Accounts = &AccountClient{baseClient: c}
	c.Authentication = &AuthenticationClient{baseClient: c}
	c.Certifications = &CertificationClient{baseClient: c}
	c.Changes = &ChangesClient{baseClient: c}
	c.Collections = &CollectionsClient{baseClient: c}
	c.Companies = &CompaniesClient{baseClient: c}
	c.Configuration = &ConfigurationsClient{baseClient: c}
	c.Credits = &CreditsClient{baseClient: c}
	c.Discover = &DiscoverClient{baseClient: c}
	c.Find = &FindClient{baseClient: c}
	c.Genres = &GenreClient{baseClient: c}
	c.GuestSessions = &GuestSessionsClient{baseClient: c}
	c.Keywords = &KeywordsClient{baseClient: c}
	c.Lists = &ListsClient{baseClient: c}
	c.MovieLists = &MovieListsClient{baseClient: c}
	c.Movies = &MoviesClient{baseClient: c}
	c.Networks = &NetworksClient{baseClient: c}
	c.PeopleLists = &PeopleListsClient{baseClient: c}
	c.People = &PeopleClient{baseClient: c}
	c.Reviews = &ReviewsClient{baseClient: c}
	c.Search = &SearchClient{baseClient: c}
	c.Trending = &TrendingClient{baseClient: c}
	c.TvEpisodeGroups = &TvEpisodeGroupsClient{baseClient: c}
	c.TvEpisodes = &TvEpisodesClient{baseClient: c}
	c.TvSeasons = &TvSeasonsClient{baseClient: c}
	c.TvSeriesLists = &TvSeriesListsClient{baseClient: c}
	c.TvSeries = &TvSeriesClient{baseClient: c}
	c.WatchProviders = &WatchProvidersClient{baseClient: c}

	return c, nil
}

func WithRetries(maxRetries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

func WithHttpClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

func WithBaseUrl(baseUrl *url.URL) ClientOption {
	return func(c *Client) {
		c.baseUrl = baseUrl
	}
}

type queryParam interface {
	apply(url.Values)
	getKey() string
}

type SingleQueryParam struct {
	Key   string
	Value interface{}
}

func (p SingleQueryParam) getKey() string {
	return p.Key
}

func (p SingleQueryParam) apply(v url.Values) {
	v.Set(p.Key, fmt.Sprintf("%v", p.Value))
}

type MultiQueryParam struct {
	Key    string
	Values []interface{}
}

func (p MultiQueryParam) getKey() string {
	return p.Key
}

func (p MultiQueryParam) apply(v url.Values) {
	for _, value := range p.Values {
		v.Add(p.Key, fmt.Sprintf("%v", value))
	}
}

func (c *Client) request(ctx context.Context, method, path string, queryParams ...queryParam) (*http.Response, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("/%s/%s", apiVersion, path))
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	for _, param := range queryParams {
		param.apply(v)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.clientType == BearerAuth {
		req.Header.Set("Authorization", "Bearer "+c.token)
	} else {
		v.Set("api_key", c.token)
	}

	req.URL.RawQuery = v.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: add rate limiting

	err = c.checkResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	// decode resp.body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// decode body into TmdbError
	var tmdbError TmdbError
	err = json.Unmarshal(body, &tmdbError)
	if err != nil {
		return err
	}
	return &tmdbError
}

func IsTmdbError(err error) bool {
	_, ok := err.(*TmdbError)
	return ok
}
