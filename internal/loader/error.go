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

func newEnvFolderNotExistError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.EnvFolderNotExistCode,
		Message:    "Folder path does not exist.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if you are correctly specifying an existing folder path to the 'path' parameter of the 'Environment' provided to the Load() function.",
		Details:    map[string]string{},
	}
}

func newFailedToReadFolderError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.FailedToReadFolderCode,
		Message:    "Failed to read folder.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the application has permission to read the specified folder path.",
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

func newPathIsNotFolderError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Issuer:     moduleIssuer,
		Code:       errors.PathIsNotFolderCode,
		Message:    "Path does not point to a folder.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Check if the 'path' of the 'Environment' provided to the Load() function points to a folder not a file.",
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
		Message:    "Could not find the root folder.",
		Timestamp:  time.Now(),
		StackTrace: debug.Stack(),
		Suggestion: "Assert you are executing the project in the root folder or its subfolders.",
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
		Suggestion: "Assert you are executing the project in the root folder or its subfolders.",
		Details:    map[string]string{},
	}
}
