package utilities

import (
	"reflect"
	"regexp"
	"strings"
)

// ValidateQueryParams validates the query parameters based on the provided validators.
func ValidateQueryParams(query interface{}, validators map[string]func(interface{}) bool) bool {
	queryValue := reflect.ValueOf(query)
	if queryValue.Kind() == reflect.Ptr {
		queryValue = queryValue.Elem()
	}
	queryType := queryValue.Type()

	for i := 0; i < queryValue.NumField(); i++ {
		fieldValue := queryValue.Field(i)
		fieldType := queryType.Field(i)

		// Skip non-pointer fields
		if fieldValue.Kind() != reflect.Ptr {
			continue
		}

		// Skip nil fields
		if fieldValue.IsNil() {
			continue
		}

		fieldName := strings.Split(fieldType.Tag.Get("json"), ",")[0]

		validator, exists := validators[fieldName]
		if exists {
			if !validator(fieldValue.Interface()) {
				return false
			}
		}
	}

	return true
}

// Check if the constructed path matches any of the allowed path patterns
func IsPathValid(path string) bool {
	for _, pattern := range validPaths {
		match, _ := regexp.MatchString(pattern, path)
		if match {
			return true
		}
	}
	return false
}
