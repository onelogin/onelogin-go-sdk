# Contributing to OneLogin Go SDK

We love your input! We want to make contributing to OneLogin Go SDK as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Development Setup

1. Install Go 1.19 or later
2. Clone the repository
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Install development tools:
   ```bash
   # Install golangci-lint
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   
   # Install gosec
   go install github.com/securego/gosec/v2/cmd/gosec@latest
   ```

## Testing

Run the full test suite:
```bash
make test
```

Run linters:
```bash
make lint
```

Run all checks (recommended before submitting PR):
```bash
make check
```

## Pull Request Process

1. Update the README.md with details of changes to the interface, if applicable.
2. Update the documentation with details of any new functionality.
3. The PR will be merged once you have the sign-off of at least one maintainer.

## Coding Standards

- Follow standard Go formatting rules (use `gofmt`)
- Write tests for new code
- Document exported functions and types
- Keep functions focused and small
- Use meaningful variable names
- Add comments for complex logic

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/master/CONTRIBUTING.md). 