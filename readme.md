# Onelogin Go SDK

This is the Onelogin SDK, a Go package that provides a convenient interface for interacting with Onelogin's API. The SDK aims to simplify the integration process and provide developers with an easy-to-use tool for managing authentication, making API requests, and handling responses.

**Note: This project is currently undergoing a rearchitecture, and the README will be updated with more information once the rearchitecture is complete.**

## Features

- Authentication: Obtain access tokens for authenticating API requests.
- API Requests: Construct and send requests to Onelogin's API.
- Response Handling: Parse and handle API responses.
- Error Management: Handle and recover from errors effectively.
- Data Models: Represent Onelogin entities and resources.
- Utilities: Provide utility functions and helper methods.

## Installation

To use the Onelogin SDK in your Go project, you need to have Go installed and set up. Then, you can install the SDK using the `go get` command:

```shell
go get github.com/onelogin/onelogin-go-sdk
```

**Note: Replace `github.com/onelogin/onelogin-go-sdk` with the actual import path of your project.**

## Usage

**Note: The usage examples provided below are based on the previous version of the SDK. Please refer to the updated documentation once the rearchitecture is complete.**

Here's a basic example demonstrating how to use the Onelogin SDK:

```go
package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	"github.com/onelogin/onelogin-go-sdk/pkg/onelogin"
)

func main() {
	// Create a new instance of the Authenticator
	auth := authentication.NewAuthenticator()

	// Authenticate and get the access token
	err := auth.GetToken()
	if err != nil {
		fmt.Printf("Failed to authenticate: %s\n", err)
		return
	}

	// Create a new instance of the Onelogin SDK
	ol := onelogin.NewOnelogin(auth)

	// Use the Onelogin SDK to make API calls
	// Example: Get user information
	user, err := ol.GetUser("user_id")
	if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
		return
	}

	// Print the user information
	fmt.Printf("User: %v\n", user)

	// Revoke the access token
	err = auth.RevokeToken(auth.GetAccessToken())
	if err != nil {
		fmt.Printf("Failed to revoke token: %s\n", err)
		return
	}

	fmt.Println("Token revoked successfully")
}
```

## Documentation

**Note: The documentation provided below is based on the previous version of the SDK. Please refer to the updated documentation once the rearchitecture is complete.**

Comprehensive documentation for the Onelogin SDK is available in the `docs` directory. The following documents provide detailed information on using the SDK and its various modules:

- `api.md`: Documentation for the API module, including request construction, communication, and response handling.
- `authentication.md`: Documentation for the authentication module, covering the process of obtaining and revoking access tokens.
- `error_handling.md`: Documentation for error handling, including information on error types and codes.
- `index.md`: Introduction and overview of the SDK, including goals and architecture.
- `models.md`: Documentation for the models module, describing the data models that represent Onelogin

entities and resources.

- `usage_examples.md`: Contains usage examples and code snippets demonstrating various SDK functionalities.

## Contributing

Contributions to the Onelogin SDK are welcome! If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request. We appreciate your feedback and contributions.

## Updates
