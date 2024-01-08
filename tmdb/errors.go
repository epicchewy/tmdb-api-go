package tmdb

import (
	"errors"
	"fmt"
)

var (
	ErrBearerTokenMissing = errors.New("bearer token missing")
	ErrApiKeyMissing      = errors.New("api key missing")
)

type TmdbError struct {
	StatusMessage string `json:"status_message"`
	StatusCode    int    `json:"status_code"`
	Success       bool   `json:"success"`
}

func (e *TmdbError) Error() string {
	return fmt.Sprintf("Code: %d, Status Message: %s", e.StatusCode, e.StatusMessage)
}

var (
	ErrInvalidQueryParams = errors.New("invalid query params")
)
