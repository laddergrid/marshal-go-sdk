package marshal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Marshal API
	DefaultBaseURL = "https://api.laddergrid.com"
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
)

// Client is the main client for interacting with the Marshal Engine API
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

// Option is a functional option for configuring the Client
type Option func(*Client)

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient creates a new Marshal API client with the provided API key.
// The API key is required and must be in the format "lg-m-..."
func NewClient(apiKey string, opts ...Option) *Client {
	client := &Client{
		baseURL: DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		apiKey: apiKey,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// doRequest performs an HTTP request and decodes the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}, requireAuth bool) error {
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Set API key authentication header for endpoints that require it
	if requireAuth && c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseErrorResponse(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// Ping performs a health check on the API
func (c *Client) Ping(ctx context.Context) (*PingResponse, error) {
	var result PingResponse
	err := c.doRequest(ctx, http.MethodGet, "/ping", nil, &result, false)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
