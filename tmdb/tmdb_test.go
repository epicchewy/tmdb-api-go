package tmdb

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

func newTestClientAndServerWithFile(statusCode int, filePath string) (*Client, *httptest.Server, error) {
	// get file contents
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		_, _ = io.Copy(w, f)
		r.Body.Close()
	}))
	baseUrl, _ := url.Parse(server.URL)
	client, err := NewClientWithBearerAuth("test", WithBaseUrl(baseUrl))
	if err != nil {
		server.Close()
		return nil, nil, err
	}
	return client, server, nil
}
