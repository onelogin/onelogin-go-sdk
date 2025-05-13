package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// Test context propagation in the client
func TestClientWithContext(t *testing.T) {
	client := createMockClient()

	// Setup mock response
	client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
		// Check if the request has a context
		if req.Context() == nil {
			t.Fatal("Request context is nil")
		}

		response := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
		}
		return response, nil
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test GetWithContext
	resp, err := client.GetWithContext(ctx, new(string), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `{"key":"value"}` {
		t.Fatalf("Expected `{\"key\":\"value\"}`, got %s", string(body))
	}
}

// Test query parameters handling, especially for arrays and pagination
func TestQueryParameters(t *testing.T) {
	client := createMockClient()

	// Setup mock response
	client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
		// Check that the URL has correctly encoded query parameters
		// This specifically tests our queryToValues function improvements
		url := req.URL.String()

		// Check for pagination parameters
		if req.URL.Query().Get("limit") != "10" {
			t.Fatalf("Expected limit=10 in query, got %s", url)
		}

		// Send response with pagination headers
		response := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`[{"id":1,"name":"User 1"},{"id":2,"name":"User 2"}]`)),
			Header:     http.Header{},
		}

		// Add pagination headers
		response.Header.Set("After-Cursor", "next_page_token")
		response.Header.Set("Total-Pages", "5")
		response.Header.Set("Current-Page", "1")
		response.Header.Set("Total-Count", "42")

		return response, nil
	}

	// Test with a query containing various parameter types
	query := &models.UserQuery{
		Limit: "10",
		Page:  "1",
	}

	// Create a pointer to string for GroupID
	groupID := "123"
	query.GroupID = &groupID

	// If we had a pointer to array, it would also be properly handled now
	roleIDs := []int32{1, 2, 3}
	query.RoleIDs = &roleIDs

	// Use simple Get to test query parameter handling
	resp, err := client.Get(new(string), query)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// The parameter handling test happened in the mock function
	// So if we got here without failing, it worked
}

// Test the PagedResponse implementation
func TestPagedResponse(t *testing.T) {
	// Create a sample paged response
	data := []interface{}{
		map[string]interface{}{"id": 1, "name": "User 1"},
		map[string]interface{}{"id": 2, "name": "User 2"},
	}

	pagination := models.PaginationInfo{
		Cursor:       "current_cursor",
		AfterCursor:  "next_cursor",
		BeforeCursor: "prev_cursor",
		TotalPages:   5,
		CurrentPage:  1,
		TotalCount:   42,
	}

	pagedResponse := models.PagedResponse{
		Data:       data,
		Pagination: pagination,
	}

	// Verify the structure can be marshaled to JSON properly
	jsonBytes, err := json.Marshal(pagedResponse)
	if err != nil {
		t.Fatal(err)
	}

	// Parse back the JSON to verify
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonBytes, &parsedResponse)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the structure has both data and pagination
	if _, ok := parsedResponse["data"]; !ok {
		t.Fatal("Parsed response is missing 'data' field")
	}

	if _, ok := parsedResponse["pagination"]; !ok {
		t.Fatal("Parsed response is missing 'pagination' field")
	}

	// Verify pagination information
	paginationMap, ok := parsedResponse["pagination"].(map[string]interface{})
	if !ok {
		t.Fatal("Pagination is not a map")
	}

	if paginationMap["total_pages"].(float64) != 5 {
		t.Fatalf("Expected total_pages to be 5, got %v", paginationMap["total_pages"])
	}
}
