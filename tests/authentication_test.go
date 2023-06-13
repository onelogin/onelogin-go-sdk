package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	olError "github.com/onelogin/onelogin-go-sdk/internal/error"
)

func TestNewAuthenticator(t *testing.T) {
	auth := authentication.NewAuthenticator()
	if auth == nil {
		t.Error("NewAuthenticator should not return nil")
	}
}

func TestGenerateToken(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Set the required environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test_client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test_client_secret")

	// Create a test server to mock the authentication endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != authentication.TokenPath {
			t.Errorf("Expected path %s, got %s", authentication.TokenPath, r.URL.Path)
		}
		if r.Header.Get("Authorization") == "" {
			t.Error("Authorization header is missing")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "test_access_token"}`))
	}))
	defer testServer.Close()

	// Set the authentication URL to the test server URL
	authentication.TokenPath = testServer.URL + authentication.TokenPath

	// Create a new authenticator and generate a token
	auth := authentication.NewAuthenticator()
	err := auth.GenerateToken()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	tk, err := auth.GetToken()
	// Check that the access token was stored
	if tk != "test_access_token" {
		t.Errorf("Expected access token test_access_token, got %s", tk)
	}
}

func TestGenerateTokenMissingEnvVars(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Unset the required environment variables
	os.Unsetenv("ONELOGIN_CLIENT_ID")
	os.Unsetenv("ONELOGIN_CLIENT_SECRET")

	// Create a new authenticator and generate a token
	auth := authentication.NewAuthenticator()
	err := auth.GenerateToken()

	// Check that the error is of the correct type
	var authErr *olError.AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected error of type AuthenticationError, got %T", err)
	}
}

func TestGenerateTokenAuthenticationError(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Set the required environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test_client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test_client_secret")

	// Create a test server to mock the authentication endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer testServer.Close()

	// Set the authentication URL to the test server URL
	authentication.TokenPath = testServer.URL + authentication.TokenPath

	// Create a new authenticator and generate a token
	auth := authentication.NewAuthenticator()
	err := auth.GenerateToken()

	// Check that the error is of the correct type
	var authErr *olError.AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected error of type AuthenticationError, got %T", err)
	}
}

func TestRevokeTokenMissingEnvVars(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Unset the required environment variables
	os.Unsetenv("ONELOGIN_CLIENT_ID")
	os.Unsetenv("ONELOGIN_CLIENT_SECRET")

	// Create a new authenticator and try to revoke a token
	auth := authentication.NewAuthenticator()
	err := auth.RevokeToken(nil, nil)

	// Check that the error is of the correct type
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestRevokeTokenRequestError(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Set the required environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test_client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test_client_secret")

	// Create a test server that always returns an error
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	// Set the revocation URL to the test server URL
	authentication.RevokePath = testServer.URL + authentication.RevokePath

	// Create a new authenticator and try to revoke a token
	auth := authentication.NewAuthenticator()
	err := auth.RevokeToken(nil, nil)

	// Check that the error is of the correct type
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestRevokeTokenAuthenticationError(t *testing.T) {
	// Save the current environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	defer func() {
		os.Setenv("ONELOGIN_CLIENT_ID", clientID)
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}()

	// Set the required environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test_client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test_client_secret")

	// Create a test server to mock the revocation endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer testServer.Close()

	// Set the revocation URL to the test server URL
	authentication.RevokePath = testServer.URL + authentication.RevokePath

	// Create a new authenticator and try to revoke a token
	auth := authentication.NewAuthenticator()
	err := auth.RevokeToken(nil, nil)

	// Check that the error is of the correct type
	var authErr *olError.AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected error of type AuthenticationError, got %T", err)
	}
}
