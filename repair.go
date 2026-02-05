package marshal

import (
	"context"
	"net/http"
)

// FixJSON attempts to repair and fix malformed or invalid JSON strings
func (c *Client) FixJSON(ctx context.Context, jsonToFix string) (*FixJSONResponse, error) {
	req := FixJSONRequest{
		JSONToFix: jsonToFix,
	}
	var result FixJSONResponse
	err := c.doRequest(ctx, http.MethodPost, "/repair/fix-json", req, &result, true)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
