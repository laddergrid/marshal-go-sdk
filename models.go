package marshal

// PingResponse represents the response from the health check endpoint
type PingResponse struct {
	Message string `json:"message"`
}

// ValidateRequest represents the request to validate JSON against an OpenAPI specification
type ValidateRequest struct {
	OpenAPISpec string `json:"open_api_spec"`
	JSONString  string `json:"json_string"`
	OperationID string `json:"operation_id"`
}

// ValidateResponse represents the response from JSON validation
type ValidateResponse struct {
	IsValid bool              `json:"is_valid"`
	Errors  []ValidationError `json:"errors"`
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// FixJSONRequest represents the request to fix malformed JSON
type FixJSONRequest struct {
	JSONToFix string `json:"json_to_fix"`
}

// FixJSONResponse represents the response from JSON repair
type FixJSONResponse struct {
	FixedJSON string `json:"fixed_json"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error string `json:"error"`
}
