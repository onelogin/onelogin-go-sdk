# OneLogin Go SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/onelogin/onelogin-go-sdk)](https://goreportcard.com/report/github.com/onelogin/onelogin-go-sdk)
[![GoDoc](https://godoc.org/github.com/onelogin/onelogin-go-sdk?status.svg)](https://godoc.org/github.com/onelogin/onelogin-go-sdk)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A Go client library for the OneLogin API (v2). This SDK provides a convenient interface for Go applications to interact with OneLogin services.

## Installation

```bash
go get github.com/onelogin/onelogin-go-sdk/v4
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/onelogin/onelogin-go-sdk/v4"
)

func main() {
    client, err := onelogin.New()
    if err != nil {
        panic(err)
    }

    // Example: Get Users
    users, err := client.Users.List(nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Found %d users\n", len(users))
}
```

## Configuration

Configure the SDK using environment variables:

```bash
export ONELOGIN_CLIENT_ID="your-client-id"
export ONELOGIN_CLIENT_SECRET="your-client-secret"
export ONELOGIN_SUBDOMAIN="your-subdomain"
export ONELOGIN_TIMEOUT=15  # Optional, defaults to 15 seconds
```

## Features

- Full support for OneLogin API v2
- Automatic token management
- Configurable timeout and retry settings
- Comprehensive error handling
- Type-safe request/response structures

## Documentation

- [API Reference](https://godoc.org/github.com/onelogin/onelogin-go-sdk)
- [Contributing](CONTRIBUTING.md)

## Development

```bash
# Run tests
make test

# Run linters
make lint

# Run all checks
make check

# Generate docs
make docs
```

### Required Tools

For SDK development, you'll need:

- Go 1.19 or later
- golangci-lint (installed automatically via make)
- gosec (installed automatically via make)
- godoc (installed automatically via make)

The CI pipeline will install these tools automatically in the GitHub Actions environment.

## Support

Issues and feature requests can be filed on our [GitHub issue tracker](https://github.com/onelogin/onelogin-go-sdk/issues).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
