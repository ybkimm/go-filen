package filen

import (
	"encoding/json"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"path"
	"sync"

	"github.com/ybkimm/go-filen/internal/httpx"
)

type APIClient struct {
	mu         sync.RWMutex
	httpClient *http.Client
	endpoints  []string
}

func NewAPIClient(httpClient *http.Client) *APIClient {
	return &APIClient{
		httpClient: httpClient,
		endpoints:  defaultAPIDomains,
	}
}

func (c *APIClient) SetEndpoints(endpoints []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.endpoints = endpoints
}

func (c *APIClient) doRequestWithBody(method, apiPath string, body any, v any) (status int, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Get random endpoint
	endpoint := c.endpoints[rand.Int63n(int64(len(c.endpoints)))]
	uri := endpoint + path.Join("/", apiPath)

	// Prepare reader that using json.Encoder as data source
	var r io.Reader

	if body != nil {
		pr, pw := io.Pipe()
		go func() {
			defer pw.Close()

			encErr := json.NewEncoder(pw).Encode(body)
			if encErr != nil {
				slog.Error("api: failed to encode body", "error", encErr)
			}
		}()

		r = pr
	}

	return httpx.DoJsonRequest(c.httpClient, method, uri, r, v)
}

func (c *APIClient) getJSON(endpoint string, v any) (status int, err error) {
	return c.doRequest(http.MethodGet, endpoint, nil, v)
}

func (c *APIClient) postJSON(endpoint string, v any) (status int, err error) {
	return c.doRequest(http.MethodPost, endpoint, nil, v)
}

func (c *APIClient) postJSONWithBody(endpoint string, body any, v any) (status int, err error) {
	return c.doRequest(http.MethodPost, endpoint, body, v)
}

func (c *APIClient) putJSON(endpoint string, v any) (status int, err error) {
	return c.doRequest(http.MethodPut, endpoint, nil, v)
}

func (c *APIClient) putJSONWithBody(endpoint string, body any, v any) (status int, err error) {
	return c.doRequest(http.MethodPut, endpoint, body, v)
}

func (c *APIClient) deleteJSON(endpoint string, v any) (status int, err error) {
	return c.doRequest(http.MethodDelete, endpoint, nil, v)
}

func (c *APIClient) deleteJSONWithBody(endpoint string, body any, v any) (status int, err error) {
	return c.doRequest(http.MethodDelete, endpoint, body, v)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
