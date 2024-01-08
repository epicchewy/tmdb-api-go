package main

import (
	"context"
	"fmt"
	"os"

	"github.com/epicchewy/tmdb-api-go/tmdb"
)

const (
	apiKey = "insert-your-api-token-here"
)

func main() {
	// instantiate client
	client, err := tmdb.NewClientWithApiKey(apiKey)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(0)
	}

	// get movie certifications
	certifications, err := client.Certifications.GetTVCertifications(context.Background())
	if err != nil {
		if tmdb.IsTmdbError(err) {
			fmt.Printf("tmdb error: %+v\n", err)
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		os.Exit(0)
	}

	// print out the certifications
	fmt.Printf("%+v\n", certifications)

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
		if tmdb.IsTmdbError(err) {
			fmt.Printf("tmdb error: %+v\n", err)
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		os.Exit(0)
	}

	// print out the tv shows
	fmt.Printf("%+v\n", tvShows)
}
