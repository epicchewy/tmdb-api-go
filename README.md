# tmdb-api-go

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A Go wrapper for the [TMDb API](https://developer.themoviedb.org/reference/intro/getting-started).

## Installation

To install the library: 

```
go get github.com/epicchewy/tmdb-api-go
```

## Usage

### Initializing a client

To initialize an instance of a client:

```
import "github.com/epicchewy/tmdb-api-go/tmdb

// create client with api key auth
client, err := tmdb.NewClientWithApiKey("insert-api-key-here")
if err != nil {
  panic(err)
}

// create client with bearer token auth
client, err := tmdb.NewClientWithBearerAuth("insert-bearer-token-here")
if err != nil {
  panic(err)
}
```

### Using the client

Simply call the methods you want on the client to call the TMDb API. Here are a few examples below:

```
// get details for the movie Tokyo Story
details, err := client.Movies.GetDetails(context.Background(), 18148)
if err != nil {
  // handle error here
}

// user discover api to find tv shows
tvShows, err := client.Discover.GetTVShows(
  context.Background(),
  tmdb.SingleQueryParam{Key: "with_genres", Value: "18"},
  tmdb.SingleQueryParam{Key: "sort_by", Value: "popularity.desc"},
  tmdb.SingleQueryParam{Key: "page", Value: 1},
  tmdb.SingleQueryParam{Key: "language", Value: "en-US"},
  tmdb.SingleQueryParam{Key: "with_status", Value: "2|3"}, // example of piping multiple values into query param
)
if err != nil {
  // handle tmdb error
  if tmdb.IsTmdbError(err) {
    fmt.Printf("tmdb error: %+v\n", err)
  }
  // more error handling
}
```

## Examples

Examples of API usage can be found in the `./examples` directory.

You can also play with the API directly in the [TMDb API docs](https://developer.themoviedb.org/reference/intro/getting-started). TMDb hosts their docs on [ReadMe](https://docs.readme.com/), a technical documentation solution that allows developers to test API endpoints directly on their platform with their authentication credentials.

## TODOs

- Rate Limiting
- Non GET request API endpoints
- Unit Tests (currently 11% coverage)
- Integration Tests