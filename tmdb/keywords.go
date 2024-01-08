package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type KeywordsService interface {
	GetDetails(ctx context.Context, keywordID int) (*KeywordDetailsResponse, error)
}

type KeywordsClient struct {
	baseClient *Client
}

type KeywordDetailsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (kc *KeywordsClient) GetDetails(ctx context.Context, keywordID int) (*KeywordDetailsResponse, error) {
	resp, err := kc.baseClient.request(ctx, http.MethodGet, fmt.Sprintf("/keyword/%d", keywordID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result KeywordDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
