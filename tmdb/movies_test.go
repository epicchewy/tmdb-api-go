package tmdb

import (
	"context"
	"net/http"
	"testing"
)

const (
	TokyoStoryName = "Tokyo Story"
)

func TestMoviesClient(t *testing.T) {
	t.Run("Get Details Tokyo Story", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movie_details.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetDetails(context.TODO(), 18148)
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if result.Title != TokyoStoryName {
			t.Errorf("expected movie title Tokyo Story, got %s", result.Title)
		}

		if result.OriginalTitle != "東京物語" {
			t.Errorf("expected movie original title 東京物語, got %s", result.OriginalTitle)
		}

		if result.ProductionCompanies[0].ID != 192 {
			t.Errorf("expected production company ID 192, got %d", result.ProductionCompanies[0].ID)
		}

		if result.BelongsToCollection != "" {
			t.Errorf("expected empty collection, got %s", result.BelongsToCollection)
		}
	})

	t.Run("Get Alternative Titles Tokyo Story", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_alternative_titles.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetAlternativeTitles(context.Background(), 18148)
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if result.Titles[0].Title != TokyoStoryName {
			t.Errorf("expected movie title Tokyo Story, got %s", result.Titles[0].Title)
		}

		if result.Titles[0].Iso_3166_1 != "US" {
			t.Errorf("expected movie title Iso_3166_1 US, got %s", result.Titles[0].Iso_3166_1)
		}

		if len(result.Titles) != 12 {
			t.Errorf("expected 12 titles, got %d", len(result.Titles))
		}
	})

	t.Run("Get External Ids Tokyo Story", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_external_ids.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetExternalIDs(context.Background(), 18148)
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if result.ImdbID != "tt0046438" {
			t.Errorf("expected movie IMDB ID tt0046438, got %s", result.ImdbID)
		}
	})

	t.Run("Get Keywords Tokyo Story", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_keywords.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetKeywords(context.Background(), 18148)
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if len(result.Keywords) != 24 {
			t.Errorf("expected 24 keywords, got %d", len(result.Keywords))
		}
	})

	t.Run("Get Movies Latest", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_latest.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetLatest(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 1225590 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}
	})

	t.Run("Get Movies Lists Tokyo Story Page 3", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_lists.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetLists(context.Background(), 18148, SingleQueryParam{"page", "3"}, SingleQueryParam{"language", "en-US"})
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if result.Page != 3 {
			t.Errorf("expected page 3, got %d", result.Page)
		}

		if len(result.Results) != 20 {
			t.Errorf("expected 20 results, got %d", len(result.Results))
		}

		if result.TotalPages != 7 {
			t.Errorf("expected 7 total pages, got %d", result.TotalPages)
		}

		if result.TotalResults != 137 {
			t.Errorf("expected 137 total results, got %d", result.TotalResults)
		}
	})

	t.Run("Get Movies Tokyo Story Recommendations", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_recommendations.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetRecommendations(context.Background(), 18148, SingleQueryParam{"page", "0"}, SingleQueryParam{"language", "en-US"})
		if err != nil {
			t.Fatal(err)
		}

		if result.Page != 1 {
			t.Errorf("expected page 1, got %d", result.Page)
		}

		if result.TotalPages != 2 {
			t.Errorf("expected 2 total pages, got %d", result.TotalPages)
		}

		if result.TotalResults != 40 {
			t.Errorf("expected 40 total results, got %d", result.TotalResults)
		}
	})

	t.Run("Get Movies Tokyo Story Reviews", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_reviews.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetReviews(context.Background(), 18148, SingleQueryParam{"page", "0"}, SingleQueryParam{"language", "en-US"})
		if err != nil {
			t.Fatal(err)
		}

		if result.ID != 18148 {
			t.Errorf("expected movie ID 18148, got %d", result.ID)
		}

		if result.Page != 1 {
			t.Errorf("expected page 1, got %d", result.Page)
		}

		if len(result.Results) != 2 {
			t.Errorf("expected 2 results, got %d", len(result.Results))
		}

		if result.TotalPages != 1 {
			t.Errorf("expected 1 total pages, got %d", result.TotalPages)
		}

		if result.TotalResults != 2 {
			t.Errorf("expected 2 total results, got %d", result.TotalResults)
		}
	})

	t.Run("Get Movies Tokyo Story Similar", func(t *testing.T) {
		testClient, testServer, err := newTestClientAndServerWithFile(http.StatusOK, "testdata/get_movies_similar.json")
		defer testServer.Close()

		if err != nil {
			t.Fatal(err)
		}

		result, err := testClient.Movies.GetSimilar(context.Background(), 18148, SingleQueryParam{"page", "0"}, SingleQueryParam{"language", "en-US"})
		if err != nil {
			t.Fatal(err)
		}

		if result.Page != 1 {
			t.Errorf("expected page 1, got %d", result.Page)
		}

		if result.TotalPages != 10647 {
			t.Errorf("expected 10647 total pages, got %d", result.TotalPages)
		}

		if result.TotalResults != 212929 {
			t.Errorf("expected 212929 total results, got %d", result.TotalResults)
		}
	})
}
