package tmdb

import (
	"context"
	"net/http"
	"testing"
)

func TestFindClient(t *testing.T) {
	// test find service interface
	t.Run("Find", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_find_by_id.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Find.FindByID(context.Background(), "tt0076759", SingleQueryParam{"external_source", "imdb_id"})
		if err != nil {
			t.Fatal(err)
		}

		if len(result.MovieResults) != 1 {
			t.Errorf("expected 1 movie result, got %d", len(result.MovieResults))
		}

		if result.MovieResults[0].ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.MovieResults[0].ID)
		}

		if result.MovieResults[0].Title != "Tokyo Story" {
			t.Errorf("expected movie title Tokyo Story, got %s", result.MovieResults[0].Title)
		}
	})
}
