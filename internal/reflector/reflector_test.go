package reflector

import (
	"reflect"
	"testing"
	"time"

	"github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
)

func Test_IsEmpty(t *testing.T) {
	testcases := []struct {
		name     string
		data     any
		expected bool
	}{
		{
			name:     "empty struct",
			data:     &(struct{}{}),
			expected: true,
		},
		{
			name:     "non empty struct",
			data:     &(struct{ Field1 string }{}),
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			assert := assert.New(t)

			// Act
			r, err := New(tc.data)

			if err != nil {
				assert.FailNow(err.Error())
			}

			isEmpty := r.IsEmpty()

			// Assert
			assert.Equal(isEmpty, tc.expected)
		})
	}
}

func Test_New(t *testing.T) {
	testcases := []struct {
		name        string
		data        any
		expectedErr *errors.GofidentialError
	}{
		{
			name:        "data is nil",
			data:        nil,
			expectedErr: newDataIsNilError(),
		},
		{
			name:        "data is not a pointer",
			data:        1,
			expectedErr: newDataIsNotPtrError(),
		},
		{
			name: "data is nil pointer",
			data: func() any {
				var data *struct{} = nil
				return data
			}(),
			expectedErr: newDataIsNilPtrError(),
		},
		{
			name:        "data does not point to a struct",
			data:        &([]string{}),
			expectedErr: newDataIsNotStructError(),
		},
		{
			name: "valid pointer",
			data: &(struct{}{}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			assert := assert.New(t)

			// Act
			r, err := New(tc.data)

			// Assert
			if tc.expectedErr != nil {
				if err == nil {
					assert.FailNow("Error should not be nil.")
				}

				castErr, ok := err.(*errors.GofidentialError)

				if !ok {
					assert.FailNow("Error is not of type GofidentialError")
				}

				assert.Nil(r)
				assert.Equal(castErr.Issuer, tc.expectedErr.Issuer)
				assert.Equal(castErr.Code, tc.expectedErr.Code)
				assert.Equal(castErr.Message, tc.expectedErr.Message)
				assert.WithinDuration(castErr.Timestamp, tc.expectedErr.Timestamp, 5*time.Second)
				assert.NotEmpty(castErr.StackTrace)
				assert.Equal(castErr.Suggestion, tc.expectedErr.Suggestion)
				assert.Equal(castErr.Details, tc.expectedErr.Details)
			} else {
				if err != nil {
					assert.FailNow(err.Error())
				}

				expectedPtr := reflect.ValueOf(tc.data).Elem()

				assert.True(r.ptrval.Equal(expectedPtr))
				assert.Nil(err)
			}
		})
	}
}
