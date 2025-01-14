package loader

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isValidDir(t *testing.T) {
	t.Run("should return true if the path points to a valid dir", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		tempDirName := "test"

		tempDir, err := os.MkdirTemp("", tempDirName)
		require.NoError(err, "Failed to create temporary dir")

		t.Cleanup(func() {
			err = os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		err = isValidDir(tempDir)
		require.NoError(err, "An unexpected error ocurred")

		require.Nil(err)
	})

	t.Run("should return an error if the specified folder was not found", func(t *testing.T) {
		require := require.New(t)

		expectedErr := newEnvDirNotExistError()

		err := isValidDir("unknown")
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		require.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})

	t.Run("should return an error if the specified path does not point to a dir", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		tempDirName := "test"
		tempFilename := "example.txt"

		expectedErr := newPathIsNotDirError()

		tempDir, err := os.MkdirTemp("", tempDirName)
		require.NoError(err, "Failed to create temporary dir")

		fp := filepath.Join(tempDir, tempFilename)

		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
		require.NoError(err, "Failed to create file in temporary dir.")

		t.Cleanup(func() {
			err := tempFile.Close()
			assert.NoError(err, "Failed to close file in temporary dir.")

			err = os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		err = isValidDir(fp)
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		require.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})
}

func Test_fileExists(t *testing.T) {
	t.Run("should return true if the file exists", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		tempDirName := "test"
		tempFilename := "example.txt"

		tempDir, err := os.MkdirTemp("", tempDirName)
		require.NoError(err, "Failed to create temporary dir")

		fp := filepath.Join(tempDir, tempFilename)

		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
		require.NoError(err, "Failed to create file in temporary dir.")

		t.Cleanup(func() {
			err := tempFile.Close()
			assert.NoError(err, "Failed to close file in temporary dir.")

			err = os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		exists, err := fileExists(fp)
		require.NoError(err, "An unexpected error ocurred")

		require.True(exists)
		require.Nil(err)
	})

	t.Run("should return false if the file does not exist", func(t *testing.T) {
		require := require.New(t)
		tempFilename := "unknown.txt"

		exists, err := fileExists(tempFilename)
		require.NoError(err, "An unexpected error ocurred")

		require.False(exists)
		require.Nil(err)
	})

	t.Run("should return an error if the specified path does not point to a file", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		tempName := "test"

		expectedErr := newPathIsNotFileError()

		tempDir, err := os.MkdirTemp("", tempName)
		require.NoError(err, "Failed to create temporary dir")

		t.Cleanup(func() {
			err := os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		exists, err := fileExists(tempDir)
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		require.False(exists)
		require.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})
}

func Test_openFile(t *testing.T) {
	t.Run("should return a valid byte buffer if file exists", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		tempName := "test"
		tempFilename := "example.txt"

		fileContent := "Hello world!"

		tempDir, err := os.MkdirTemp("", tempName)
		require.NoError(err, "Failed to create temporary dir")

		fp := filepath.Join(tempDir, tempFilename)

		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
		require.NoError(err, "Failed to create file in temporary dir.")

		t.Cleanup(func() {
			err = os.RemoveAll(tempDir)
			assert.NoError(err, "Failed to remove temporary dir")
		})

		_, err = tempFile.WriteString(fileContent)
		require.NoError(err, "Failed writing string to file.")

		err = tempFile.Close()
		assert.NoError(err, "Failed to close file in temporary dir.")

		actualBuffer, err := openFile(fp)
		require.NoError(err, "An unexpected error ocurred")

		var expectedBuffer bytes.Buffer

		file, err := os.Open(fp)
		require.NoError(err, "Could not open file to fill expected buffer.")

		_, err = io.Copy(&expectedBuffer, file)

		err = file.Close()
		assert.NoError(err, "Could not close expected buffer file.")

		assert.Equal(actualBuffer.String(), expectedBuffer.String())
		assert.NotEmpty(actualBuffer.String())
		assert.NotEqual(actualBuffer.Len(), 0)
		assert.Nil(err)
	})

	t.Run("should return an error if the specified file path does not exist", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		filename := "unknown.txt"

		expectedErr := newFailedToReadFileError()

		buffer, err := openFile(filename)
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		assert.Empty(buffer.String())
		assert.Equal(buffer.Len(), 0)
		require.Error(err, "An error was expected. But got none.")
		require.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})
}
