package reflector

import (
	"runtime/debug"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
	errorCode "github.com/fueripe-desu/gofidential/errors/reflector"
)

const moduleIssuer string = "GoFidential/Reflector"

func newDataIsNilError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.DataIsNilCode,
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
		Code:       errorCode.DataIsNotPtrCode,
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
		Code:       errorCode.DataIsNilPtrCode,
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
		Code:       errorCode.DataIsNilPtrCode,
		Message:    "The 'data' pointer must point to a valid struct.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are passing a pointer to a valid struct to the the Load() function.",
		Details:    map[string]string{},
	}
}

func newDuplicateKeyError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.DuplicateKeyCode,
		Message:    "Duplicate keys are not allowed.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "If you are using GoFidential v1, this error happened due to a bug. Please open an issue at github.com/fueripe-desu/gofidential",
		Details: map[string]string{
			"normalized_key": key,
		},
	}
}

func newMissingFieldError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.MissingFieldCode,
		Message:    "The expected '" + field + "' field is missing.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Add the missing field, or remove the unused key from the .env file.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newUnsupportedTypeError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.MissingFieldCode,
		Message:    "The '" + field + "' field has an unsupported type",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the type of the invalid field or remove it if unusued.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidIntError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidIntCode,
		Message:    "The '" + field + "' field is not a valid int.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid integer string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidUintError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidUintCode,
		Message:    "The '" + field + "' field is not a valid unsigned int.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid unsigned integer string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidFloatError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidFloatCode,
		Message:    "The '" + field + "' field is not a valid float",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid float string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidComplexError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidComplexCode,
		Message:    "The '" + field + "' field is not a valid complex number.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid complex number string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidBoolError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidBoolCode,
		Message:    "The '" + field + "' field is not a valid boolean.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid boolean string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newInvalidTimeError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidTimeCode,
		Message:    "The '" + field + "' field is not a valid time value.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Change the value so it becomes a valid time value string.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newUnsettableFieldError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidTimeCode,
		Message:    "The '" + field + "' field is not settable.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are passing a valid pointer to a struct.",
		Details: map[string]string{
			"field": field,
		},
	}
}

func newUnexportedFieldError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidTimeCode,
		Message:    "The provided struct contains unexported fields.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if all fields in the struct start with an uppercase letter.",
		Details:    map[string]string{},
	}
}

func newInvalidEnvDataError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidEnvDataCode,
		Message:    "Failed to read env data. Data is invalid.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "If this error happened, there is a bug in the code. Please open an issue at github.com/fueripe-desu/gofidential",
		Details:    map[string]string{},
	}
}
