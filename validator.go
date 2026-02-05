package marshal

import (
	"context"
	"net/http"
)

// ValidateJSON validates a JSON string against an OpenAPI specification
func (c *Client) ValidateJSON(ctx context.Context, openAPISpec, jsonString, operationID string) (*ValidateResponse, error) {
	req := ValidateRequest{
		OpenAPISpec: openAPISpec,
		JSONString:  jsonString,
		OperationID: operationID,
	}
	var result ValidateResponse
	err := c.doRequest(ctx, http.MethodPost, "/validator/validate", req, &result, true)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
