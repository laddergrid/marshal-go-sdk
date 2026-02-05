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

	// Create a new client with your API key
	client := marshal.NewClient(apiKey)

	// Perform a health check
	resp, err := client.Ping(context.Background())
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	fmt.Printf("API Status: %s\n", resp.Message)
	fmt.Println("The Marshal API is ready to use!")
}
