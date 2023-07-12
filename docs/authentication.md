# Authentication Module Documentation

The Authentication module provides functionalities to handle the authentication process with the OneLogin API in the Go SDK. This process includes generating and revoking access tokens, which are necessary for making authenticated requests to the API. The module is implemented in the `authentication.go` file in the `authentication` package.

## Authenticator Struct

The `Authenticator` struct represents an authenticator object that can handle authentication processes. It holds the `accessToken` as a field. The `accessToken` field stores the generated access token used for making authenticated API calls.

```go
type Authenticator struct {
	accessToken string
}
```

## NewAuthenticator Function

The `NewAuthenticator` function is used to create a new `Authenticator` instance. It does not require any arguments and it initializes a new Authenticator with an empty `accessToken`.

```go
func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}
```

## GenerateToken Function

The `GenerateToken` function is used to generate a new access token. It reads the `ONELOGIN_CLIENT_ID` and `ONELOGIN_CLIENT_SECRET` environment variables, creates an authentication request, sends it, and handles the response. The newly generated access token is stored in the `Authenticator` instance.

```go
func (a *Authenticator) GenerateToken() error {
    // implementation details
}
```

## RevokeToken Function

The `RevokeToken` function is used to revoke an existing access token. It reads the `ONELOGIN_CLIENT_ID` and `ONELOGIN_CLIENT_SECRET` environment variables, creates a revocation request, sends it, and handles the response. If the revocation is successful, a confirmation message is printed.

```go
func (a *Authenticator) RevokeToken(token, domain *string) error {
    // implementation details
}
```

## GetToken Function

The `GetToken` function is used to retrieve the current access token from the `Authenticator` instance. It returns the `accessToken` field of the `Authenticator` struct.

```go
func (a *Authenticator) GetToken() (string, error) {
    return a.accessToken, nil
}
```

