package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthenticationService interface {
	CreateGuestSession(ctx context.Context) (*GuestSessionResponse, error)
	CreateRequestToken(ctx context.Context) (*RequestTokenResponse, error)
	ValidateKey(ctx context.Context) (*ValidateResponse, error)
}

type AuthenticationClient struct {
	baseClient *Client
}

type GuestSessionResponse struct {
	Success   bool   `json:"success"`
	GuestID   string `json:"guest_session_id"`
	ExpiresAt string `json:"expires_at"`
}

type RequestTokenResponse struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}

type ValidateResponse struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

func (ac *AuthenticationClient) CreateGuestSession(ctx context.Context) (*GuestSessionResponse, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, "/authentication/guest_session/new")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GuestSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AuthenticationClient) CreateRequestToken(ctx context.Context) (*RequestTokenResponse, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, "/authentication/token/new")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RequestTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ac *AuthenticationClient) ValidateKey(ctx context.Context) (*ValidateResponse, error) {
	resp, err := ac.baseClient.request(ctx, http.MethodGet, "/authentication")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ValidateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
