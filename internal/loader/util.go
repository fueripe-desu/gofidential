package loader

import (
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
)

// The [loader.isLowerUnderscore] function checks if the input string consists
// entirely of lowercase letters (a-z) and underscores (_).
//
// Parameters:
//   - v (string): The string to be checked.
//
// Returns:
//   - bool: Returns true if the string contains only lowercase letters and underscores,
//     otherwise returns false.
//
// Notes:
//   - An empty string is considered as not containing lowercase letters and underscores.
func isLowerUnderscore(v string) bool {
	regex := regexp.MustCompile(`^[a-z_]+$`)
	return regex.MatchString(v)
}

// The [loader.isUnderscore] function checks if the input string is composed entirely
// of underscores.
//
// Parameters:
//   - v (string): The string to be checked.
//
// Returns:
//   - bool: Returns true if the string consists solely of underscores, otherwise returns false.
//
// Notes:
//   - An empty string is considered as not composed solely of underscores.
func isUnderscore(v string) bool {
	regex := regexp.MustCompile(`^[_]+$`)
	return regex.MatchString(v)
}

// The [loader.isValidDir] function checks if the input path is a valid directory.
//
// Parameters:
//   - path (string): The path to be checked.
//
// Returns:
//   - error: Returns an error if the input path is not a valid directory
//     or if there are permission issues accessing the path.
//
// Notes:
//
//   - If the application lacks permission to read the directory, an error may
//     be returned even if the path is valid.
//   - The input path can be either relative or absolute. Both formats are supported.
func isValidDir(path string) error {
	// Get file or directory info
	info, err := os.Stat(path)

	if err != nil {
		// If the path doesn't exist or is invalid
		if os.IsNotExist(err) {
			return newInexistentDirError()
		}
		// Other errors (e.g., permission issues)
		return newFailedToReadDirError()
	}

	// Check if the path is a directory
	if info.IsDir() {
		return nil
	}

	return newPathIsNotDirError()
}

// The [loader.fileExists] function checks if the input path exists and
// is a valid file.
//
// Parameters:
//   - path (string): The path to be checked.
//
// Returns:
//   - error: Returns an error if the input path is not a valid file or
//     if there are permission issues accessing the file.
//
// Notes:
//   - If the application lacks permission to read the file, an error may be
//     returned even if the path is valid.
//   - The input path can be either relative or absolute. Both formats
//     are supported.
func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, newFailedToReadEnvError()
	}

	if info.IsDir() {
		return false, newPathIsNotFileError()
	}

	return true, nil
}

// The [loader.openFile] function opens the specified file, loads it into
// memory, and returns a buffer containing the file's bytes.
//
// Parameters:
//   - path (string): The path to the file.
//
// Returns:
//   - error: Returns an error if the input path is not a valid file or
//     if there are permission issues accessing the file.
//
// Notes:
//   - If the application lacks permission to read the file, an error may
//     be returned even if the path is valid.
//   - The input path can be either relative or absolute. Both formats
//     are supported.
func openFile(path string) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	file, err := os.Open(path)

	if err != nil {
		return buffer, newFailedToReadEnvError()
	}

	defer file.Close()

	_, err = io.Copy(&buffer, file)

	if err != nil {
		return buffer, newFailedToReadEnvError()
	}

	return buffer, nil
}
