# Marshal Go SDK Examples

This directory contains example applications demonstrating how to use the Marshal Go SDK.

## Prerequisites

1. Install the SDK:
```bash
go get github.com/laddergrid/marshal-go-sdk
```

2. Set your API key as an environment variable:
```bash
export MARSHAL_API_KEY="lg-m-your-api-key"
```

## Running the Examples

### Quick Start

The simplest example to verify the API is working:

```bash
cd examples/quickstart
go run main.go
```

### Health Check

Check if the Marshal API is available:

```bash
cd examples/health
go run main.go
```

### JSON Validation

Validate JSON against OpenAPI specifications:

```bash
cd examples/validation
go run main.go
```

This demonstrates:
- Valid JSON validation
- Invalid JSON detection (missing required field)
- Type validation errors

### JSON Repair

Automatically fix malformed JSON:

```bash
cd examples/repair
go run main.go
```

This shows how to repair:
- Unquoted property names
- Missing closing braces
- Trailing commas
- Single quotes instead of double quotes
- Complex nested structures

## Environment Variables

- `MARSHAL_API_KEY` - Your Marshal API key (format: `lg-m-...`)

## Example Output

### Validation Example
```
=== JSON Validation Example ===

Test 1: Validating valid JSON...
✓ JSON is valid!

Test 2: Validating invalid JSON (missing required field)...
✗ JSON is invalid:
  - Field 'age': required field missing

Test 3: Validating invalid JSON (wrong type)...
✗ JSON is invalid:
  - Field 'age': Expected type number, got string

=== Validation Complete ===
```

### Repair Example
```
=== JSON Repair Example ===

Test 1: Fixing JSON with unquoted property name...
Before: {name: "John Doe", age: 30}
After:  {"name": "John Doe", "age": 30}

=== JSON Repair Complete ===
All malformed JSON has been successfully repaired!
```

## Notes

- All examples use `context.Background()`, but you can use custom contexts with timeouts or cancellation
- Error handling is simplified in examples for clarity
- In production, handle errors appropriately and use structured logging
- Never commit your API key to version control

## Building Examples

You can build any example to create a standalone binary:

```bash
cd examples/validation
go build -o validation
./validation
```

## Testing with Custom Base URL

If you're testing against a local or staging server:

```bash
# Modify the example to use WithBaseURL option
client := marshal.NewClient(apiKey, marshal.WithBaseURL("http://localhost:8080"))
```
