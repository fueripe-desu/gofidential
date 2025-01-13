package errors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Equal(t *testing.T) {
	testcases := []struct {
		name     string
		expected bool
		err1     *GofidentialError
		err2     *GofidentialError
	}{
		{
			name: "should return true if both errors are equal",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: true,
		},
		{
			name: "should return false if issuer is different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer 1",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer 2",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if code is different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR 1",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR 2",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if message is different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is the first test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is the second test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if timestamp is different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 2, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if the first stack trace is empty",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte{},
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if the second stack trace is empty",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte{},
				Suggestion: "This is an example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if suggestion is different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion 1.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion 2.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			expected: false,
		},
		{
			name: "should return false if details is completely different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion.",
				Details: map[string]string{
					"value1": "example 3",
					"value2": "example 4",
				},
			},
			expected: false,
		},
		{
			name: "should return false if details is partially different",
			err1: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 2",
				},
			},
			err2: &GofidentialError{
				Issuer:     "Mock issuer",
				Code:       "TEST_ERROR",
				Message:    "This is a test error.",
				Timestamp:  time.Date(2020, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
				StackTrace: []byte("example stacktrace"),
				Suggestion: "This is the example suggestion.",
				Details: map[string]string{
					"value1": "example 1",
					"value2": "example 3",
				},
			},
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			assert := assert.New(t)

			// Act
			isEqual := tc.err1.Equal(tc.err2)

			// Assert
			assert.Equal(isEqual, tc.expected)
		})
	}
}
