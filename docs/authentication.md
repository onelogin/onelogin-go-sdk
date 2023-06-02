# Authentication Module Documentation

The Authentication module provides functionalities to handle the authentication process with the OneLogin API. This process includes generating and revoking access tokens, which are necessary for making authenticated requests to the API. The module is implemented in the `authentication.go` file in the `authentication` package.

## Authenticator Struct

The `Authenticator` struct represents an authenticator object that can handle authentication processes. It holds the `accessToken` and `OLDomain` as fields:

```go
type Authenticator struct {
	accessToken *string
	OLDomain    *string
}
```

## NewAuthenticator Function

The `NewAuthenticator` function is used to create a new `Authenticator` instance:

```go
func NewAuthenticator() *Authenticator {
	var domain = Domain
	var token string = ""
	return &Authenticator{
		&token, &domain,
	}
}
```

## GenerateToken Function

The `GenerateToken` function is used to generate a new access token. It reads the `ONELOGIN_CLIENT_ID` and `ONELOGIN_CLIENT_SECRET` environment variables, creates an authentication request, sends it, and handles the response. The newly generated access token is stored in the `Authenticator` instance and also returned by the function:

```go
func (a *Authenticator) GenerateToken() (string, error) {
    // implementation details
}
```

## RevokeToken Function

The `RevokeToken` function is used to revoke an existing access token. Similar to `GenerateToken`, it reads the `ONELOGIN_CLIENT_ID` and `ONELOGIN_CLIENT_SECRET` environment variables, creates a revocation request, sends it, and handles the response:

```go
func (a *Authenticator) RevokeToken(token, domain *string) error {
    // implementation details
}
```

## GetToken Function

The `GetToken` function is used to retrieve the current access token from the `Authenticator` instance:

```go
func (a *Authenticator) GetToken() (string, error) {
    return *a.accessToken, nil
}
```

## Error Handling

Errors in the Authentication module are represented using the `olError` package, and they can be of the following types:

- `AuthenticationError`: Represents an error that occurred during the authentication process, e.g., missing environment variables, failed authentication request, failed token generation or revocation.
- `RequestError`: Represents an error that occurred while creating or sending an HTTP request.
- `SerializationError`: Represents an error that occurred while marshalling or unmarshalling JSON data.

Each error type is associated with a specific error message that provides more details about the error.

## Conclusion

The Authentication module plays a crucial role in using the OneLogin Go SDK, as it handles the process of generating and revoking access tokens, which are necessary for making authenticated requests to the API. By understanding how this module works, you can handle authentication in your applications more effectively.
