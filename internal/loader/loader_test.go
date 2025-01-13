package loader

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := ""
		defaultPath := ""
		ignoreFilename := false
		expectedErr := newMissingEnvNameError()

		// Act
		buffer, err := Load(name, defaultPath, ignoreFilename)

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

	t.Run("no path override file exists in cwd", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := false
		tempName := "test"
		filename := "dev.env"
		fileContents := "Hello world!"

		// Act
		originalDir, err := os.Getwd()
		if err != nil {
			assert.FailNow(err.Error())
		}

		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer func() {
			if err := os.Chdir(originalDir); err != nil {
				assert.FailNow(err.Error())
			}

			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		if err = os.Chdir(tempDir); err != nil {
			assert.FailNow(err.Error())
		}

		tempFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, "", ignoreFilename)

		// Assert
		if err != nil {
			assert.FailNow(err.Error())
		}

		assert.Equal(buffer.String(), fileContents)
		assert.Equal(buffer.Len(), len([]byte(fileContents)))
		assert.Nil(err)
	})

	t.Run("no path override file exists in cwd but with a different name", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := false
		tempName := "test"
		filename := "prod.env"
		expectedErr := newEnvNotFoundError()

		// Act
		originalDir, err := os.Getwd()
		if err != nil {
			assert.FailNow(err.Error())
		}

		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer func() {
			if err := os.Chdir(originalDir); err != nil {
				assert.FailNow(err.Error())
			}

			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		if err = os.Chdir(tempDir); err != nil {
			assert.FailNow(err.Error())
		}

		tempFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, "", ignoreFilename)

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

	t.Run("no path override file exists in cwd but with a different name ignore filename", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := true
		tempName := "test"
		filename := ".env"
		fileContents := "Hello world!"

		// Act
		originalDir, err := os.Getwd()
		if err != nil {
			assert.FailNow(err.Error())
		}

		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer func() {
			if err := os.Chdir(originalDir); err != nil {
				assert.FailNow(err.Error())
			}

			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		if err = os.Chdir(tempDir); err != nil {
			assert.FailNow(err.Error())
		}

		tempFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, "", ignoreFilename)

		// Assert
		if err != nil {
			assert.FailNow(err.Error())
		}

		assert.Equal(buffer.String(), fileContents)
		assert.Equal(buffer.Len(), len([]byte(fileContents)))
		assert.Nil(err)
	})

	t.Run("no path override file exists in cwd but with a different name ignore filename", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := true
		tempName := "test"
		filename := "dev.env"
		expectedErr := newEnvNotFoundError()

		// Act
		originalDir, err := os.Getwd()
		if err != nil {
			assert.FailNow(err.Error())
		}

		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer func() {
			if err := os.Chdir(originalDir); err != nil {
				assert.FailNow(err.Error())
			}

			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		if err = os.Chdir(tempDir); err != nil {
			assert.FailNow(err.Error())
		}

		tempFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, "", ignoreFilename)

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

	t.Run("specified path file exists", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := false
		tempName := "test"
		filename := "dev.env"
		fileContents := "Hello world!"

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		pathOverride := tempDir

		defer func() {
			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		fp := filepath.Join(tempDir, filename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, pathOverride, ignoreFilename)

		// Assert
		if err != nil {
			assert.FailNow(err.Error())
		}

		assert.Equal(buffer.String(), fileContents)
		assert.Equal(buffer.Len(), len([]byte(fileContents)))
		assert.Nil(err)
	})

	t.Run("specified path file exists but with a different name", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := false
		tempName := "test"
		filename := "prod.env"
		fileContents := "Hello world!"
		expectedErr := newEnvNotFoundError()

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		pathOverride := tempDir

		defer func() {
			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		fp := filepath.Join(tempDir, filename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, pathOverride, ignoreFilename)

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

	t.Run("specified path file exists ignore filename", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := true
		tempName := "test"
		filename := ".env"
		fileContents := "Hello world!"

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		pathOverride := tempDir

		defer func() {
			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		fp := filepath.Join(tempDir, filename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, pathOverride, ignoreFilename)

		// Assert
		if err != nil {
			assert.FailNow(err.Error())
		}

		assert.Equal(buffer.String(), fileContents)
		assert.Equal(buffer.Len(), len([]byte(fileContents)))
		assert.Nil(err)
	})

	t.Run("specified path file exists but with a different name ignore filename", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := true
		tempName := "test"
		filename := "dev.env"
		fileContents := "Hello world!"
		expectedErr := newEnvNotFoundError()

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		pathOverride := tempDir

		defer func() {
			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		fp := filepath.Join(tempDir, filename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		buffer, err := Load(name, pathOverride, ignoreFilename)

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

	t.Run("specified path is not a valid directory", func(t *testing.T) {
		// Arrange
		assert := assert.New(t)
		name := "dev"
		ignoreFilename := false
		tempName := "test"
		filename := "dev.env"
		fileContents := "Hello world!"
		expectedErr := newPathIsNotFolderError()

		// Act
		tempDir, err := os.MkdirTemp("", tempName)

		if err != nil {
			assert.FailNow(err.Error())
		}

		defer func() {
			if err = os.RemoveAll(tempDir); err != nil {
				assert.FailNow(err.Error())
			}
		}()

		fp := filepath.Join(tempDir, filename)
		tempFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			assert.FailNow(err.Error())
		}

		_, err = tempFile.WriteString(fileContents)

		if err != nil {
			assert.FailNow(err.Error())
		}

		tempFile.Close()

		pathOverride := fp
		buffer, err := Load(name, pathOverride, ignoreFilename)

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
