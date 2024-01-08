package tmdb

import (
	"context"
	"encoding/json"
	"net/http"
)

type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}

type CertificationResults struct {
	Certifications map[string][]Certification `json:"certifications"`
}

type CertificationService interface {
	GetMovieCertifications(ctx context.Context) (*CertificationResults, error)
	GetTVCertifications(ctx context.Context) (*CertificationResults, error)
}

type CertificationClient struct {
	baseClient *Client
}

func (cc *CertificationClient) GetMovieCertifications(ctx context.Context) (*CertificationResults, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/certification/movie/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CertificationResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cc *CertificationClient) GetTVCertifications(ctx context.Context) (*CertificationResults, error) {
	resp, err := cc.baseClient.request(ctx, http.MethodGet, "/certification/tv/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CertificationResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
