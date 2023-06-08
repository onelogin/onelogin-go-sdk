package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
)

func TestGenerateToken(t *testing.T) {
	// Setup environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test-id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test-secret")

	// Create an HTTP test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"access_token": "test-token",
		}

		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Update the domain in the Authenticator
	authenticator := authentication.NewAuthenticator()

	// Generate a token
	token, err := authenticator.GetToken()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if token != "test-token" {
		t.Errorf("Expected token 'test-token', but got '%v'", token)
	}
}

func TestRevokeToken(t *testing.T) {
	// Setup environment variables
	os.Setenv("ONELOGIN_CLIENT_ID", "test-id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test-secret")

	// Create an HTTP test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Update the domain in the Authenticator
	authenticator := authentication.NewAuthenticator()

	// Revoke a token
	domain := server.URL
	token := "test-token"
	err := authenticator.RevokeToken(&token, &domain)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
