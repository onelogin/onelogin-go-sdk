package models

import (
	"time"
)

type Queryable interface {
	GetKeyValidators() map[string]func(interface{}) bool
}

// validateString checks if the provided value is a string.
func validateString(val interface{}) bool {
	_, ok := val.(string)
	return ok
}

// validateTime checks if the provided value is a time.Time.
func validateTime(val interface{}) bool {
	_, ok := val.(*time.Time)
	return ok
}

// validateInt checks if the provided value is an int.
func validateInt(val interface{}) bool {
	_, ok := val.(int)
	return ok
}

// validateBool checks if the provided value is a bool.
func validateBool(val interface{}) bool {
	_, ok := val.(bool)
	return ok
}
