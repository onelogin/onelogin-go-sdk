# SDK Design Document

## Introduction

This design document outlines the architecture and implementation approach for building a new SDK for Onelogin's API. The SDK aims to provide developers with a convenient and user-friendly interface to interact with Onelogin's services.

## Goals

- Modularity: Design the SDK with separate modules for authentication, API request handling, data models, error handling, and utility functions.
- Encapsulation: Define well-defined interfaces for each module to encapsulate their functionality and hide implementation details.
- Simplified Interface: Use the Facade pattern to provide a unified and simplified interface for users of the SDK, shielding them from underlying complexities.
- Customizability: Implement dependency injection to allow users to customize or replace certain components, such as authentication or HTTP client implementations.

## Architecture Overview

The SDK architecture will be based on a combination of Modular Architecture, Facade Pattern, and Dependency Injection.

### Modular Architecture

Identify the following functional components for the SDK:

1. Authentication: Responsible for handling authentication mechanisms.
2. API: Handles API request construction, communication, and response handling.
3. Models: Represents entities and resources in the Onelogin API.
4. Error Handling: Manages errors and provides error types and codes.
5. Utilities: Contains utility functions and helper methods.

Create separate modules or packages for each component to ensure clear separation of concerns and better organization.

### Facade Pattern

Design a facade module or package, "onelogin," that acts as the entry point for SDK users. The facade will provide a simplified interface and shield users from the complexities of underlying modules.

The facade will offer methods for authentication, making API requests, handling responses, and error management. It will internally coordinate and interact with the authentication and API modules to fulfill these functionalities.

### Dependency Injection

Identify components within the SDK that can benefit from customization or replacement by users, such as the authentication module or the HTTP client implementation.

Design interfaces for these components, such as "Authenticator" and "HTTPClient," specifying the required methods or functionality. Implement default implementations for these interfaces, which will be used if users do not provide custom implementations.

Allow users to provide their custom implementations of these interfaces by accepting them as parameters in relevant methods or constructors. Use dependency injection to inject the appropriate implementation instances into the modules or components that require them.

## Implementation Steps

1. Set up the project structure with separate directories for each module: authentication, api, models, error, and utilities.
2. Implement the authentication module:
   - Create an "auth.go" file within the authentication package to handle authentication mechanisms (e.g., API key, OAuth, or SAML).
   - Implement methods to authenticate and generate authentication tokens.
3. Implement the API module:
   - Create a "client.go" file within the api package to handle API request construction, communication, and response handling.
   - Implement methods for making HTTP requests, handling response parsing, and error handling.
4. Implement the models module:
   - Create files (e.g., "user.go" and "group.go") within the models package to define the data models that represent Onelogin entities and resources.
   - Define appropriate struct types, fields, and methods for convenient data access and manipulation.
5. Implement the error handling module:
   - Create an "error.go" file within the error package to define error types and codes that reflect possible failures and error scenarios in the Onelogin API.
   - Implement error handling mechanisms that allow users to handle and recover from errors effectively.
6. Implement the utilities module:
   - Create a "util.go" file within the utilities package to include utility functions and helper methods used across the SDK.
   - Implement commonly used functions, such as data serialization, date/time formatting, or string manipulation.

7. Design and implement the facade module:
   - Create an "onelogin.go" file within the onelogin package to define the facade module.
   - Implement methods that provide a simplified interface for users, delegating calls to the relevant modules.
   - Coordinate authentication and API request handling using the appropriate modules.
8. Apply dependency injection:
   - Identify components that can be customized or replaced by users, such as the authentication module or the HTTP client implementation.
   - Design interfaces (e.g., "Authenticator" and "HTTPClient") within their respective modules and implement default implementations.
   - Allow users to provide their custom implementations by accepting them as parameters in relevant methods or constructors.
   - Use dependency injection to inject the appropriate implementation instances into the modules or components that require them.
9. Write comprehensive documentation:
   - Create documentation files (e.g., index.md, authentication.md, api.md, models.md, error_handling.md, usage_examples.md) within the docs directory.
   - Provide clear explanations of the SDK's features, functionalities, and usage examples.
   - Include code samples and guidelines to help users integrate and utilize the SDK efficiently.
   - Document any customization options and how to use dependency injection.
10. Conduct thorough testing:
    - Write unit tests for each module, covering different functionalities and edge cases.
    - Create test files (e.g., authentication_test.go, api_test.go, models_test.go, error_handling_test.go) within the tests directory.
    - Ensure that the SDK behaves as expected and handles errors appropriately.
11. Continuously iterate and gather feedback from users, making improvements and addressing any issues that arise.

## Conclusion

By following this architecture pattern, we aim to develop a modular, user-friendly SDK for Onelogin's API. The combination of Modular Architecture, Facade Pattern, and Dependency Injection will ensure maintainability, extensibility, and ease of use for developers integrating with Onelogin.
