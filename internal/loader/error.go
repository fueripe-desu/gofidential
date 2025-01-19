package loader

import (
	"github.com/fueripe-desu/gofidential/errors"
	errorCode "github.com/fueripe-desu/gofidential/errors/loader"
)

// The [loader.newMissingNameError] function creates an error with the
// code "MissingName".
func newMissingNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MissingName,
		Message: "Environment name is required.",
		Hint:    "Did you set the 'Name' field in the 'Environment' struct passed to the Load() function?",
	}
}

// The [loader.newInvalidNameError] function creates an error with the
// code "InvalidName".
func newInvalidNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidName,
		Message: "Environment name must contain only lowercase letters and underscores.",
		Hint:    "Check if the 'Name' field set in the 'Environment' struct contains only lowercase letters and underscores.",
	}
}

// The [loader.newUnderscoreOnlyNameError] function creates an error with the
// code "UnderscoreOnlyName".
func newUnderscoreOnlyNameError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnderscoreOnlyName,
		Message: "Environment name must not be composed by underscores only.",
		Hint:    "Try using lowercase letters in the 'Name' field of the 'Environment' struct passed to the Load() function.",
	}
}

// The [loader.newTrailingUnderscoreError] function creates an error with the
// code "TrailingUnderscore".
func newTrailingUnderscoreError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.TrailingUnderscore,
		Message: "Environment name must not have trailing underscores.",
		Hint:    "Remove any trailing underscores from the 'Name' field in the 'Environment' struct you passed to the Load() function.",
	}
}

// The [loader.newLeadingUnderscoreError] function creates an error with the
// code "LeadingUnderscore".
func newLeadingUnderscoreError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.LeadingUnderscore,
		Message: "Environment name must not have leading underscores.",
		Hint:    "Remove any leading underscores from the 'Name' field in the 'Environment' struct you passed to the Load() function.",
	}
}

// The [loader.newInexistentDirError] function creates an error with the
// code "InexistentDir".
func newInexistentDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InexistentDir,
		Message: "Specified env directory path does not exist.",
		Hint:    "Have you set the 'OverridePath' field of the 'Environment' struct to an existing path?",
	}
}

// The [loader.newFailedToReadDirError] function creates an error with the
// code "FailedToReadDir".
func newFailedToReadDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.FailedToReadDir,
		Message: "Failed to read specified directory.",
		Hint:    "Does the application have permission to read the path set in the 'OverridePath' field of the 'Environment' struct?",
	}
}

// The [loader.newFailedToReadEnvError] function creates an error with the
// code "FailedToReadEnv".
func newFailedToReadEnvError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.FailedToReadEnv,
		Message: "Failed to read .env file.",
		Hint:    "Does the application have the necessary permissions to read the target .env file?",
	}
}

// The [loader.newPathIsNotDirError] function creates an error with the
// code "PathIsNotDir".
func newPathIsNotDirError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.PathIsNotDir,
		Message: "The specified path does not point to a directory.",
		Hint:    "Check if the 'OverridePath' field in the 'Environment' struct is set to a directory, not a file.",
	}
}

// The [loader.newPathIsNotFileError] function creates an error with the
// code "PathIsNotFile".
func newPathIsNotFileError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.PathIsNotFile,
		Message: "Path does not point to a file.",
		Hint:    "This is an internal error. If you can’t find a solution it, please open an issue at https://github.com/fueripe-desu/gofidential.",
	}
}

// The [loader.newEnvNotFoundError] function creates an error with the
// code "EnvNotFound".
func newEnvNotFoundError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.EnvNotFound,
		Message: "Could not find target .env file.",
		Hint:    "Ensure that you have created a .env file with the same name as the 'Name' field in the 'Environment' struct that you passed to the Load() function.",
	}
}
