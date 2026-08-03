package tests

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

// roleMemberQuery is a minimal Queryable for exercising the paginated role
// sub-endpoint getters. The V2 API treats cursor and limit/page as mutually
// exclusive, so callers clear limit/page once a cursor is in hand.
type roleMemberQuery struct {
	Limit  string `json:"limit,omitempty"`
	Page   string `json:"page,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

func (q *roleMemberQuery) GetKeyValidators() map[string]func(interface{}) bool {
	isString := func(v interface{}) bool {
		_, ok := v.(string)
		return ok
	}
	return map[string]func(interface{}) bool{
		"limit":  isString,
		"page":   isString,
		"cursor": isString,
	}
}

// respondWith builds a 200 response carrying body and the supplied After-Cursor.
func respondWith(body string, afterCursor string) *http.Response {
	header := make(http.Header)
	if afterCursor != "" {
		header.Set("After-Cursor", afterCursor)
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     header,
	}
}

// TestRoleMemberGettersSurfacePagination covers the three role sub-endpoints
// (apps, users, admins), which are paginated by the V2 API. Each getter must
// pass query parameters through to the request and surface the After-Cursor
// response header so callers can walk every page.
func TestRoleMemberGettersSurfacePagination(t *testing.T) {
	t.Run("apps getter passes query params and returns After-Cursor", func(t *testing.T) {
		client := createMockClient()
		sdk := &onelogin.OneloginSDK{Client: client}

		var gotPath, gotQuery string
		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			gotPath = req.URL.Path
			gotQuery = req.URL.RawQuery
			return respondWith(`[{"id":1,"name":"app-one"}]`, "cursor-page-2"), nil
		}

		result, pagination, err := sdk.GetRoleAppsWithPaginationAndContext(
			context.Background(), 42, &roleMemberQuery{Limit: "100"})
		if err != nil {
			t.Fatalf("GetRoleAppsWithPaginationAndContext failed: %v", err)
		}
		if gotPath != "/api/2/roles/42/apps" {
			t.Errorf("expected path /api/2/roles/42/apps, got %s", gotPath)
		}
		if gotQuery != "limit=100" {
			t.Errorf("expected query limit=100, got %q", gotQuery)
		}
		if pagination == nil || pagination.AfterCursor != "cursor-page-2" {
			t.Fatalf("expected After-Cursor to be surfaced, got %+v", pagination)
		}
		items, ok := result.([]interface{})
		if !ok || len(items) != 1 {
			t.Fatalf("expected one app in result, got %#v", result)
		}
	})

	t.Run("users getter passes query params and returns After-Cursor", func(t *testing.T) {
		client := createMockClient()
		sdk := &onelogin.OneloginSDK{Client: client}

		var gotPath, gotQuery string
		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			gotPath = req.URL.Path
			gotQuery = req.URL.RawQuery
			return respondWith(`[{"id":10,"username":"joe.user"}]`, "cursor-page-2"), nil
		}

		_, pagination, err := sdk.GetRoleUsersWithPaginationAndContext(
			context.Background(), 42, &roleMemberQuery{Cursor: "cursor-page-1"})
		if err != nil {
			t.Fatalf("GetRoleUsersWithPaginationAndContext failed: %v", err)
		}
		if gotPath != "/api/2/roles/42/users" {
			t.Errorf("expected path /api/2/roles/42/users, got %s", gotPath)
		}
		if gotQuery != "cursor=cursor-page-1" {
			t.Errorf("expected query cursor=cursor-page-1, got %q", gotQuery)
		}
		if pagination == nil || pagination.AfterCursor != "cursor-page-2" {
			t.Fatalf("expected After-Cursor to be surfaced, got %+v", pagination)
		}
	})

	t.Run("admins getter passes query params and returns After-Cursor", func(t *testing.T) {
		client := createMockClient()
		sdk := &onelogin.OneloginSDK{Client: client}

		var gotPath string
		client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
			gotPath = req.URL.Path
			return respondWith(`[{"id":7,"username":"admin.user"}]`, ""), nil
		}

		_, pagination, err := sdk.GetRoleAdminsWithPaginationAndContext(
			context.Background(), 42, nil)
		if err != nil {
			t.Fatalf("GetRoleAdminsWithPaginationAndContext failed: %v", err)
		}
		if gotPath != "/api/2/roles/42/admins" {
			t.Errorf("expected path /api/2/roles/42/admins, got %s", gotPath)
		}
		// An absent After-Cursor is how the API signals the last page.
		if pagination == nil || pagination.AfterCursor != "" {
			t.Fatalf("expected empty After-Cursor on last page, got %+v", pagination)
		}
	})
}

// TestRoleMemberGettersPropagateContext verifies the context reaches the
// outbound request, so a cancelled Terraform plan stops issuing calls.
func TestRoleMemberGettersPropagateContext(t *testing.T) {
	client := createMockClient()
	sdk := &onelogin.OneloginSDK{Client: client}

	type ctxKey string
	const key ctxKey = "trace"

	var sawValue interface{}
	client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
		sawValue = req.Context().Value(key)
		return respondWith(`[]`, ""), nil
	}

	ctx := context.WithValue(context.Background(), key, "present")
	if _, _, err := sdk.GetRoleUsersWithPaginationAndContext(ctx, 42, nil); err != nil {
		t.Fatalf("GetRoleUsersWithPaginationAndContext failed: %v", err)
	}
	if sawValue != "present" {
		t.Errorf("expected context to reach the request, got %v", sawValue)
	}
}

// TestRoleMemberGettersWalkAllPages exercises the cursor loop callers are
// expected to run: follow After-Cursor until it comes back empty.
func TestRoleMemberGettersWalkAllPages(t *testing.T) {
	client := createMockClient()
	sdk := &onelogin.OneloginSDK{Client: client}

	pages := []struct {
		body        string
		afterCursor string
	}{
		{`[{"id":1},{"id":2}]`, "cursor-2"},
		{`[{"id":3}]`, ""},
	}
	callCount := 0
	client.HttpClient.(*MockHttpClient).DoFunc = func(req *http.Request) (*http.Response, error) {
		page := pages[callCount]
		callCount++
		return respondWith(page.body, page.afterCursor), nil
	}

	var ids []int
	query := &roleMemberQuery{Limit: "100"}
	for {
		result, pagination, err := sdk.GetRoleUsersWithPaginationAndContext(
			context.Background(), 42, query)
		if err != nil {
			t.Fatalf("page %d failed: %v", callCount, err)
		}
		items, ok := result.([]interface{})
		if !ok {
			t.Fatalf("expected an array, got %#v", result)
		}
		for _, item := range items {
			if obj, ok := item.(map[string]interface{}); ok {
				if id, ok := obj["id"].(float64); ok {
					ids = append(ids, int(id))
				}
			}
		}
		if pagination == nil || pagination.AfterCursor == "" {
			break
		}
		// cursor and limit/page are mutually exclusive in the V2 API
		query.Cursor = pagination.AfterCursor
		query.Limit, query.Page = "", ""
	}

	if callCount != 2 {
		t.Errorf("expected 2 requests, got %d", callCount)
	}
	if len(ids) != 3 || ids[0] != 1 || ids[2] != 3 {
		t.Errorf("expected ids [1 2 3] across both pages, got %v", ids)
	}
}
