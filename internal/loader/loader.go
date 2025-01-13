package loader

import (
	"bytes"
	"path/filepath"
	"strings"
)

func Load(name string, overridePath string, ignoreFilename bool) (bytes.Buffer, error) {
	// Must be trimmed to pass subsequent error checks.
	name = strings.TrimSpace(name)
	overridePath = strings.TrimSpace(overridePath)

	// If the name of the environment is not defined, then it's not possible
	// to know the .env filename.
	if name == "" {
		return bytes.Buffer{}, newMissingEnvNameError()
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
	if err := isValidFolder(overridePath); err != nil {
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
