package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	olerror "github.com/onelogin/onelogin-go-sdk/internal/error"
	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
)

// receive http response, check error code status, if good return json of resp.Body
// else return error
func CheckHTTPResponse(resp *http.Response) (interface{}, error) {
	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
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

	// Try to unmarshal the response body into a map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		// If unmarshaling fails, return the response body as a string
		return string(body), nil
	}

	return data, nil
}

func ModelToJson(model interface{}) (string, error) {
	jsonData, err := json.Marshal(model)
	if err != nil {
		return "", olerror.NewSDKError("Error marshalling model into json.")
	}

	return string(jsonData), nil
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
func AddQueryToPath(path *string, model *mod.Queryable) (string, error) {
	// Convert the model to a JSON string
	jsonBytes, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	// URL-encode the JSON string
	encodedModel := url.QueryEscape(string(jsonBytes))

	// Parse the original path
	u, err := url.Parse(*path)
	if err != nil {
		return "", err
	}

	// Add the encoded model to the path's query string
	q := u.Query()
	q.Set("model", encodedModel)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
