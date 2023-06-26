# API Module Documentation

The API module is a crucial part of the OneLogin Go SDK. It is responsible for constructing and sending HTTP requests to the OneLogin API, and processing the responses. This document will explain the main components of the API module and their role in the SDK.

## Client

The `Client` struct is the primary interface for interacting with the OneLogin API. It uses an `HttpClient` to make HTTP requests, an `Auth` object to handle authentication, and the `OLdomain` (OneLogin domain) for routing the requests.

### `NewClient` Function

The `NewClient` function creates a new `Client` and initializes it with the `http.DefaultClient` as the `HttpClient`, a new `Authenticator` for `Auth`, and sets the `OLdomain` from the `ONELOGIN_SUBDOMAIN` environment variable. This function is typically called at the start of the program to instantiate the client, which is then used to make API calls.

### `newRequest` Function

This function creates a new HTTP request with the specified method, path, query parameters, and request body. It is a helper function, used by the HTTP methods (Get, Post, Delete, Put) of the `Client` to construct a new request. The function takes the method type (GET, POST, etc.), the API path, an object for query parameters, and a body for the request, and returns an HTTP request that is ready to be sent.

### `sendRequest` Function

This function sends an HTTP request and returns the HTTP response. It is used by the HTTP methods (Get, Post, Delete, Put) of the `Client` to send requests. This function also checks the response status code, and if it detects a `http.StatusUnauthorized` (HTTP 401), it attempts to refresh the token and retry the request.

## HTTP Methods

The `Client` struct provides the following methods for making HTTP requests:

- `Get`: Makes a GET request to the specified path with optional query parameters. It is used to retrieve information from the OneLogin API.
- `Post`: Makes a POST request to the specified path. It sends data to the OneLogin API to create a new resource.
- `Delete`: Makes a DELETE request to the specified path. It is used to delete a resource from the OneLogin API.
- `Put`: Makes a PUT request to the specified path. It is used to update a resource in the OneLogin API.

Each of these methods uses the `newRequest` function to create the HTTP request, and the `sendRequest` function to send the request and retrieve the response. These methods make the process of interacting with the OneLogin API simpler and more intuitive.

## Authenticator

The `Authenticator` interface is used for handling authentication. It uses the `GetToken` method for retrieving authentication tokens. The tokens are needed for authenticating requests to the OneLogin API.

The `Authenticator` is initialized with the `NewAuthenticator` function and the token is generated with the `GenerateToken` method within the `NewClient` function. If a request is unauthorized (HTTP 401), the token is refreshed using the `GenerateToken` method in the `sendRequest` function.

In summary, the API module simplifies the process of interacting with the OneLogin API by encapsulating the details of creating, sending, and processing HTTP requests. It uses environment variables for the API credentials and handles error scenarios such as unauthorized requests and token refresh. It forms the backbone of the OneLogin Go SDK, providing a streamlined interface for making API calls.