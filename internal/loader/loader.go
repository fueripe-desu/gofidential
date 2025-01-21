// Package loader provides functionality to load a .env file into memory
// and return its contents as a byte buffer.
//
// The core component of this package is the [loader.loader] struct, which
// manages the file loading process. The package exposes the [loader.Load]
// function, which initiates the file loading and generates the byte buffer
// containing the file's contents.
package loader

import (
	"bytes"
	"path/filepath"
	"strings"
)

// The [loader.Load] function loads the specified .env file into memory and
// returns a buffer containing the file's bytes.
//
// Parameters:
//   - name (string): The name of the environment. This field is required.
//   - overridePath (string): An optional custom directory path to look for the .env file.
//   - ignoreFilename (bool): Determines the naming convention for the .env file.
//     If true, the filename will always be ".env". If false, the filename will include
//     the environment name (e.g., "dev.env", "test.env", "prod.env").
//
// Returns:
//   - bytes.Buffer: The loaded file's byte buffer.
//   - error: An error explaning why the loading process failed, or nil if the operation was
//     successful.
//
// Constraints:
//   - The "name" parameter must not be empty.
//   - The "name" parameter must contain only lowercase letters and underscores.
//   - The "name" parameter must not have leading or trailing underscores.
//   - The "name" parameter must not consist solely of underscores.
//   - If specified, the "overridePath" parameter must be a valid directory path
//     (not a filepath).
//
// Notes:
//   - The "overridePath" and "ignoreFilename" parameters are optional.
func Load(name string, overridePath string, ignoreFilename bool) (bytes.Buffer, error) {
	// Must be trimmed to pass subsequent error checks.
	name = strings.TrimSpace(name)
	overridePath = strings.TrimSpace(overridePath)

	// If the name of the environment is not defined, then it's not possible
	// to know the .env filename.
	if name == "" {
		return bytes.Buffer{}, newMissingNameError()
	}

	// If the env name contains uppercase letters, numbers or special characters,
	// it must be rejected.
	if !isLowerUnderscore(name) {
		return bytes.Buffer{}, newInvalidNameError()
	}

	// If the env name is composed by underscores only, it must be rejected.
	if isUnderscore(name) {
		return bytes.Buffer{}, newUnderscoreOnlyNameError()
	}

	// If the env name contains a leading underscore, it must be rejected.
	if name[0] == '_' {
		return bytes.Buffer{}, newLeadingUnderscoreError()
	}

	// If the env name contains a trailing underscore, it must be rejected.
	if name[len(name)-1] == '_' {
		return bytes.Buffer{}, newTrailingUnderscoreError()
	}

	// If the filename must be ignored, then the non-empty name is set to an empty string,
	// this way, the resulting filename will be only ".env".
	if ignoreFilename {
		name = ""
	}

	filename := name + ".env"

	if overridePath == "" {
		// If there is no specified path, first check if the filename exists
		// in the current working directory.
		exists, err := fileExists(filename)

		if err != nil {
			return bytes.Buffer{}, err
		}

		if exists {
			return openFile(filename)
		}

		// As a fallback, if it does not exist, check for it in the special
		// "env" directory in the current working directory.
		envPath := filepath.Join("env", filename)
		exists, err = fileExists(envPath)

		if err != nil {
			return bytes.Buffer{}, err
		}

		if !exists {
			return bytes.Buffer{}, newEnvNotFoundError()
		}

		return openFile(envPath)
	}

	// If a path was specified, then check if it points to a valid directory.
	if err := isValidDir(overridePath); err != nil {
		return bytes.Buffer{}, err
	}

	// Then lastly, check if the .env file exists in this specified directory.
	envPath := filepath.Join(overridePath, filename)
	exists, err := fileExists(envPath)

	if err != nil {
		return bytes.Buffer{}, err
	}

	if !exists {
		return bytes.Buffer{}, newEnvNotFoundError()
	}

	return openFile(envPath)
}
