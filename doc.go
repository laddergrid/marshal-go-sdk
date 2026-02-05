// Package marshal provides a Go SDK for the Marshal Engine API.
//
// Marshal Engine is a JSON validation and repair service that helps you
// ensure your JSON data conforms to OpenAPI specifications and automatically
// fixes malformed JSON strings.
//
// # Installation
//
// Install the package using go get:
//
//	go get github.com/laddergrid/marshal-go-sdk
//
// # Quick Start
//
// Create a client with your API key and perform operations:
//
//	client := marshal.NewClient("lg-m-your-api-key")
//	resp, err := client.Ping(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp.Message) // "pong"
//
// # Authentication
//
// The SDK requires an API key for all operations. Pass your API key when
// creating the client:
//
//	client := marshal.NewClient("lg-m-your-api-key")
//
// For security, store your API key in an environment variable:
//
//	apiKey := os.Getenv("MARSHAL_API_KEY")
//	client := marshal.NewClient(apiKey)
//
// # JSON Validation
//
// Validate JSON against an OpenAPI specification:
//
//	spec := `{
//	  "type": "object",
//	  "properties": {
//	    "name": {"type": "string"},
//	    "age": {"type": "number"}
//	  },
//	  "required": ["name"]
//	}`
//
//	json := `{"name": "John Doe", "age": 30}`
//
//	result, err := client.ValidateJSON(ctx, spec, json, "createUser")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if !result.IsValid {
//	    for _, err := range result.Errors {
//	        fmt.Printf("%s: %s\n", err.Field, err.Error)
//	    }
//	}
//
// # JSON Repair
//
// Fix malformed JSON automatically:
//
//	malformed := `{name: "John", age: 30}`
//	result, err := client.FixJSON(ctx, malformed)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result.FixedJSON) // {"name": "John", "age": 30}
//
// # Context Support
//
// All methods accept a context for timeout and cancellation:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	result, err := client.Ping(ctx)
//
// # Error Handling
//
// API errors can be type-asserted for detailed information:
//
//	if err != nil {
//	    if apiErr, ok := err.(*marshal.APIError); ok {
//	        fmt.Printf("Status: %d, Message: %s\n",
//	            apiErr.StatusCode, apiErr.Message)
//	    }
//	}
//
// For more examples, see the examples directory in the repository.
package marshal
