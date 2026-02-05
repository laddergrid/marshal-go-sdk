# Contributing to Marshal Go SDK

Thank you for your interest in contributing to the Marshal Go SDK! This document provides guidelines and instructions for contributing.

## Development Setup

1. Fork the repository
2. Clone your fork:
```bash
git clone https://github.com/YOUR_USERNAME/marshal-go-sdk.git
cd marshal-go-sdk
```

3. Install dependencies (if any):
```bash
go mod download
```

## Making Changes

1. Create a new branch for your feature or bugfix:
```bash
git checkout -b feature/my-new-feature
```

2. Make your changes following the coding standards below

3. Run tests and ensure everything passes:
```bash
go test ./...
go vet ./...
go fmt ./...
```

4. Commit your changes with a clear commit message:
```bash
git commit -m "Add feature: description of your changes"
```

5. Push to your fork:
```bash
git push origin feature/my-new-feature
```

6. Open a Pull Request with a clear description of your changes

## Coding Standards

### Go Style

- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Run `go vet` to catch common mistakes
- Write clear, idiomatic Go code

### Documentation

- Add comments for all exported types, functions, and methods
- Use complete sentences in comments
- Start comments with the name of the thing being described
- Include examples in documentation when helpful

Example:
```go
// ValidateJSON validates a JSON string against an OpenAPI specification.
// It returns a ValidateResponse indicating whether the JSON is valid and
// any validation errors found.
func (c *Client) ValidateJSON(ctx context.Context, openAPISpec, jsonString, operationID string) (*ValidateResponse, error) {
    // implementation
}
```

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting a PR
- Aim for good test coverage
- Use table-driven tests when appropriate

Example:
```go
func TestValidateJSON(t *testing.T) {
    tests := []struct {
        name    string
        spec    string
        json    string
        wantErr bool
    }{
        // test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Error Handling

- Return errors instead of panicking
- Use descriptive error messages
- Wrap errors with context when appropriate:
```go
if err != nil {
    return fmt.Errorf("failed to validate JSON: %w", err)
}
```

### Context

- All API methods should accept `context.Context` as the first parameter
- Respect context cancellation and timeouts
- Pass context through the call chain

## Commit Message Guidelines

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests when relevant

Example:
```
Add JSON schema validation support

- Implement schema validation endpoint
- Add comprehensive error messages
- Include examples in documentation

Fixes #123
```

## Pull Request Process

1. Ensure your code follows the coding standards
2. Update the README.md with details of changes if needed
3. Add examples if introducing new features
4. Update documentation and comments
5. The PR will be merged once you have the sign-off of a maintainer

## Types of Contributions

### Bug Reports

- Use the GitHub issue tracker
- Describe the bug clearly with steps to reproduce
- Include error messages and stack traces if applicable
- Mention your Go version and OS

### Feature Requests

- Use the GitHub issue tracker
- Clearly describe the feature and its use case
- Explain why this feature would be useful
- Be open to discussion and feedback

### Code Contributions

We welcome contributions for:
- Bug fixes
- New features
- Performance improvements
- Documentation improvements
- Example code
- Test coverage improvements

## Questions?

If you have questions about contributing, feel free to:
- Open an issue for discussion
- Contact the maintainers

## License

By contributing to Marshal Go SDK, you agree that your contributions will be licensed under the GNU General Public License v3.0.

Thank you for contributing!
