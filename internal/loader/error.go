package loader

import (
	"runtime/debug"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
)

const moduleIssuer string = "GoFidential/File loader"

func newMissingEnvNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.MissingEnvNameCode,
		Message:    "Name is required.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are correctly specifying the 'name' parameter of the 'Environment' provided to the Load() function.",
		Details:    map[string]string{},
	}
}

func newEnvDirNotExistError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.EnvDirNotExistCode,
		Message:    "Dir path does not exist.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are correctly specifying an existing dir path to the 'path' parameter of the 'Environment' provided to the Load() function.",
		Details:    map[string]string{},
	}
}

func newFailedToReadDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.FailedToReadDirCode,
		Message:    "Failed to read dir.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the application has permission to read the specified dir path.",
		Details:    map[string]string{},
	}
}

func newFailedToReadFileError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.FailedToReadFileCode,
		Message:    "Failed to read file.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the application has permission to read the specified file path.",
		Details:    map[string]string{},
	}
}

func newPathIsNotDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.PathIsNotDirCode,
		Message:    "Path does not point to a dir.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the 'path' of the 'Environment' provided to the Load() function points to a dir not a file.",
		Details:    map[string]string{},
	}
}

func newPathIsNotFileError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.PathIsNotFileCode,
		Message:    "Path does not point to a file.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "This is an internal error. Please, if you can not find a solution, open an issue at \"https://github.com/fueripe-desu/gofidential\".",
		Details:    map[string]string{},
	}
}

func newRootNotFoundError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.RootNotFoundCode,
		Message:    "Could not find the root dir.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Assert you are executing the project in the root dir or its sub directories.",
		Details:    map[string]string{},
	}
}

func newEnvNotFoundError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.EnvNotFoundCode,
		Message:    "Could not find target .env file.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Assert you are executing the project in the root dir or its sub directories.",
		Details:    map[string]string{},
	}
}
