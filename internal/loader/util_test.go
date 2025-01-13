package loader

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
)

func Test_isValidFolder(t *testing.T) {
	t.Run("valid folder", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempName := "test"

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer os.RemoveAll(tempDir)

		err = isValidFolder(tempDir)

		if err != nil {
			assert.FailNow(err.Error())
		}

		// Assert
		assert.Nil(err)
	})

	t.Run("folder not found", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		expectedErr := newEnvFolderNotExistError()

		// Act
		err := isValidFolder("unknown")

		// Assert
		if err == nil {
			assert.FailNow("Error should not be nil.")
		}

		castErr, ok := err.(*errors.GofidentialError)

		if !ok {
			assert.FailNow("Error is not of type GofidentialError.")
		}

		assert.Equal(castErr.Issuer, expectedErr.Issuer)
		assert.Equal(castErr.Code, expectedErr.Code)
		assert.Equal(castErr.Message, expectedErr.Message)
		assert.WithinDuration(castErr.Timestamp, expectedErr.Timestamp, 5*time.Second)
		assert.Equal(castErr.Suggestion, expectedErr.Suggestion)
		assert.NotEmpty(castErr.StackTrace, expectedErr.StackTrace)
		assert.Equal(castErr.Details, expectedErr.Details)
	})

	t.Run("filepath", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempName := "test"
		tempFilename := "example.txt"
		expectedErr := newPathIsNotFolderError()

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer os.RemoveAll(tempDir)

		fp := filepath.Join(tempDir, tempFilename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}
		defer tempFile.Close()

		err = isValidFolder(fp)

		// Assert
		if err == nil {
			assert.FailNow("Error should not be nil.")
		}

		castErr, ok := err.(*errors.GofidentialError)

		if !ok {
			assert.FailNow("Error is not of type GofidentialError.")
		}

		assert.Equal(castErr.Issuer, expectedErr.Issuer)
		assert.Equal(castErr.Code, expectedErr.Code)
		assert.Equal(castErr.Message, expectedErr.Message)
		assert.WithinDuration(castErr.Timestamp, expectedErr.Timestamp, 5*time.Second)
		assert.Equal(castErr.Suggestion, expectedErr.Suggestion)
		assert.NotEmpty(castErr.StackTrace, expectedErr.StackTrace)
		assert.Equal(castErr.Details, expectedErr.Details)
	})
}

func Test_fileExists(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempName := "test"
		tempFilename := "example.txt"

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer os.RemoveAll(tempDir)

		fp := filepath.Join(tempDir, tempFilename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}
		defer tempFile.Close()

		exists, err := fileExists(fp)

		if err != nil {
			assert.FailNow("Error should not be nil.")
		}

		assert.True(exists)
		assert.Nil(err)
	})

	t.Run("inexistent file", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempFilename := "unknown.txt."

		// Act
		exists, err := fileExists(tempFilename)

		// Assert
		if err != nil {
			assert.FailNow(err.Error())
		}

		assert.False(exists)
		assert.Nil(err)
	})

	t.Run("inexistent file", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempName := "test"
		expectedErr := newPathIsNotFileError()

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer os.RemoveAll(tempDir)

		exists, err := fileExists(tempDir)

		// Assert
		if err == nil {
			assert.FailNow("Error should not be nil.")
		}

		castErr, ok := err.(*errors.GofidentialError)

		if !ok {
			assert.FailNow("Error is not of type GofidentialError.")
		}

		assert.False(exists)
		assert.Equal(castErr.Issuer, expectedErr.Issuer)
		assert.Equal(castErr.Code, expectedErr.Code)
		assert.Equal(castErr.Message, expectedErr.Message)
		assert.WithinDuration(castErr.Timestamp, expectedErr.Timestamp, 5*time.Second)
		assert.Equal(castErr.Suggestion, expectedErr.Suggestion)
		assert.NotEmpty(castErr.StackTrace, expectedErr.StackTrace)
		assert.Equal(castErr.Details, expectedErr.Details)
	})
}

func Test_openFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		tempName := "test"
		tempFilename := "example.txt"
		fileContent := "Hello world!"

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer os.RemoveAll(tempDir)

		fp := filepath.Join(tempDir, tempFilename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer tempFile.Close()

		if _, err = tempFile.WriteString(fileContent); err != nil {
			assert.FailNow(err.Error())
		}

		actualBuffer, err := openFile(fp)

		if err != nil {
			assert.FailNow(err.Error())
		}

		var expectedBuffer bytes.Buffer
		file, err := os.Open(fp)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer file.Close()

		_, err = io.Copy(&expectedBuffer, file)

		assert.Equal(actualBuffer.String(), expectedBuffer.String())
		assert.NotEmpty(actualBuffer.String())
		assert.NotEqual(actualBuffer.Len(), 0)
		assert.Nil(err)
	})

	t.Run("inexistent file", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		filename := "unknown.txt"
		expectedErr := newFailedToReadFileError()

		// Act
		buffer, err := openFile(filename)

		// Assert
		if err == nil {
			assert.FailNow("Error should not be nil.")
		}

		castErr, ok := err.(*errors.GofidentialError)

		if !ok {
			assert.FailNow("Error is not of type GofidentialError.")
		}

		assert.Empty(buffer.String())
		assert.Equal(buffer.Len(), 0)
		assert.Equal(castErr.Issuer, expectedErr.Issuer)
		assert.Equal(castErr.Code, expectedErr.Code)
		assert.Equal(castErr.Message, expectedErr.Message)
		assert.WithinDuration(castErr.Timestamp, expectedErr.Timestamp, 5*time.Second)
		assert.Equal(castErr.Suggestion, expectedErr.Suggestion)
		assert.NotEmpty(castErr.StackTrace, expectedErr.StackTrace)
		assert.Equal(castErr.Details, expectedErr.Details)
	})
}
