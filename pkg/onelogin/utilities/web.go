package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	olerror "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/error"
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// receive http response, check error code status, if good return json of resp.Body
// else return error
func CheckHTTPResponse(resp *http.Response) (any, error) {
	// Handle 204 No Content responses - this is a success but with no content
	if resp.StatusCode == http.StatusNoContent {
		return map[string]any{"status": "success"}, nil
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Close the response body
	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return unmarshalBody(body)
}

// unmarshalBody decodes a response body the way this package always has: an
// array into []any, an object into map[string]any, anything else left as the
// raw string.
func unmarshalBody(body []byte) (any, error) {
	var data any
	bodyStr := string(body)
	if strings.HasPrefix(bodyStr, "[") {
		var slice []any
		if err := json.Unmarshal(body, &slice); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body into []any: %w", err)
		}
		data = slice
	} else if strings.HasPrefix(bodyStr, "{") {
		var dict map[string]any
		if err := json.Unmarshal(body, &dict); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body into map[string]any: %w", err)
		}
		data = dict
	} else {
		data = bodyStr
	}
	return data, nil
}

// CheckHTTPResponseWithErrorBody is CheckHTTPResponse with the API's own error
// message kept. CheckHTTPResponse returns the moment it sees a non-2xx status
// and never reads the body, so a caller is told
//
//	request failed with status: 422
//
// and nothing about what was actually wrong. On an endpoint with dozens of
// writable attributes that is not something a user can act on -- the body is
// what names the offending field.
//
// The "status: %d" wording is deliberate and load-bearing: callers match on it
// to recognise a 404 and treat the resource as gone. Anything that reformats
// this has to keep that substring.
//
// OneLogin reports errors as {"name":..., "message":..., "statusCode":...}. A
// body in any other shape falls back to the status alone, so this is never
// worse than CheckHTTPResponse.
func CheckHTTPResponseWithErrorBody(resp *http.Response) (any, error) {
	if resp.StatusCode == http.StatusNoContent {
		return map[string]any{"status": "success"}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if closeErr := resp.Body.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		if message := apiErrorMessage(body); message != "" {
			return nil, fmt.Errorf("request failed with status: %d: %s", resp.StatusCode, message)
		}
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return unmarshalBody(body)
}

// apiErrorMessage pulls the message out of a OneLogin error body, returning ""
// for a body that is not one.
func apiErrorMessage(body []byte) string {
	var apiError struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &apiError); err != nil {
		return ""
	}
	return apiError.Message
}

// CheckHTTPResponseWithPagination parses the response body and extracts pagination
// cursor headers (After-Cursor, Before-Cursor) from the V2 API response.
func CheckHTTPResponseWithPagination(resp *http.Response) (any, *mod.PaginationInfo, error) {
	data, err := CheckHTTPResponse(resp)
	if err != nil {
		return nil, nil, err
	}

	pagination := &mod.PaginationInfo{
		AfterCursor:  resp.Header.Get("After-Cursor"),
		BeforeCursor: resp.Header.Get("Before-Cursor"),
		Cursor:       resp.Header.Get("Cursor"),
	}
	if tp := resp.Header.Get("Total-Pages"); tp != "" {
		pagination.TotalPages, _ = strconv.Atoi(tp)
	}
	if cp := resp.Header.Get("Current-Page"); cp != "" {
		pagination.CurrentPage, _ = strconv.Atoi(cp)
	}
	if tc := resp.Header.Get("Total-Count"); tc != "" {
		pagination.TotalCount, _ = strconv.Atoi(tc)
	}

	return data, pagination, nil
}

// CheckHTTPResponseAndUnmarshal checks the HTTP response and unmarshals the response body into the target struct
func CheckHTTPResponseAndUnmarshal(resp *http.Response, target any) error {
	// Handle 204 No Content responses - this is a success but with no content
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Close the response body
	err = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}

	// Unmarshal the response body into the target struct
	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}

func BuildAPIPath(parts ...any) (string, error) {
	var path string
	for _, part := range parts {
		switch p := part.(type) {
		case string:
			path += "/" + p
		case int:
			path += fmt.Sprintf("/%d", p)
		case int32:
			path += fmt.Sprintf("/%d", p)
		case int64:
			path += fmt.Sprintf("/%d", p)
		default:
			// Handle other types if needed
			return path, olerror.NewSDKError("Unsupported path type")
		}
	}

	// Check if the path is valid
	if !IsPathValid(path) {
		return path, olerror.NewSDKError("Invalid path")
	}

	return path, nil
}

// AddQueryToPath adds the model as a JSON-encoded query parameter to the path and returns the new path.
func AddQueryToPath(path string, query any) (string, error) {
	if query == nil {
		return path, nil
	}

	// Convert query parameters to URL-encoded string
	values, err := queryToValues(query)
	if err != nil {
		return "", err
	}

	// Append query parameters to path
	if values.Encode() != "" {
		path += "?" + values.Encode()
	}

	return path, nil
}

func queryToValues(query any) (url.Values, error) {
	values := url.Values{}

	// Convert query parameters to URL-encoded string using reflection
	if query != nil {
		// First, get the json field names from struct tags
		queryBytes, err := json.Marshal(query)
		if err != nil {
			return nil, err
		}

		// Unmarshal to map[string]interface{} to handle all types of values
		var data map[string]any
		if err := json.Unmarshal(queryBytes, &data); err != nil {
			return nil, err
		}

		// Add each field to query parameters
		for key, value := range data {
			if value != nil {
				// Handle different value types
				switch v := value.(type) {
				case string:
					values.Set(key, v)
				case float64:
					values.Set(key, fmt.Sprintf("%v", v))
				case []any:
					// For arrays, convert to comma-separated string
					if len(v) > 0 {
						// Convert array to comma-separated string
						strItems := make([]string, len(v))
						for i, item := range v {
							strItems[i] = fmt.Sprintf("%v", item)
						}
						values.Set(key, strings.Join(strItems, ","))
					}
				default:
					// Convert other types to string
					values.Set(key, fmt.Sprintf("%v", v))
				}
			}
		}
	}

	return values, nil
}
