package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type ReviewsService interface {
	GetDetails(ctx context.Context, reviewID string) (*Review, error)
}

type ReviewsClient struct {
	baseClient *Client
}

type Review struct {
	ID            string `json:"id"`
	Author        string `json:"author"`
	AuthorDetails struct {
		Name       string `json:"name"`
		Username   string `json:"username"`
		AvatarPath string `json:"avatar_path"`
		Rating     int    `json:"rating"`
	} `json:"author_details"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Iso_639_1  string `json:"iso_639_1"`
	MediaID    int    `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	UpdatedAt  string `json:"updated_at"`
	URL        string `json:"url"`
}

func (rc *ReviewsClient) GetDetails(ctx context.Context, reviewID string) (*Review, error) {
	resp, err := rc.baseClient.request(ctx, http.MethodGet, "/review/"+reviewID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Review
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
