# API Module Documentation

This document describes the use of the API module in the OneLogin Go SDK. The API module provides the basic HTTP methods (GET, POST, DELETE, PUT) for communicating with the OneLogin API, and handles the construction and sending of requests, as well as processing the responses.

## Client

The `Client` struct in the API module is responsible for making requests to the OneLogin API. It contains an `HttpClient` (for making the HTTP requests), an `Auth` (for handling authentication), and the `OLdomain` (the OneLogin domain to which requests are sent).

The `NewClient` function is used to create a new `Client`. It uses the `http.DefaultClient` as the `HttpClient`, creates a new `Authenticator`, and sets the `OLdomain` based on the `ONELOGIN_SUBDOMAIN` environment variable.

```go
client := api.NewClient()
```

## HTTP Methods

The `Client` struct provides the following methods for making HTTP requests:

- `Get`: This method makes a GET request to the specified path. It takes a `path` and a `queryParams` map as arguments, and returns the response body as a byte array.

```go
response, err := client.Get("/api/1/users", nil)
```

- `Post`: This method makes a POST request to the specified path. It takes a `path`, a `queryParams` map, and a `body` interface{} as arguments, and returns the response body as a byte array.

```go
response, err := client.Post("/api/1/users", nil, user)
```

- `Delete`: This method makes a DELETE request to the specified path. It takes a `path` and a `queryParams` map as arguments, and returns the response body as a byte array.

```go
response, err := client.Delete("/api/1/users/1", nil)
```

- `Put`: This method makes a PUT request to the specified path. It takes a `path`, a `queryParams` map, and a `body` interface{} as arguments, and returns the response body as a byte array.

```go
response, err := client.Put("/api/1/users/1", nil, updatedUser)
```

## Request

The `Request` struct represents an API request. It contains the `Url`, `Headers`, and `Body` of the request.

A new `Request` can be created using the `NewRequest` function, which takes a `url`, a `headers` map, and a `body` interface{} as arguments.

The `Send` method of the `Request` struct sends the API request and returns an `http.Response`.

```go
request := api.NewRequest("https://your-api-subdomain.onelogin.com/api/1/users", headers, body)
response, err := request.Send()
```

## Response

The `Response` struct represents an API response. It contains the `StatusCode`, `Body`, and `Headers` of the response.

A new `Response` can be created using the `NewResponse` function, which takes an `http.Response` as an argument.

```go
response, err := api.NewResponse(httpResponse)
```

Note: Make sure to check the `StatusCode` of the `Response` to ensure that the request was successful. A `StatusCode` in the 200-299 range indicates success, while a `StatusCode` in the 400-499 range indicates an error.
