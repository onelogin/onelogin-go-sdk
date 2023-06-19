package tests

import (
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/error"
)

func TestValidationError(t *testing.T) {
	expectedMessage := "mock message"
	expectedCode := "mock code"
	validationError := error.ValidationError{
		Message: expectedMessage,
		Code:    expectedCode,
	}

	t.Run("when Message is accessed", func(t *testing.T) {
		message := validationError.Message

		if message != expectedMessage {
			t.Errorf("unexpected message: got %q, want %q", message, expectedMessage)
		}
	})

	t.Run("when Code is accessed", func(t *testing.T) {
		code := validationError.Code

		if code != expectedCode {
			t.Errorf("unexpected code: got %q, want %q", code, expectedCode)
		}
	})
}

func TestNewSerializationError(t *testing.T) {
	expectedMessage := "mock message"
	err := error.NewSerializationError(expectedMessage)

	serializationError, ok := err.(error.SerializationError)
	if !ok {
		t.Errorf("unexpected error type: got %T, want *error.SerializationError", err)
		return
	}

	if serializationError.Message != expectedMessage {
		t.Errorf("unexpected message: got %q, want %q", serializationError.Message, expectedMessage)
	}
}

func TestNewSDKError(t *testing.T) {
	expectedMessage := "mock message"
	err := error.NewSDKError(expectedMessage)

	sdkError, ok := err.(error.SDKError)
	if !ok {
		t.Errorf("unexpected error type: got %T, want *error.SDKError", err)
		return
	}

	if sdkError.Message != expectedMessage {
		t.Errorf("unexpected message: got %q, want %q", sdkError.Message, expectedMessage)
	}
}

func TestSerializationError_Is(t *testing.T) {
	expectedMessage := "mock message"
	serializationError := error.NewSerializationError(expectedMessage)

	t.Run("when the error is a SerializationError", func(t *testing.T) {
		var targetErr error.SerializationError
		if !errors.As(serializationError, &targetErr) {
			t.Errorf("expected error to be a SerializationError")
			return
		}
		if targetErr.Message != expectedMessage {
			t.Errorf("unexpected message: got %q, want %q", targetErr.Message, expectedMessage)
		}
	})

	t.Run("when the error is not a SerializationError", func(t *testing.T) {
		var targetErr error.SerializationError
		if !errors.As(serializationError, &targetErr) {
			t.Errorf("expected error to be a SerializationError")
			return
		}
	})
}

func TestSDKError_Is(t *testing.T) {
	expectedMessage := "mock message"
	sdkError := error.NewSDKError(expectedMessage)

	t.Run("when the error is an SDKError", func(t *testing.T) {
		targetErr := &error.SDKError{}
		if !errors.As(sdkError, targetErr) {
			t.Errorf("expected error to be an SDKError")
			return
		}
		if targetErr.Message != expectedMessage {
			t.Errorf("unexpected message: got %q, want %q", targetErr.Message, expectedMessage)
		}
	})

	t.Run("when the error is not an SDKError", func(t *testing.T) {
		var targetErr error.SDKError
		if !errors.As(sdkError, &targetErr) {
			t.Errorf("expected error to be an SDKError")
			return
		}
	})
}
