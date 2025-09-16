package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// TestGetAppUsersWithPagination tests the new paginated GetAppUsers functionality
func TestGetAppUsersWithPagination(t *testing.T) {
	client := createMockClient()
	sdk := &onelogin.OneloginSDK{Client: client}

	t.Run("GetAppUsers without query params (backward compatibility)", func(t *testing.T) {
		expectedUsers := []models.UserApp{
			{ID: int32Ptr(1), LoginID: int32Ptr(100)},
			{ID: int32Ptr(2), LoginID: int32Ptr(200)},
		}

		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			// Verify the URL doesn't contain query parameters
			if req.URL.RawQuery != "" {
				t.Errorf("Expected no query parameters, got: %s", req.URL.RawQuery)
			}

			usersJSON, _ := json.Marshal(expectedUsers)
			resp := &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(usersJSON)),
				Header:     make(http.Header),
			}
			return resp, nil
		}

		result, err := sdk.GetAppUsers(123)
		if err != nil {
			t.Fatalf("GetAppUsers failed: %v", err)
		}

		// The result should be the same as before
		if result == nil {
			t.Fatal("Expected result to not be nil")
		}
	})

	t.Run("GetAppUsersWithQuery with pagination parameters", func(t *testing.T) {
		expectedUsers := []models.UserApp{
			{ID: int32Ptr(3), LoginID: int32Ptr(300)},
			{ID: int32Ptr(4), LoginID: int32Ptr(400)},
		}

		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			// Verify query parameters are included
			query := req.URL.Query()
			if query.Get("limit") != "50" {
				t.Errorf("Expected limit=50, got: %s", query.Get("limit"))
			}
			if query.Get("page") != "2" {
				t.Errorf("Expected page=2, got: %s", query.Get("page"))
			}

			usersJSON, _ := json.Marshal(expectedUsers)
			resp := &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(usersJSON)),
				Header:     make(http.Header),
			}
			return resp, nil
		}

		query := &models.AppUserQuery{
			Limit: "50",
			Page:  "2",
		}

		result, err := sdk.GetAppUsersWithQuery(123, query)
		if err != nil {
			t.Fatalf("GetAppUsersWithQuery failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected result to not be nil")
		}
	})

	t.Run("GetAppUsersWithPagination with full pagination info", func(t *testing.T) {
		expectedUsers := []models.UserApp{
			{ID: int32Ptr(5), LoginID: int32Ptr(500)},
			{ID: int32Ptr(6), LoginID: int32Ptr(600)},
		}

		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			usersJSON, _ := json.Marshal(expectedUsers)
			resp := &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(usersJSON)),
				Header:     make(http.Header),
			}
			
			// Add pagination headers
			resp.Header.Set("Current-Page", "2")
			resp.Header.Set("Total-Pages", "5")
			resp.Header.Set("Total-Count", "245")
			resp.Header.Set("After-Cursor", "cursor_next_page")
			resp.Header.Set("Before-Cursor", "cursor_prev_page")
			
			return resp, nil
		}

		query := &models.AppUserQuery{
			Limit: "50",
			Page:  "2",
		}

		result, err := sdk.GetAppUsersWithPagination(123, query)
		if err != nil {
			t.Fatalf("GetAppUsersWithPagination failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected result to not be nil")
		}

		// Verify pagination info
		if result.Pagination.CurrentPage != 2 {
			t.Errorf("Expected CurrentPage=2, got: %d", result.Pagination.CurrentPage)
		}
		if result.Pagination.TotalPages != 5 {
			t.Errorf("Expected TotalPages=5, got: %d", result.Pagination.TotalPages)
		}
		if result.Pagination.TotalCount != 245 {
			t.Errorf("Expected TotalCount=245, got: %d", result.Pagination.TotalCount)
		}
		if result.Pagination.AfterCursor != "cursor_next_page" {
			t.Errorf("Expected AfterCursor=cursor_next_page, got: %s", result.Pagination.AfterCursor)
		}
		if result.Pagination.BeforeCursor != "cursor_prev_page" {
			t.Errorf("Expected BeforeCursor=cursor_prev_page, got: %s", result.Pagination.BeforeCursor)
		}
	})

	t.Run("GetAppUsersWithPagination with cursor", func(t *testing.T) {
		expectedUsers := []models.UserApp{
			{ID: int32Ptr(7), LoginID: int32Ptr(700)},
		}

		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			// Verify cursor parameter is included
			query := req.URL.Query()
			if query.Get("cursor") != "abc123" {
				t.Errorf("Expected cursor=abc123, got: %s", query.Get("cursor"))
			}

			usersJSON, _ := json.Marshal(expectedUsers)
			resp := &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(usersJSON)),
				Header:     make(http.Header),
			}
			return resp, nil
		}

		query := &models.AppUserQuery{
			Cursor: "abc123",
			Limit:  "25",
		}

		result, err := sdk.GetAppUsersWithPagination(123, query)
		if err != nil {
			t.Fatalf("GetAppUsersWithPagination with cursor failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected result to not be nil")
		}
	})

	t.Run("AppUserQuery validation", func(t *testing.T) {
		query := &models.AppUserQuery{
			Limit: "invalid",
		}

		validators := query.GetKeyValidators()
		if validators == nil {
			t.Fatal("Expected validators to not be nil")
		}

		// Test limit validator exists
		if _, exists := validators["limit"]; !exists {
			t.Error("Expected limit validator to exist")
		}

		// Test page validator exists  
		if _, exists := validators["page"]; !exists {
			t.Error("Expected page validator to exist")
		}

		// Test cursor validator exists
		if _, exists := validators["cursor"]; !exists {
			t.Error("Expected cursor validator to exist")
		}
	})
}

// Helper function to create int32 pointers
func int32Ptr(i int32) *int32 {
	return &i
}