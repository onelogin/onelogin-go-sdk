package customerrors

import "errors"

// Errors in services
var (
	ErrValueMissing = errors.New("A required parameter was not given")
)
