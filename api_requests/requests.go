package api_requests

import (
	"fmt"
	"io"
	"net/http"
	"terminal-illness/utils"
	"time"
)

// HTTPMethod represents valid HTTP methods
type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
)

// Validate checks if the HTTP method is valid
func (m HTTPMethod) Validate() error {
	switch m {
	case GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD:
		return nil
	}
	return fmt.Errorf("invalid HTTP method: %s", m)
}

// MakeRequest sends an HTTP request to the specified endpoint
func MakeRequest(url string, method HTTPMethod, bearerToken string) (*http.Response, error) {
	// Validate the HTTP method
	if err := method.Validate(); err != nil {
		return nil, err
	}

	// Create a new HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create the request
	req, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		return nil, err
	}

	// Add bearer token if provided
	if bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+bearerToken)
	}

	// Add common headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Helper function to read response body
func ReadResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	responseURL := resp.Request.URL.String()
	if err := utils.SaveURL(responseURL); err != nil {
		// Handle error if needed
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Helper function to read response body
func ReadStatus(resp *http.Response) (string, error) {

	return resp.Status, nil
}
