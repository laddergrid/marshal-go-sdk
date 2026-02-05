# Marshal Go SDK

Official Go SDK for the [Marshal Engine API](https://api.laddergrid.com) - JSON validation and repair services.

[![Go Reference](https://pkg.go.dev/badge/github.com/laddergrid/marshal-go-sdk.svg)](https://pkg.go.dev/github.com/laddergrid/marshal-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/laddergrid/marshal-go-sdk)](https://goreportcard.com/report/github.com/laddergrid/marshal-go-sdk)

## Installation

```bash
go get github.com/laddergrid/marshal-go-sdk
```

## Features

- **Health Check**: Verify API availability
- **JSON Validation**: Validate JSON against OpenAPI specifications
- **JSON Repair**: Fix malformed or invalid JSON strings
- **Context Support**: All methods support context for timeout and cancellation
- **Type-Safe**: Strongly typed Go structs for all API requests and responses
- **Zero Dependencies**: Uses only Go standard library

## Quick Start

### Prerequisites

You need an API key to use the Marshal Engine API. Your API key should be in the format `lg-m-...`

### Health Check

```go
package main

import (
    "context"
    "fmt"
    "log"

    marshal "github.com/laddergrid/marshal-go-sdk"
)

func main() {
    client := marshal.NewClient("lg-m-your-api-key")

    resp, err := client.Ping(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp.Message) // "pong"
}
```

### JSON Validation

Validate JSON data against an OpenAPI specification:

```go
package main

import (
    "context"
    "fmt"
    "log"

    marshal "github.com/laddergrid/marshal-go-sdk"
)

func main() {
    client := marshal.NewClient("lg-m-your-api-key")

    openAPISpec := `{
  "type": "object",
  "properties": {
    "name": {"type": "string"},
    "age": {"type": "number"}
  },
  "required": ["name"]
}`

    jsonString := `{"name": "John Doe", "age": 30}`

    result, err := client.ValidateJSON(
        context.Background(),
        openAPISpec,
        jsonString,
        "createUser",
    )
    if err != nil {
        log.Fatal(err)
    }

    if result.IsValid {
        fmt.Println("JSON is valid!")
    } else {
        fmt.Println("Validation errors:")
        for _, err := range result.Errors {
            fmt.Printf("  - %s: %s\n", err.Field, err.Error)
        }
    }
}
```

### JSON Repair

Automatically fix malformed JSON:

```go
package main

import (
    "context"
    "fmt"
    "log"

    marshal "github.com/laddergrid/marshal-go-sdk"
)

func main() {
    client := marshal.NewClient("lg-m-your-api-key")

    malformedJSON := `{"name": "John", age: 30, "city": "New York"`

    result, err := client.FixJSON(context.Background(), malformedJSON)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Fixed JSON:", result.FixedJSON)
    // Output: {"name": "John", "age": 30, "city": "New York"}
}
```

## Client Configuration

### Custom Base URL

For development or self-hosted instances:

```go
client := marshal.NewClient(
    "lg-m-your-api-key",
    marshal.WithBaseURL("http://localhost:8080"),
)
```

### Custom HTTP Client

Configure timeouts, proxies, or other HTTP settings:

```go
import (
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        IdleConnTimeout:     90 * time.Second,
        DisableCompression:  false,
    },
}

client := marshal.NewClient(
    "lg-m-your-api-key",
    marshal.WithHTTPClient(httpClient),
)
```

### Multiple Options

```go
client := marshal.NewClient(
    "lg-m-your-api-key",
    marshal.WithBaseURL("https://api.laddergrid.com"),
    marshal.WithHTTPClient(customHTTPClient),
)
```

## Error Handling

All methods return errors that can be type-asserted to `*marshal.APIError` for detailed information:

```go
result, err := client.ValidateJSON(ctx, spec, json, opID)
if err != nil {
    if apiErr, ok := err.(*marshal.APIError); ok {
        fmt.Printf("Status: %d\n", apiErr.StatusCode)
        fmt.Printf("Message: %s\n", apiErr.Message)

        // Access the raw HTTP response if needed
        if apiErr.Response != nil {
            fmt.Printf("Headers: %v\n", apiErr.Response.Header)
        }
    } else {
        // Network error or other non-API error
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

### Common Error Status Codes

- `400` - Bad request (invalid input)
- `401` - Unauthorized (invalid or missing API key)
- `404` - Not found
- `500` - Internal server error

## Context Support

All API methods accept a `context.Context` parameter for timeout and cancellation:

```go
import (
    "context"
    "time"
)

// Timeout after 10 seconds
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

result, err := client.Ping(ctx)
if err != nil {
    // Handle timeout or other errors
    log.Fatal(err)
}

// Cancel on demand
ctx, cancel := context.WithCancel(context.Background())
go func() {
    // Cancel after some condition
    time.Sleep(5 * time.Second)
    cancel()
}()

result, err := client.FixJSON(ctx, malformedJSON)
```

## Environment Variables

For security, consider storing your API key in an environment variable:

```go
import "os"

apiKey := os.Getenv("MARSHAL_API_KEY")
if apiKey == "" {
    log.Fatal("MARSHAL_API_KEY environment variable is required")
}

client := marshal.NewClient(apiKey)
```

Then run your application:

```bash
export MARSHAL_API_KEY="lg-m-your-api-key"
go run main.go
```

## Examples

See the [examples](./examples) directory for complete working examples:

- [Quick Start](./examples/quickstart/main.go) - Simple health check
- [Health Check](./examples/health/main.go) - Verify API status
- [JSON Validation](./examples/validation/main.go) - Validate JSON against schemas
- [JSON Repair](./examples/repair/main.go) - Fix malformed JSON

Run an example:

```bash
export MARSHAL_API_KEY="lg-m-your-api-key"
cd examples/validation
go run main.go
```

## API Reference

### Client Creation

```go
func NewClient(apiKey string, opts ...Option) *Client
```

Creates a new Marshal API client with the provided API key. The API key is required.

**Parameters:**
- `apiKey` (string) - Your Marshal API key in the format "lg-m-..."
- `opts` (...Option) - Optional configuration options

**Returns:**
- `*Client` - A configured Marshal client instance

### Options

- `WithBaseURL(baseURL string)` - Set custom base URL
- `WithHTTPClient(httpClient *http.Client)` - Set custom HTTP client

### Health Check

```go
func (c *Client) Ping(ctx context.Context) (*PingResponse, error)
```

Performs a health check on the API to verify it's available and responding.

**Parameters:**
- `ctx` (context.Context) - Context for request cancellation and timeout

**Returns:**
- `*PingResponse` - Response containing a status message
- `error` - Error if the request fails

### JSON Validation

```go
func (c *Client) ValidateJSON(ctx context.Context, openAPISpec, jsonString, operationID string) (*ValidateResponse, error)
```

Validates a JSON string against an OpenAPI specification schema.

**Parameters:**
- `ctx` (context.Context) - Context for request cancellation and timeout
- `openAPISpec` (string) - The OpenAPI specification as a JSON string
- `jsonString` (string) - The JSON data to validate
- `operationID` (string) - The operationId from the OpenAPI spec

**Returns:**
- `*ValidateResponse` - Validation result with IsValid flag and any errors
- `error` - Error if the request fails

### JSON Repair

```go
func (c *Client) FixJSON(ctx context.Context, jsonToFix string) (*FixJSONResponse, error)
```

Attempts to repair and fix malformed or invalid JSON strings.

**Parameters:**
- `ctx` (context.Context) - Context for request cancellation and timeout
- `jsonToFix` (string) - The malformed JSON string to repair

**Returns:**
- `*FixJSONResponse` - Response containing the repaired JSON string
- `error` - Error if the request fails

## Best Practices

### 1. Use Context with Timeouts

Always use contexts with appropriate timeouts to prevent hanging requests:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

result, err := client.ValidateJSON(ctx, spec, json, opID)
```

### 2. Handle Errors Appropriately

Check for specific error types and handle them accordingly:

```go
result, err := client.FixJSON(ctx, jsonStr)
if err != nil {
    if apiErr, ok := err.(*marshal.APIError); ok {
        if apiErr.StatusCode == 401 {
            log.Fatal("Invalid API key")
        }
    }
    log.Fatal(err)
}
```

### 3. Reuse Client Instances

Create one client instance and reuse it for multiple requests:

```go
// Good - reuse client
client := marshal.NewClient(apiKey)
for _, item := range items {
    result, err := client.ValidateJSON(ctx, spec, item.JSON, opID)
    // ...
}

// Avoid - creating new clients for each request
for _, item := range items {
    client := marshal.NewClient(apiKey) // wasteful
    result, err := client.ValidateJSON(ctx, spec, item.JSON, opID)
    // ...
}
```

### 4. Secure Your API Key

- Never commit API keys to version control
- Use environment variables or secure secret management
- Rotate keys periodically
- Use different keys for development and production

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Support

- GitHub: [https://github.com/laddergrid/marshal-engine](https://github.com/laddergrid/marshal-engine)
- Issues: [Report bugs and feature requests](https://github.com/laddergrid/marshal-go-sdk/issues)

## Changelog

### v1.0.0 (Initial Release)

- Health check endpoint
- JSON validation against OpenAPI specifications
- JSON repair functionality
- Full context support
- Comprehensive error handling
- API key authentication
