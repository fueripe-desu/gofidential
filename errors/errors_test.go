package errors

import (
	"testing"

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
				Code:    "TEST_ERROR",
				Message: "This is a test error.",
				Hint:    "This is an example hint.",
			},
			err2: &GofidentialError{
				Code:    "TEST_ERROR",
				Message: "This is a test error.",
				Hint:    "This is an example hint.",
			},
			expected: true,
		},
		{
			name: "should return false if code is different",
			err1: &GofidentialError{
				Code:    "TEST_ERROR 1",
				Message: "This is a test error.",
				Hint:    "This is an example hint.",
			},
			err2: &GofidentialError{
				Code:    "TEST_ERROR 2",
				Message: "This is a test error.",
				Hint:    "This an example hint.",
			},
			expected: false,
		},
		{
			name: "should return false if message is different",
			err1: &GofidentialError{
				Code:    "TEST_ERROR",
				Message: "This is the first test error.",
				Hint:    "This is an example hint.",
			},
			err2: &GofidentialError{
				Code:    "TEST_ERROR",
				Message: "This is the second test error.",
				Hint:    "This is an example hint.",
			},
			expected: false,
		},
		{
			name: "should return false if hint is different",
			err1: &GofidentialError{
				Code:    "TEST_ERROR",
				Message: "This is a test error.",
				Hint:    "This is the example hint 1.",
			},
			err2: &GofidentialError{
				Code:    "TEST_ERROR",
				Message: "This is a test error.",
				Hint:    "This is the example hint 2.",
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
