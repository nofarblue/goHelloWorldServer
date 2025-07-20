# Go Hello World Server

A production-ready HTTP server written in Go that provides a simple greeting API with comprehensive input validation, security features, and Docker support. This project demonstrates best practices for Go web services including input sanitization, comprehensive testing, and CI/CD integration.

## Features

- **RESTful API**: Simple HTTP endpoint for generating personalized greetings
- **Input Sanitization**: Advanced security features to prevent injection attacks
- **Input Validation**: Robust handling of edge cases and malformed input
- **Comprehensive Testing**: Extensive test suite covering all functionality and edge cases
- **Docker Support**: Containerized deployment with Alpine Linux base
- **CI/CD Integration**: Automated testing and builds with CircleCI
- **Graceful Shutdown**: Proper signal handling for clean server termination
- **Structured Logging**: Safe logging with input sanitization

## Security Features

This server includes several security enhancements:

- **Control Character Sanitization**: Removes newlines, tabs, and other control characters to prevent log injection
- **Input Length Limiting**: Restricts names to 100 characters maximum to prevent abuse
- **Whitespace Handling**: Properly trims and handles whitespace-only inputs
- **Safe Logging**: Sanitizes all logged inputs to prevent log injection attacks
- **XSS Prevention**: Strips potentially dangerous characters from user input

## Prerequisites

- Go 1.15 or higher
- Docker (optional, for containerized deployment)
- Git

## Installation

### Local Development

1. **Clone the repository**:
   ```bash
   git clone https://github.com/nofarblue/goHelloWorldServer.git
   cd goHelloWorldServer
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the application**:
   ```bash
   go build -o go-sample-app
   ```

### Docker Deployment

1. **Build the Docker image**:
   ```bash
   # Build the Go binary for Linux
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o go-sample-app
   
   # Build Docker image
   docker build -t go-hello-server .
   ```

2. **Run the container**:
   ```bash
   docker run -p 8080:8080 go-hello-server
   ```

## Usage

### Starting the Server

**Local development**:
```bash
./go-sample-app
```

The server will start on port 8080 and display:
```
2019/02/03 11:38:11 Starting Server
```

### API Endpoints

#### GET /

**Description**: Returns a personalized greeting message.

**Parameters**:
- `name` (query parameter, optional): The name to include in the greeting

**Examples**:

```bash
# Basic greeting
curl http://localhost:8080
# Response: Hello, Guest

# Personalized greeting
curl http://localhost:8080?name=Nofar
# Response: Hello, Nofar

# Whitespace handling
curl "http://localhost:8080?name=   John   "
# Response: Hello, John

# Empty name handling
curl "http://localhost:8080?name="
# Response: Hello, Guest
```

**Security Examples**:

The server automatically handles various security scenarios:

```bash
# Control character removal
curl "http://localhost:8080?name=John%0AAdmin"
# Response: Hello, JohnAdmin

# Length limiting (names over 100 characters are truncated)
curl "http://localhost:8080?name=$(python3 -c 'print("a"*150)')"
# Response: Hello, followed by exactly 100 'a' characters

# Special characters are preserved
curl "http://localhost:8080?name=Jane!@#$%"
# Response: Hello, Jane!@#$%
```

## Testing

This project includes comprehensive test coverage for all functionality and edge cases.

### Running Tests

```bash
# Run all tests
go test -v

# Run tests with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Coverage

The test suite includes:

- **Basic functionality tests**: Standard greeting generation
- **Edge case handling**: Empty inputs, whitespace-only inputs
- **Security tests**: Control character injection, length limiting
- **Input validation tests**: Whitespace trimming, special characters
- **Utility function tests**: Sanitization helpers and logging functions

**Test Results**:
```
=== RUN   TestGreetingSpecificJohn
--- PASS: TestGreetingSpecificJohn (8.00s)
=== RUN   TestGreetingSpecificDemo  
--- PASS: TestGreetingSpecificDemo (0.00s)
=== RUN   TestGreetingDefault
--- PASS: TestGreetingDefault (0.00s)
=== RUN   TestGreeting_WhitespaceOnly
--- PASS: TestGreeting_WhitespaceOnly (0.00s)
=== RUN   TestGreeting_LongName
--- PASS: TestGreeting_LongName (0.00s)
=== RUN   TestGreeting_NewlineInjection
--- PASS: TestGreeting_NewlineInjection (0.00s)
=== RUN   TestGreeting_SpecialSymbols
--- PASS: TestGreeting_SpecialSymbols (0.00s)
=== RUN   TestGreeting_TrimWhitespace
--- PASS: TestGreeting_TrimWhitespace (0.00s)
=== RUN   TestSanitizeControlChars
--- PASS: TestSanitizeControlChars (0.00s)
=== RUN   TestSanitizeForLogging
--- PASS: TestSanitizeForLogging (0.00s)
PASS
ok      github.com/harness/go-sample-app    8.003s
```

## CI/CD

This project uses CircleCI for continuous integration and deployment.

### Pipeline Features

- **Automated Testing**: Runs full test suite on every commit
- **Dependency Caching**: Optimized build times with Go module caching  
- **Test Reporting**: JUnit XML format test results
- **Multi-environment Support**: Tests across different Go versions

### Configuration

The CI/CD pipeline is configured in `.circleci/config.yml` and includes:

1. **Build Environment**: Go 1.16 with CircleCI convenience image
2. **Dependency Management**: Automatic Go module download and caching
3. **Test Execution**: Comprehensive test suite with formatted output
4. **Artifact Storage**: Test results stored for analysis

## Project Structure

```
.
├── .circleci/              # CircleCI configuration
│   └── config.yml         # CI/CD pipeline definition
├── .gitignore             # Git ignore rules
├── Dockerfile             # Container image definition
├── Dockerfile.build       # Build container definition
├── README.md              # This file
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── hello_server.go        # Main server implementation
└── hello_server_test.go   # Comprehensive test suite
```

## Dependencies

- **gorilla/mux v1.6.2**: HTTP router and URL matcher
- **gorilla/context v1.1.1**: Request context management

## Contributing

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/new-feature`
3. **Write tests**: Ensure your changes include appropriate test coverage
4. **Run tests**: `go test -v` to verify all tests pass
5. **Commit changes**: `git commit -am 'Add new feature'`
6. **Push to branch**: `git push origin feature/new-feature`
7. **Submit pull request**: Create a PR with a clear description of changes

### Code Standards

- Follow Go best practices and conventions
- Include comprehensive tests for new functionality
- Ensure all tests pass before submitting PR
- Document any new features or breaking changes
- Maintain backwards compatibility where possible

## Production Deployment

### Environment Variables

The server can be configured using environment variables:

- `PORT`: Server port (default: 8080)
- `HOST`: Server host (default: all interfaces)

### Health Checks

The server includes graceful shutdown handling for production deployments:

```go
// Graceful shutdown on SIGTERM/SIGINT
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt, syscall.SIGTERM)
```

### Monitoring

The server logs all requests with sanitized input for monitoring and debugging:

```
2019/02/03 11:38:11 Received request for John
2019/02/03 11:38:15 Received request for <empty>
```

## License

This project is open source. Please check the repository for license details.

## Maintainer

**Nofar Bluestein** - nofarb@gmail.com 
