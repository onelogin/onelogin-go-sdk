package tests

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// TestAppWritesKeepAPIErrorMessage covers CreateApp and UpdateApp returning the
// API's own message for a rejected write. The body below is what the app
// endpoint answers for a configuration key the connector does not have.
func TestAppWritesKeepAPIErrorMessage(t *testing.T) {
	const rejected = `{"name":"BadRequestError","message":"Unknown parameter on configuration section: not_a_real_field","statusCode":400}`

	answer := func(status int, body string) *onelogin.OneloginSDK {
		client := createMockClient()
		client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: status, Body: io.NopCloser(bytes.NewBufferString(body))}, nil
		}
		return &onelogin.OneloginSDK{Client: client}
	}

	name := "SAML app"
	connectorID := int32(45504)
	app := mod.App{Name: &name, ConnectorID: &connectorID}

	writes := map[string]func(*onelogin.OneloginSDK) (interface{}, error){
		"CreateApp": func(sdk *onelogin.OneloginSDK) (interface{}, error) { return sdk.CreateApp(app) },
		"UpdateApp": func(sdk *onelogin.OneloginSDK) (interface{}, error) { return sdk.UpdateApp(1504547, app) },
	}

	for name, write := range writes {
		t.Run(name+" names the refused key", func(t *testing.T) {
			_, err := write(answer(400, rejected))
			if err == nil {
				t.Fatal("expected an error")
			}
			want := "request failed with status: 400: Unknown parameter on configuration section: not_a_real_field"
			if err.Error() != want {
				t.Fatalf("expected %q, got %q", want, err)
			}
		})

		// Callers recognise a missing app by this substring, so it has to
		// survive the message being appended.
		t.Run(name+" still reports a 404 as status: 404", func(t *testing.T) {
			_, err := write(answer(404, `{"name":"NotFoundError","message":"Not Found","statusCode":404}`))
			if err == nil || !strings.Contains(err.Error(), "status: 404") {
				t.Fatalf("expected an error containing \"status: 404\", got %v", err)
			}
		})

		t.Run(name+" returns the body on success", func(t *testing.T) {
			got, err := write(answer(200, `{"id":1504547}`))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			body, ok := got.(map[string]interface{})
			if !ok || body["id"] != float64(1504547) {
				t.Fatalf("expected the decoded body, got %#v", got)
			}
		})
	}
}
