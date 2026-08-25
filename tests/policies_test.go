package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	ol "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

// BuildAPIPath refuses anything not matched by the validPaths whitelist, so a
// policy method with no entry there fails on "Invalid path" rather than on
// anything to do with the request. These are the three shapes the policy
// endpoints use.
func TestPolicyPathsAreValid(t *testing.T) {
	for _, tt := range []struct {
		name  string
		parts []any
		want  string
	}{
		{"collection", []any{ol.PoliciesPath}, "/api/2/policies"},
		{"member", []any{ol.PoliciesPath, 955633}, "/api/2/policies/955633"},
		{"set_default", []any{ol.PoliciesPath, 955633, "set_default"}, "/api/2/policies/955633/set_default"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utl.BuildAPIPath(tt.parts...)
			if err != nil {
				t.Fatalf("BuildAPIPath(%v): %v", tt.parts, err)
			}
			if got != tt.want {
				t.Fatalf("expected %s, got %s", tt.want, got)
			}
		})
	}
}

func response(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestCheckHTTPResponseWithErrorBody(t *testing.T) {
	// The whole point: the message names the field, and a bare status does
	// not, on an endpoint with seventy of them.
	t.Run("surfaces the API's message", func(t *testing.T) {
		_, err := utl.CheckHTTPResponseWithErrorBody(response(422,
			`{"name":"UnprocessableEntityError","message":"password_expiration_days is not applicable to app policies","statusCode":422}`))
		if err == nil {
			t.Fatal("expected an error")
		}
		if !strings.Contains(err.Error(), "password_expiration_days") {
			t.Fatalf("expected the field to be named, got %q", err)
		}
	})

	// Callers recognise a missing resource by this substring and drop it from
	// state. Reformatting the error without it turns a deleted policy into a
	// hard failure on every plan.
	t.Run("keeps the status readable for 404 detection", func(t *testing.T) {
		_, err := utl.CheckHTTPResponseWithErrorBody(response(404,
			`{"name":"NotFoundError","message":"The resource with the given id could not be found","statusCode":404}`))
		if err == nil {
			t.Fatal("expected an error")
		}
		if !strings.Contains(err.Error(), "status: 404") {
			t.Fatalf("expected %q to contain \"status: 404\"", err)
		}
	})

	t.Run("falls back to the status for a body it cannot read", func(t *testing.T) {
		_, err := utl.CheckHTTPResponseWithErrorBody(response(502, "<html>bad gateway</html>"))
		if err == nil {
			t.Fatal("expected an error")
		}
		if !strings.Contains(err.Error(), "status: 502") {
			t.Fatalf("expected the status, got %q", err)
		}
	})

	t.Run("returns the decoded body on success", func(t *testing.T) {
		got, err := utl.CheckHTTPResponseWithErrorBody(response(200, `{"id":955633,"name":"MFA Required"}`))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		policy, ok := got.(map[string]any)
		if !ok {
			t.Fatalf("expected a map, got %T", got)
		}
		if policy["name"] != "MFA Required" {
			t.Fatalf("unexpected body: %v", policy)
		}
	})

	t.Run("treats 204 as success", func(t *testing.T) {
		// DELETE answers 204 with no body; decoding that as JSON would fail.
		got, err := utl.CheckHTTPResponseWithErrorBody(response(204, ""))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatal("expected a success value")
		}
	})

	t.Run("accepts a 2xx the older check rejected", func(t *testing.T) {
		// CheckHTTPResponse allows only 200, 201 and 202, so the status here
		// has to be outside that set for this to test anything.
		if _, err := utl.CheckHTTPResponseWithErrorBody(response(203, `{"ok":true}`)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	// A body that is never closed keeps its connection out of the pool, and
	// the paths that skip it are the easy ones to miss: the early return for
	// 204 and the early return for an error status.
	t.Run("closes the body on every path", func(t *testing.T) {
		for _, tt := range []struct {
			name   string
			status int
			body   string
		}{
			{"no content", 204, ""},
			{"success", 200, `{"id":955633}`},
			{"api error", 422, `{"message":"nope","statusCode":422}`},
			{"unreadable error", 502, "<html>bad gateway</html>"},
		} {
			t.Run(tt.name, func(t *testing.T) {
				body := &trackingBody{Reader: strings.NewReader(tt.body)}
				_, _ = utl.CheckHTTPResponseWithErrorBody(&http.Response{
					StatusCode: tt.status,
					Body:       body,
				})
				if !body.closed {
					t.Fatalf("status %d: response body was left open", tt.status)
				}
			})
		}
	})
}

// trackingBody records whether it was closed.
type trackingBody struct {
	io.Reader
	closed bool
}

func (b *trackingBody) Close() error {
	b.closed = true
	return nil
}

// The same leak existed in CheckHTTPResponse, on both of its early returns.
// Every DELETE in this SDK answers 204, so that path is well travelled.
func TestCheckHTTPResponseClosesBody(t *testing.T) {
	for _, tt := range []struct {
		name   string
		status int
		body   string
	}{
		{"no content", 204, ""},
		{"error status", 422, `{"message":"nope"}`},
		{"success", 200, `{"id":955633}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body := &trackingBody{Reader: strings.NewReader(tt.body)}
			_, _ = utl.CheckHTTPResponse(&http.Response{
				StatusCode: tt.status,
				Body:       body,
			})
			if !body.closed {
				t.Fatalf("status %d: response body was left open", tt.status)
			}
		})
	}
}
