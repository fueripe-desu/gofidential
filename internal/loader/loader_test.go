package loader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Load(t *testing.T) {
	t.Run("should return an error if name is empty", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		name := ""
		defaultPath := ""
		ignoreFilename := false

		expectedErr := newMissingEnvNameError()

		buffer, err := Load(name, defaultPath, ignoreFilename)
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		assert.Empty(buffer.String())
		assert.Equal(buffer.Len(), 0)
		assert.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})

	defaultPathTestcases := []struct {
		name string

		envName        string
		ignoreFilename bool

		tempDirName      string
		tempFilename     string
		tempFileContents string

		expectedErr *errors.GofidentialError
	}{
		{
			name:             "should return a byte buffer if file exists in default path",
			envName:          "dev",
			ignoreFilename:   false,
			tempDirName:      "test",
			tempFilename:     "dev.env",
			tempFileContents: "Hello world!",
		},
		{
			name:             "should return an error if filename does not match the environment",
			envName:          "dev",
			ignoreFilename:   false,
			tempDirName:      "test",
			tempFilename:     "prod.env",
			tempFileContents: "Hello world!",
			expectedErr:      newEnvNotFoundError(),
		},
		{
			name:             "should return a byte buffer if filename is ignored and a nameless file exists",
			envName:          "dev",
			ignoreFilename:   true,
			tempDirName:      "test",
			tempFilename:     ".env",
			tempFileContents: "Hello world!",
		},
		{
			name:             "should return an error if filename is ignored and a nameless file does not exist",
			envName:          "dev",
			ignoreFilename:   true,
			tempDirName:      "test",
			tempFilename:     "dev.env",
			tempFileContents: "Hello world!",
			expectedErr:      newEnvNotFoundError(),
		},
	}

	for _, tc := range defaultPathTestcases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			originalDir, err := os.Getwd()
			require.NoError(err, "Failed to get the current working directory.")

			tempDir, err := os.MkdirTemp("", tc.tempDirName)
			require.NoError(err, "Failed to create temporary dir")

			err = os.Chdir(tempDir)
			require.NoError(err, "Failed to change working directory to the temporary directory.")

			tempFile, err := os.OpenFile(tc.tempFilename, os.O_CREATE|os.O_WRONLY, 0644)
			require.NoError(err, "Failed to create file in temporary dir.")

			_, err = tempFile.WriteString(tc.tempFileContents)
			require.NoError(err, "Failed writing string to file.")

			t.Cleanup(func() {
				err := tempFile.Close()
				assert.NoError(err, "Failed to close file in temporary dir.")

				err = os.Chdir(originalDir)
				assert.NoError(err, "Failed to return to the original working directory.")

				err = os.RemoveAll(tempDir)
				assert.NoError(err, "Failed to remove temporary dir")
			})

			buffer, err := Load(tc.envName, "", tc.ignoreFilename)

			if tc.expectedErr == nil {
				require.NoError(err, "An unexpected error ocurred")

				assert.Equal(buffer.String(), tc.tempFileContents)
				assert.Equal(buffer.Len(), len([]byte(tc.tempFileContents)))
				assert.Nil(err)
			} else {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

				assert.Empty(buffer.String())
				assert.Equal(buffer.Len(), 0)
				assert.Error(err, "An error was expected. But got none.")
				assert.True(
					castErr.Equal(tc.expectedErr),
					"The actual error does not match the expected one. Actual: %v, Expected: %v",
					castErr,
					tc.expectedErr,
				)
			}
		})
	}

	pathOverrideTestcases := []struct {
		name string

		envName        string
		ignoreFilename bool

		tempDirName      string
		tempFilename     string
		tempFileContents string

		expectedErr *errors.GofidentialError
	}{
		{
			name:             "should return a byte buffer if file exists in overriden path",
			envName:          "dev",
			ignoreFilename:   false,
			tempDirName:      "test",
			tempFilename:     "dev.env",
			tempFileContents: "Hello world!",
		},
		{
			name:             "should return an if filename does not match the environment in overriden path",
			envName:          "dev",
			ignoreFilename:   false,
			tempDirName:      "test",
			tempFilename:     "prod.env",
			tempFileContents: "Hello world!",
			expectedErr:      newEnvNotFoundError(),
		},
		{
			name:             "should return a byte buffer if filename is ignored and a nameless file exist in overriden path",
			envName:          "dev",
			ignoreFilename:   true,
			tempDirName:      "test",
			tempFilename:     ".env",
			tempFileContents: "Hello world!",
		},
		{
			name:             "should return an error if filename is ignored and a nameless does not file exist in overriden path",
			envName:          "dev",
			ignoreFilename:   true,
			tempDirName:      "test",
			tempFilename:     "dev.env",
			tempFileContents: "Hello world!",
			expectedErr:      newEnvNotFoundError(),
		},
	}

	for _, tc := range pathOverrideTestcases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			tempDir, err := os.MkdirTemp("", tc.tempDirName)
			require.NoError(err, "Failed to create temporary dir")

			t.Cleanup(func() {
				err := os.RemoveAll(tempDir)
				assert.NoError(err, "Failed to remove temporary dir")
			})

			fp := filepath.Join(tempDir, tc.tempFilename)

			tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
			require.NoError(err, "Failed to create file in temporary dir.")

			_, err = tempFile.WriteString(tc.tempFileContents)
			require.NoError(err, "Failed writing string to file.")

			err = tempFile.Close()
			assert.NoError(err, "Failed to close file in temporary dir.")

			pathOverride := tempDir

			buffer, err := Load(tc.envName, pathOverride, tc.ignoreFilename)

			if tc.expectedErr == nil {
				require.NoError(err, "An unexpected error ocurred")
				assert.Equal(buffer.String(), tc.tempFileContents)
				assert.Equal(buffer.Len(), len([]byte(tc.tempFileContents)))
				assert.Nil(err)
			} else {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

				assert.Empty(buffer.String())
				assert.Equal(buffer.Len(), 0)
				assert.Error(err, "An error was expected. But got none.")
				assert.True(
					castErr.Equal(tc.expectedErr),
					"The actual error does not match the expected one. Actual: %v, Expected: %v",
					castErr,
					tc.expectedErr,
				)
			}
		})
	}

	t.Run("should return an error if overriden path is a file instead of a dir", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		name := "dev"
		ignoreFilename := false

		tempDirName := "test"
		tempFilename := "dev.env"
		tempFileContents := "Hello world!"

		expectedErr := newPathIsNotDirError()

		tempDir, err := os.MkdirTemp("", tempDirName)
		require.NoError(err, "Failed to create temporary dir")

		t.Cleanup(func() {
			err := os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		fp := filepath.Join(tempDir, tempFilename)

		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
		require.NoError(err, "Failed to create file in temporary dir.")

		_, err = tempFile.WriteString(tempFileContents)
		require.NoError(err, "Failed writing string to file.")

		err = tempFile.Close()
		assert.NoError(err, "Failed to close file in temporary dir.")

		pathOverride := fp
		buffer, err := Load(name, pathOverride, ignoreFilename)
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		assert.Empty(buffer.String())
		assert.Equal(buffer.Len(), 0)
		assert.Error(err, "An error was expected. But got none.")
		assert.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})
}
