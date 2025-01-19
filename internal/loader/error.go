package loader

import (
	"runtime/debug"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
	errorCode "github.com/fueripe-desu/gofidential/errors/loader"
)

const moduleIssuer string = "GoFidential/File loader"

func newMissingNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.MissingName,
		Message:    "Name is required.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are correctly specifying the 'name' parameter of the 'Environment' provided to the Load() function.",
		Details:    map[string]string{},
	}
}

func newInvalidNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InvalidName,
		Message:    "Environment name is invalid.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the provided environment name contains only lowercase letters (a-z) and underscores (_).",
		Details:    map[string]string{},
	}
}

func newUnderscoreOnlyNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.UnderscoreOnlyName,
		Message:    "Environment name must not be composed by underscores only.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Add lowercases letters to the name so it does not contain only underscores.",
		Details:    map[string]string{},
	}
}

func newTrailingUnderscoreError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.TrailingUnderscore,
		Message:    "Environment name must not have trailing underscores.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove trailing underscores from the environment name.",
		Details:    map[string]string{},
	}
}

func newLeadingUnderscoreError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.LeadingUnderscore,
		Message:    "Environment name must not have leading underscores.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Remove leading underscores from the environment name.",
		Details:    map[string]string{},
	}
}

func newInexistentDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.InexistentDir,
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
		Code:       errorCode.FailedToReadDir,
		Message:    "Failed to read dir.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the application has permission to read the specified dir path.",
		Details:    map[string]string{},
	}
}

func newFailedToReadEnvError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.FailedToReadEnv,
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
		Code:       errorCode.PathIsNotDir,
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
		Code:       errorCode.PathIsNotFile,
		Message:    "Path does not point to a file.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "This is an internal error. Please, if you can not find a solution, open an issue at \"https://github.com/fueripe-desu/gofidential\".",
		Details:    map[string]string{},
	}
}

func newEnvNotFoundError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errorCode.EnvNotFound,
		Message:    "Could not find target .env file.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Assert you are executing the project in the root dir or its sub directories.",
		Details:    map[string]string{},
	}
}
