package reflector

import (
	"runtime/debug"
	"time"

	errors "github.com/fueripe-desu/gofidential/errors"
)

const moduleIssuer string = "GoFidential/Reflector"

func newDataIsNilError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.DataIsNilCode,
		Message:    "The 'data' parameter must not be nil.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are actually passing a pointer to the Load() function and not 'nil'.",
		Details:    map[string]string{},
	}
}

func newDataIsNotPtrError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.DataIsNotPtrCode,
		Message:    "The 'data' parameter must be a pointer.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are actually passing a pointer to the Load() function.",
		Details:    map[string]string{},
	}
}

func newDataIsNilPtrError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.DataIsNilPtrCode,
		Message:    "The 'data' pointer must not be nil.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are passing a pointer to a valid struct to the the Load() function.",
		Details:    map[string]string{},
	}
}

func newDataIsNotStructError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.DataIsNilPtrCode,
		Message:    "The 'data' pointer must point to a valid struct.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are passing a pointer to a valid struct to the the Load() function.",
		Details:    map[string]string{},
	}
}
