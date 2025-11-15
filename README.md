# GoHub

A github repo maker in the terminal written in Go.

## Project Structure

```
.
├── cmd/           # Main applications for this project
├── pkg/           # Library code that's ok to use by external applications
├── internal/      # Private application and library code
├── config/        # Configuration files
├── docs/          # Documentation
├── scripts/       # Various scripts
├── test/          # Additional external test files
├── assets/        # Static assets
├── deployments/   # Deployment configurations
├── go.mod         # Go module file
├── go.sum         # Go module checksum
├── main.go        # Entry point
└── README.md      # This file
```

## Getting Started

```bash
# Run the application
go run main.go

# Build the application
go build -o bin/app main.go

# Run tests
go test ./...
```
