package loader

import (
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
)

func isLowerUnderscore(v string) bool {
	regex := regexp.MustCompile(`^[a-z_]+$`)
	return regex.MatchString(v)
}

func isUnderscore(v string) bool {
	regex := regexp.MustCompile(`^[_]+$`)
	return regex.MatchString(v)
}

func isValidDir(path string) error {
	// Get file or directory info
	info, err := os.Stat(path)

	if err != nil {
		// If the path doesn't exist or is invalid
		if os.IsNotExist(err) {
			return newEnvDirNotExistError()
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

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, newFailedToReadFileError()
	}

	if info.IsDir() {
		return false, newPathIsNotFileError()
	}

	return true, nil
}

func openFile(path string) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	file, err := os.Open(path)

	if err != nil {
		return buffer, newFailedToReadFileError()
	}

	defer file.Close()

	_, err = io.Copy(&buffer, file)

	if err != nil {
		return buffer, newFailedToReadFileError()
	}

	return buffer, nil
}
