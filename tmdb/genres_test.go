package tmdb

import (
	"context"
	"net/http"
	"testing"
)

func TestGenresClient(t *testing.T) {
	// test genres service interface

	t.Run("GetMovieList", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_genres_movie_list.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Genres.GetMovieGenres(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(result.Genres) != 19 {
			t.Errorf("expected 19 genres, got %d", len(result.Genres))
		}
	})

	t.Run("GetTVList", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_genres_tv_list.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Genres.GetTVGenres(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(result.Genres) != 16 {
			t.Errorf("expected 16 genres, got %d", len(result.Genres))
		}
	})
}
