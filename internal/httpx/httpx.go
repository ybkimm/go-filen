package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func decodeResponse(resp *http.Response, v any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func DoJsonRequest(httpClient *http.Client, method, endpoint string, body io.Reader, v any) (status int, err error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return 0, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode > 299 {
		var errorResponse ErrorResponse

		err = decodeResponse(resp, &errorResponse)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("failed to decode error response: %w", err)
		}

		return resp.StatusCode, fmt.Errorf("HTTP GET failed: %s", errorResponse.Message)
	}

	err = decodeResponse(resp, v)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
