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

	fmt.Println("=== JSON Repair Example ===")
	fmt.Println()

	// Test 1: Missing quotes on property name
	fmt.Println("Test 1: Fixing JSON with unquoted property name...")
	malformed1 := `{name: "John Doe", age: 30}`
	fmt.Printf("Before: %s\n", malformed1)

	result1, err := client.FixJSON(context.Background(), malformed1)
	if err != nil {
		log.Fatalf("Repair request failed: %v", err)
	}

	fmt.Printf("After:  %s\n\n", result1.FixedJSON)

	// Test 2: Missing closing brace
	fmt.Println("Test 2: Fixing JSON with missing closing brace...")
	malformed2 := `{"name": "Jane Doe", "age": 25, "city": "New York"`
	fmt.Printf("Before: %s\n", malformed2)

	result2, err := client.FixJSON(context.Background(), malformed2)
	if err != nil {
		log.Fatalf("Repair request failed: %v", err)
	}

	fmt.Printf("After:  %s\n\n", result2.FixedJSON)

	// Test 3: Trailing comma
	fmt.Println("Test 3: Fixing JSON with trailing comma...")
	malformed3 := `{"name": "Bob Smith", "age": 35,}`
	fmt.Printf("Before: %s\n", malformed3)

	result3, err := client.FixJSON(context.Background(), malformed3)
	if err != nil {
		log.Fatalf("Repair request failed: %v", err)
	}

	fmt.Printf("After:  %s\n\n", result3.FixedJSON)

	// Test 4: Single quotes instead of double quotes
	fmt.Println("Test 4: Fixing JSON with single quotes...")
	malformed4 := `{'name': 'Alice Johnson', 'age': 28}`
	fmt.Printf("Before: %s\n", malformed4)

	result4, err := client.FixJSON(context.Background(), malformed4)
	if err != nil {
		log.Fatalf("Repair request failed: %v", err)
	}

	fmt.Printf("After:  %s\n\n", result4.FixedJSON)

	// Test 5: Complex nested JSON with multiple issues
	fmt.Println("Test 5: Fixing complex JSON with multiple issues...")
	malformed5 := `{
  name: "Charlie Brown",
  age: 40,
  address: {
    street: "123 Main St",
    city: 'Springfield',
  },
  hobbies: ['reading', 'coding', 'gaming',]
`
	fmt.Printf("Before:\n%s\n", malformed5)

	result5, err := client.FixJSON(context.Background(), malformed5)
	if err != nil {
		log.Fatalf("Repair request failed: %v", err)
	}

	fmt.Printf("After:\n%s\n\n", result5.FixedJSON)

	fmt.Println("=== JSON Repair Complete ===")
	fmt.Println("All malformed JSON has been successfully repaired!")
}
