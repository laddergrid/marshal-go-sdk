package main

import (
	"context"
	"fmt"
	"log"
	"os"

	marshal "github.com/laddergrid/marshal-go-sdk"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("MARSHAL_API_KEY")
	if apiKey == "" {
		log.Fatal("MARSHAL_API_KEY environment variable is required")
	}

	// Create client with API key
	client := marshal.NewClient(apiKey)

	fmt.Println("=== JSON Validation Example ===")
	fmt.Println()

	// Define OpenAPI specification
	openAPISpec := `{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "minLength": 1
    },
    "age": {
      "type": "number",
      "minimum": 0,
      "maximum": 150
    },
    "email": {
      "type": "string",
      "format": "email"
    }
  },
  "required": ["name", "age"]
}`

	// Test 1: Valid JSON
	fmt.Println("Test 1: Validating valid JSON...")
	validJSON := `{"name": "John Doe", "age": 30, "email": "john@example.com"}`

	result1, err := client.ValidateJSON(
		context.Background(),
		openAPISpec,
		validJSON,
		"createUser",
	)
	if err != nil {
		log.Fatalf("Validation request failed: %v", err)
	}

	if result1.IsValid {
		fmt.Println("✓ JSON is valid!")
	} else {
		fmt.Println("✗ JSON is invalid:")
		for _, err := range result1.Errors {
			fmt.Printf("  - Field '%s': %s\n", err.Field, err.Error)
		}
	}

	// Test 2: Invalid JSON (missing required field)
	fmt.Println("\nTest 2: Validating invalid JSON (missing required field)...")
	invalidJSON1 := `{"name": "Jane Doe"}`

	result2, err := client.ValidateJSON(
		context.Background(),
		openAPISpec,
		invalidJSON1,
		"createUser",
	)
	if err != nil {
		log.Fatalf("Validation request failed: %v", err)
	}

	if result2.IsValid {
		fmt.Println("✓ JSON is valid!")
	} else {
		fmt.Println("✗ JSON is invalid:")
		for _, err := range result2.Errors {
			fmt.Printf("  - Field '%s': %s\n", err.Field, err.Error)
		}
	}

	// Test 3: Invalid JSON (wrong type)
	fmt.Println("\nTest 3: Validating invalid JSON (wrong type)...")
	invalidJSON2 := `{"name": "Bob Smith", "age": "thirty"}`

	result3, err := client.ValidateJSON(
		context.Background(),
		openAPISpec,
		invalidJSON2,
		"createUser",
	)
	if err != nil {
		log.Fatalf("Validation request failed: %v", err)
	}

	if result3.IsValid {
		fmt.Println("✓ JSON is valid!")
	} else {
		fmt.Println("✗ JSON is invalid:")
		for _, err := range result3.Errors {
			fmt.Printf("  - Field '%s': %s\n", err.Field, err.Error)
		}
	}

	fmt.Println("\n=== Validation Complete ===")
}
