# Onelogin Go SDK Documentation

This documentation provides detailed information on how to use the SDK to interact with the OneLogin platform.

## Table of Contents

- [API Reference](api.md): Learn about the available API endpoints and how to make requests using the SDK.
- [Authentication](authentication.md): Understand the authentication mechanisms supported by the SDK and how to authenticate your requests.
- [Error Handling](error_handling.md): Explore the different types of errors that can occur during SDK usage and how to handle them properly.
- [SDK Models](models.md): Get familiar with the models used by the SDK to represent data and interact with the API.
- [Usage Examples](usage_examples.md): Find example code snippets and scenarios to help you understand how to use the SDK effectively.

## Additional Resources

- [Design](../Design.md): Dive into the design considerations and architecture of the SDK.
- [License](../LICENSE): View the license agreement for the SDK.
- [README](../README.md): Read the SDK's introductory information, installation instructions, and basic usage guidelines.

## Code Structure

The SDK follows a modular structure with the following key components:

- `cmd/main.go`: Entry point of the SDK application.
- `internal/api`: Contains the implementation of the API client, request handling, and response parsing.
- `internal/authentication`: Provides authentication mechanisms and related functionality.
- `internal/error`: Defines various error types used by the SDK.
- `internal/models`: Contains the model definitions used to represent data exchanged with the API.
- `internal/utilities`: Includes utility functions and helper methods.
- `pkg/onelogin`: Contains the main implementation of the OneLogin SDK.

## Getting Started

To get started with the SDK, follow the installation instructions outlined in the [README](../README.md) file. Once the SDK is installed, refer to the usage examples in [Usage Examples](usage_examples.md) to understand how to use the SDK to perform various operations.
