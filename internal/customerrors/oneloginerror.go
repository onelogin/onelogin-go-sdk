package customerrors

import "fmt"

// OneloginError used for any errors.
type OneloginError struct {
	Context string
	Err     error
}

// OneloginErrorWrapper creates a new OneloginError and returns, if an error is passed in,
// the pointer to the error struct.
func OneloginErrorWrapper(context string, err error) error {
	if err == nil {
		return nil
	}

	return &OneloginError{
		Context: context,
		Err:     err,
	}
}

func (err *OneloginError) Error() string {
	errMsg := ""
	if err.Err != nil {
		errMsg = err.Err.Error()
	}
	return fmt.Sprintf("error: context: [%s], error_message: [%s]", err.Context, errMsg)
}
