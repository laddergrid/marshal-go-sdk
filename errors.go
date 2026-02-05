package marshal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIError represents an error returned by the Marshal API
type APIError struct {
	StatusCode int
	Message    string
	Response   *http.Response
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("marshal api error (status %d): %s", e.StatusCode, e.Message)
}

// parseErrorResponse attempts to parse an error response from the API
func parseErrorResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("failed to read error response: %v", err),
			Response:   resp,
		}
	}

	var errResp ErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
			Response:   resp,
		}
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    errResp.Error,
		Response:   resp,
	}
}
