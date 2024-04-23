package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	olerror "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/error"
)

// receive http response, check error code status, if good return json of resp.Body
// else return error
func CheckHTTPResponse(resp *http.Response) (interface{}, error) {
	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Close the response body
	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	// Try to unmarshal the response body into a map[string]interface{} or []interface{}
	var data interface{}
	bodyStr := string(body)
	//log.Printf("Response body: %s\n", bodyStr)
	if strings.HasPrefix(bodyStr, "[") {
		var slice []interface{}
		err = json.Unmarshal(body, &slice)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body into []interface{}: %w", err)
		}
		data = slice
	} else if strings.HasPrefix(bodyStr, "{") {
		var dict map[string]interface{}
		err = json.Unmarshal(body, &dict)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body into map[string]interface{}: %w", err)
		}
		data = dict
	} else {
		data = bodyStr
	}

	//log.Printf("Response body unmarshaled successfully: %v\n", data)
	return data, nil
}

func BuildAPIPath(parts ...interface{}) (string, error) {
	var path string
	for _, part := range parts {
		switch p := part.(type) {
		case string:
			path += "/" + p
		case int:
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
func AddQueryToPath(path string, query interface{}) (string, error) {
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

func queryToValues(query interface{}) (url.Values, error) {
	values := url.Values{}

	// Convert query parameters to URL-encoded string
	if query != nil {
		queryBytes, err := json.Marshal(query)
		if err != nil {
			return nil, err
		}
		var data map[string]string
		if err := json.Unmarshal(queryBytes, &data); err != nil {
			return nil, err
		}
		for key, value := range data {
			values.Set(key, value)
		}
	}

	return values, nil
}
