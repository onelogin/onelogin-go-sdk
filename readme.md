# Onelogin Go SDK

This is the Onelogin SDK, a Go package that provides a convenient interface for interacting with Onelogin's API. The SDK simplifies the integration process by providing developers an easy-to-use tool for managing authentication, making API requests, and handling responses.

## Support

OneLogin by One Identity open source projects are supported through [OneLogin GitHub issues](https://github.com/onelogin/onelogin-go-sdk/issues). This includes all scripts, plugins, SDKs, modules, code snippets or other solutions. For assistance with any OneLogin by One Identity GitHub project, please raise a new Issue on the [OneLogin GitHub issues](https://github.com/onelogin/onelogin-go-sdk/issues) page. Requests for assistance made through official One Identity Support will be referred back to GitHub where those requests can benefit all users.


## Features

- **Authentication**: Obtain and revoke access tokens for authenticating API requests using environment variables.
- **API Requests**: Construct and send requests to Onelogin's API.
- **Response Handling**: Parse and handle API responses.
- **Error Management**: Handle and recover from errors effectively.
- **Data Models**: Represent Onelogin entities and resources.
- **Utilities**: Provide utility functions and helper methods.

## Supported APIs

- [API Authorization](https://developers.onelogin.com/api-docs/2/api-authorization/overview)
- [Apps](https://developers.onelogin.com/api-docs/2/apps)
- [App Rules](https://developers.onelogin.com/api-docs/2/app-rules)
- [Branding](https://developers.onelogin.com/api-docs/2/branding)
- [Connectors](https://developers.onelogin.com/api-docs/2/connectors/list-connectors)
- [Events](https://developers.onelogin.com/api-docs/2/events)
- [Groups](https://developers.onelogin.com/api-docs/2/groups)
- [Invite Links](https://developers.onelogin.com/api-docs/2/invite-links)
- [Privileges](https://developers.onelogin.com/api-docs/1/privileges)
- [Roles](https://developers.onelogin.com/api-docs/2/roles)
- [SAML Assertions](https://developers.onelogin.com/api-docs/2/saml-assertions)
- [Smart Hooks](https://developers.onelogin.com/api-docs/2/smart-hooks)
- [Users](https://developers.onelogin.com/api-docs/2/users)
- [User Mappings](https://developers.onelogin.com/api-docs/2/user-mappings)
- [Vigilance AI](https://developers.onelogin.com/api-docs/2/vigilance/train)

## Partially Support APIs

- [Login Pages](https://developers.onelogin.com/api-docs/2/login-page)
- [MFA](https://developers.onelogin.com/api-docs/2/multi-factor-authentication)

## Installation

To use the Onelogin SDK in your Go project, you need to have Go installed and set up. Then, you can install the SDK using the `go get` command:

```shell
go get github.com/onelogin/onelogin-go-sdk/v4
```

## Requirements

The SDK expects three environment variables to be set for authentication:

- `ONELOGIN_CLIENT_ID`
- `ONELOGIN_CLIENT_SECRET`
- `ONELOGIN_SUBDOMAIN`
- `ONELOGIN_TIMEOUT`

These variables are used by the Authenticator for authentication with the OneLogin API. The Authenticator retrieves an access token using these credentials, which is then used for API requests. You can set these variables directly in your shell or in the environment of the program that will be using this SDK.

In a Unix/Linux shell, you can use the `export` command to set these variables:

```shell
export ONELOGIN_CLIENT_ID=your_client_id
export ONELOGIN_CLIENT_SECRET=your_client_secret
export ONELOGIN_SUBDOMAIN=your_subdomain
export ONELOGIN_TIMEOUT=15
```

In a Go program, you can use the `os` package to set these variables:

```go
os.Setenv("ONELOGIN_CLIENT_ID", "your_client_id")
os.Setenv("ONELOGIN_CLIENT_SECRET", "your_client_secret")
os.Setenv("ONELOGIN_SUBDOMAIN", "your_subdomain")
os.Setenv("ONELOGIN_TIMEOUT", 15)
```

Please ensure these variables are set before attempting to use the SDK to make API requests.

## Usage

Here's an example demonstrating how to use the Onelogin SDK:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {

	// As mentioned earlier, ensure that the environment variables (ONELOGIN_CLIENT_ID, ONELOGIN_CLIENT_SECRET, ONELOGIN_SUBDOMAIN)
    // are correctly loaded prior to initialization, as they are necessary for creating a new instance of the OneLogin SDK client.


	// Create a new OneLogin SDK client
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		log.Fatalf("Error creating OneLogin SDK client: %v", err)
	}

	userQuery := models.UserQuery{}
	userList, err := client.GetUsers(&userQuery)
	if err != nil {
		fmt.Println("Failed to get user:", err)
		return
	}

	// Marshal the result into JSON for printing
	userJSON, err := json.MarshalIndent(userList, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling the result to JSON: %v", err)
		return
	}

	// Print the User details in JSON format
	fmt.Println("User Details:")
	fmt.Println(string(userJSON))

	appQuery := models.AppQuery{}
	appList, err := client.GetApps(&appQuery)
	if err != nil {
		fmt.Println("Failed to get app list:", err)
		return
	}

	// Marshal the result into JSON for printing
	appJSON, err := json.MarshalIndent(appList, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling the result to JSON: %v", err)
		return
	}

	// Print the App details in JSON format
	fmt.Println("App Details:")
	fmt.Println(string(appJSON))
}
```

## Releasing

When releasing a new version of the SDK:

1. Update the version number in `pkg/onelogin/version.go` following semantic versioning:
   - MAJOR version for incompatible API changes
   - MINOR version for backwards-compatible functionality additions
   - PATCH version for backwards-compatible bug fixes

2. Commit the version change:
   ```shell
   git add pkg/onelogin/version.go
   git commit -m "Bump version to x.y.z"
   git tag vx.y.z
   git push origin main --tags
   ```

## Documentation

[Comprehensive documentation](docs/index.md) for the Onelogin SDK is available in the `docs` directory. The following documents provide detailed information on using the SDK and its various modules:

- `api.md`: Documentation for the API module, including request construction, communication, and response handling.
- `authentication.md`: Detailed documentation for the Authentication module, including the process of obtaining and revoking access tokens as well as retrieving the token for other applications or services.
- `error_handling.md`: Documentation for error handling, including information on error types and codes.
- `index.md`: Introduction and overview of the SDK, including goals and architecture.
- `models.md`: Documentation for the models module, describing the data models that represent Onelogin entities and resources.
- `usage_examples.md`: Contains usage examples and code snippets demonstrating various SDK functionalities.

## Contributing

Contributions to the Onelogin SDK are welcome! If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request. We appreciate your feedback and contributions.