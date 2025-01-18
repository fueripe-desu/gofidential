package reflector

import (
	"reflect"
	"testing"
	"time"

	errors "github.com/fueripe-desu/gofidential/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newReflector(t *testing.T) {
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
			require := require.New(t)
			assert := assert.New(t)

			r, err := newReflector(tc.data)

			if tc.expectedErr != nil {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

				assert.Nil(r)
				assert.Error(err, "An error was expected. But got none.")
				assert.True(
					castErr.Equal(tc.expectedErr),
					"The actual error does not match the expected one. Actual: %v, Expected: %v",
					castErr,
					tc.expectedErr,
				)
			} else {
				require.NoError(err, "An unexpected error ocurred")

				expectedPtr := reflect.ValueOf(tc.data).Elem()

				assert.True(r.ptrval.Equal(expectedPtr))
				assert.Nil(err)
			}
		})
	}
}

func Test_AllFields(t *testing.T) {
	testcases := []struct {
		name     string
		data     any
		expected map[string]any
	}{
		{
			name: "should return a valid map when the struct has an int field",
			data: &(struct{ IntField int }{IntField: 1}),
			expected: map[string]any{
				"IntField": 1,
			},
		},
		{
			name: "should return a valid map when struct has an int8 field",
			data: &(struct{ SmallIntField int8 }{SmallIntField: 8}),
			expected: map[string]any{
				"SmallIntField": int8(8),
			},
		},
		{
			name: "should return a valid map when struct has an int16 field",
			data: &(struct{ ShortIntField int16 }{ShortIntField: 16}),
			expected: map[string]any{
				"ShortIntField": int16(16),
			},
		},
		{
			name: "should return a valid map when struct has an int32 field",
			data: &(struct{ NormalIntField int32 }{NormalIntField: 32}),
			expected: map[string]any{
				"NormalIntField": int32(32),
			},
		},
		{
			name: "should return a valid map when struct has an int64 field",
			data: &(struct{ LargeIntField int64 }{LargeIntField: 64}),
			expected: map[string]any{
				"LargeIntField": int64(64),
			},
		},
		{
			name: "should return a valid map when the struct has an uint field",
			data: &(struct{ UintField uint }{UintField: 1}),
			expected: map[string]any{
				"UintField": uint(1),
			},
		},
		{
			name: "should return a valid map when struct has an uint8 field",
			data: &(struct{ SmallUintField uint8 }{SmallUintField: 8}),
			expected: map[string]any{
				"SmallUintField": uint8(8),
			},
		},
		{
			name: "should return a valid map when struct has an uint16 field",
			data: &(struct{ ShortUintField uint16 }{ShortUintField: 16}),
			expected: map[string]any{
				"ShortUintField": uint16(16),
			},
		},
		{
			name: "should return a valid map when struct has an uint32 field",
			data: &(struct{ NormalUintField uint32 }{NormalUintField: 32}),
			expected: map[string]any{
				"NormalUintField": uint32(32),
			},
		},
		{
			name: "should return a valid map when struct has an uint64 field",
			data: &(struct{ LargeUintField uint64 }{LargeUintField: 64}),
			expected: map[string]any{
				"LargeUintField": uint64(64),
			},
		},
		{
			name: "should return a valid map when struct has a float32 field",
			data: &(struct{ FloatField float32 }{FloatField: 32.0}),
			expected: map[string]any{
				"FloatField": float32(32),
			},
		},
		{
			name: "should return a valid map when struct has a float64 field",
			data: &(struct{ LargeFloatField float64 }{LargeFloatField: 64.0}),
			expected: map[string]any{
				"LargeFloatField": float64(64),
			},
		},
		{
			name: "should return a valid map when struct has a complex64 field",
			data: &(struct{ ComplexField complex64 }{
				ComplexField: complex(float32(3), float32(4)),
			}),
			expected: map[string]any{
				"ComplexField": complex(float32(3), float32(4)),
			},
		},
		{
			name: "should return a valid map when struct has a complex128 field",
			data: &(struct{ LargeComplexField complex128 }{
				LargeComplexField: complex(float64(3), float64(4)),
			}),
			expected: map[string]any{
				"LargeComplexField": complex(float64(3), float64(4)),
			},
		},
		{
			name: "should return a valid map when struct has a string field",
			data: &(struct{ StringField string }{StringField: "example string"}),
			expected: map[string]any{
				"StringField": "example string",
			},
		},
		{
			name: "should return a valid map when struct has a bool field",
			data: &(struct{ BoolField bool }{BoolField: true}),
			expected: map[string]any{
				"BoolField": true,
			},
		},
		{
			name: "should return a valid map when struct has a time field",
			data: &(struct{ TimeField time.Time }{
				TimeField: time.Date(2022, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
			}),
			expected: map[string]any{
				"TimeField": time.Date(2022, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
			},
		},
		{
			name: "should return a valid map when struct has fields of all types",
			data: &(struct {
				IntField          int
				SmallIntField     int8
				ShortIntField     int16
				NormalIntField    int32
				LargeIntField     int64
				UintField         uint
				SmallUintField    uint8
				ShortUintField    uint16
				NormalUintField   uint32
				LargeUintField    uint64
				FloatField        float32
				LargeFloatField   float64
				ComplexField      complex64
				LargeComplexField complex128
				StringField       string
				BoolField         bool
				TimeField         time.Time
			}{
				IntField:          1,
				SmallIntField:     8,
				ShortIntField:     16,
				NormalIntField:    32,
				LargeIntField:     64,
				UintField:         1,
				SmallUintField:    8,
				ShortUintField:    16,
				NormalUintField:   32,
				LargeUintField:    64,
				FloatField:        32.0,
				LargeFloatField:   64.0,
				ComplexField:      complex(float32(3), float32(4)),
				LargeComplexField: complex(float64(3), float64(4)),
				StringField:       "example string",
				BoolField:         true,
				TimeField:         time.Date(2022, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
			}),
			expected: map[string]any{
				"IntField":          int(1),
				"SmallIntField":     int8(8),
				"ShortIntField":     int16(16),
				"NormalIntField":    int32(32),
				"LargeIntField":     int64(64),
				"UintField":         uint(1),
				"SmallUintField":    uint8(8),
				"ShortUintField":    uint16(16),
				"NormalUintField":   uint32(32),
				"LargeUintField":    uint64(64),
				"FloatField":        float32(32),
				"LargeFloatField":   float64(64),
				"ComplexField":      complex(float32(3), float32(4)),
				"LargeComplexField": complex(float64(3), float64(4)),
				"StringField":       "example string",
				"BoolField":         true,
				"TimeField":         time.Date(2022, 1, 1, 1, 1, 1, 1, time.Now().UTC().Location()),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			r, err := newReflector(tc.data)
			require.NoError(err, "Failed to create reflector")

			result, err := r.AllFields()
			require.NoError(err, "Failed to get all fields from the struct")

			assert.Equal(result, tc.expected)
		})
	}

	t.Run("should return an error if the struct has unexported fields", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)
		expectedErr := newUnexportedFieldError()

		s := struct {
			unexported int
		}{
			unexported: 1,
		}

		r, err := newReflector(&s)
		require.NoError(err, "Failed to create reflector")

		result, err := r.AllFields()
		require.Error(err, "An error was expected. But got none.")

		castErr, ok := err.(*errors.GofidentialError)
		require.True(ok, "Error is not of type GofidentialError.")

		assert.Nil(result)
		assert.True(
			castErr.Equal(expectedErr),
			"The actual error does not match the expected one. Actual: %v, Expected: %v",
			castErr,
			expectedErr,
		)
	})
}

func Test_SetField(t *testing.T) {
	testcases := []struct {
		name         string
		data         any
		field        string
		value        string
		expectedData any
		expectedErr  *errors.GofidentialError
	}{
		{
			name:         "should be able to set an int field",
			data:         &(struct{ IntField int }{}),
			field:        "IntField",
			value:        "22",
			expectedData: &(struct{ IntField int }{IntField: 22}),
		},
		{
			name:         "should be able to set a negative int field",
			data:         &(struct{ IntField int }{}),
			field:        "IntField",
			value:        "-22",
			expectedData: &(struct{ IntField int }{IntField: -22}),
		},
		{
			name:         "should be able to set an int8 field",
			data:         &(struct{ IntField int8 }{}),
			field:        "IntField",
			value:        "22",
			expectedData: &(struct{ IntField int8 }{IntField: 22}),
		},
		{
			name:         "should be able to set a negative int8 field",
			data:         &(struct{ IntField int8 }{}),
			field:        "IntField",
			value:        "-22",
			expectedData: &(struct{ IntField int8 }{IntField: -22}),
		},
		{
			name:         "should be able to set an int16 field",
			data:         &(struct{ IntField int16 }{}),
			field:        "IntField",
			value:        "22",
			expectedData: &(struct{ IntField int16 }{IntField: 22}),
		},
		{
			name:         "should be able to set a negative int16 field",
			data:         &(struct{ IntField int16 }{}),
			field:        "IntField",
			value:        "-22",
			expectedData: &(struct{ IntField int16 }{IntField: -22}),
		},
		{
			name:         "should be able to set an int32 field",
			data:         &(struct{ IntField int32 }{}),
			field:        "IntField",
			value:        "22",
			expectedData: &(struct{ IntField int32 }{IntField: 22}),
		},
		{
			name:         "should be able to set a negative int32 field",
			data:         &(struct{ IntField int32 }{}),
			field:        "IntField",
			value:        "-22",
			expectedData: &(struct{ IntField int32 }{IntField: -22}),
		},
		{
			name:         "should be able to set an int64 field",
			data:         &(struct{ IntField int64 }{}),
			field:        "IntField",
			value:        "22",
			expectedData: &(struct{ IntField int64 }{IntField: 22}),
		},
		{
			name:         "should be able to set a negative int64 field",
			data:         &(struct{ IntField int64 }{}),
			field:        "IntField",
			value:        "-22",
			expectedData: &(struct{ IntField int64 }{IntField: -22}),
		},
		{
			name:         "should be able to set an uint field",
			data:         &(struct{ UintField uint }{}),
			field:        "UintField",
			value:        "22",
			expectedData: &(struct{ UintField uint }{UintField: 22}),
		},
		{
			name:         "should be able to set an uint8 field",
			data:         &(struct{ UintField uint8 }{}),
			field:        "UintField",
			value:        "22",
			expectedData: &(struct{ UintField uint8 }{UintField: 22}),
		},
		{
			name:         "should be able to set an uint16 field",
			data:         &(struct{ UintField uint16 }{}),
			field:        "UintField",
			value:        "22",
			expectedData: &(struct{ UintField uint16 }{UintField: 22}),
		},
		{
			name:         "should be able to set an uint32 field",
			data:         &(struct{ UintField uint32 }{}),
			field:        "UintField",
			value:        "22",
			expectedData: &(struct{ UintField uint32 }{UintField: 22}),
		},
		{
			name:         "should be able to set an uint64 field",
			data:         &(struct{ UintField uint64 }{}),
			field:        "UintField",
			value:        "22",
			expectedData: &(struct{ UintField uint64 }{UintField: 22}),
		},
		{
			name:         "should be able to set a float32 field",
			data:         &(struct{ FloatField float32 }{}),
			field:        "FloatField",
			value:        "22.7",
			expectedData: &(struct{ FloatField float32 }{FloatField: 22.7}),
		},
		{
			name:         "should be able to set a negative float32 field",
			data:         &(struct{ FloatField float32 }{}),
			field:        "FloatField",
			value:        "-22.7",
			expectedData: &(struct{ FloatField float32 }{FloatField: -22.7}),
		},
		{
			name:         "should be able to set a whole float32 field",
			data:         &(struct{ FloatField float32 }{}),
			field:        "FloatField",
			value:        "22",
			expectedData: &(struct{ FloatField float32 }{FloatField: 22}),
		},
		{
			name:         "should be able to set a negative whole float32 field",
			data:         &(struct{ FloatField float32 }{}),
			field:        "FloatField",
			value:        "-22",
			expectedData: &(struct{ FloatField float32 }{FloatField: -22}),
		},
		{
			name:         "should be able to set a float64 field",
			data:         &(struct{ FloatField float64 }{}),
			field:        "FloatField",
			value:        "22.7",
			expectedData: &(struct{ FloatField float64 }{FloatField: 22.7}),
		},
		{
			name:         "should be able to set a negative float64 field",
			data:         &(struct{ FloatField float64 }{}),
			field:        "FloatField",
			value:        "-22.7",
			expectedData: &(struct{ FloatField float64 }{FloatField: -22.7}),
		},
		{
			name:         "should be able to set a whole float64 field",
			data:         &(struct{ FloatField float64 }{}),
			field:        "FloatField",
			value:        "22",
			expectedData: &(struct{ FloatField float64 }{FloatField: 22}),
		},
		{
			name:         "should be able to set a negative whole float64 field",
			data:         &(struct{ FloatField float64 }{}),
			field:        "FloatField",
			value:        "-22",
			expectedData: &(struct{ FloatField float64 }{FloatField: -22}),
		},
		{
			name:  "should be able to set an complex64 field",
			data:  &(struct{ ComplexField complex64 }{}),
			field: "ComplexField",
			value: "3+4i",
			expectedData: &(struct{ ComplexField complex64 }{
				ComplexField: complex(float32(3), float32(4)),
			}),
		},
		{
			name:  "should be able to set an complex128 field",
			data:  &(struct{ ComplexField complex128 }{}),
			field: "ComplexField",
			value: "3+4i",
			expectedData: &(struct{ ComplexField complex128 }{
				ComplexField: complex(float64(3), float64(4)),
			}),
		},
		{
			name:         "should be able to set a string field",
			data:         &(struct{ StringField string }{}),
			field:        "StringField",
			value:        "example string",
			expectedData: &(struct{ StringField string }{StringField: "example string"}),
		},
		{
			name:         "should be able to set a bool field",
			data:         &(struct{ BoolField bool }{}),
			field:        "BoolField",
			value:        "true",
			expectedData: &(struct{ BoolField bool }{BoolField: true}),
		},
		{
			name:  "should be able to set a time field",
			data:  &(struct{ TimeField time.Time }{}),
			field: "TimeField",
			value: "2023-01-18T15:04:05Z",
			expectedData: &(struct{ TimeField time.Time }{
				TimeField: time.Date(2023, 1, 18, 15, 4, 5, 0, time.Now().UTC().Location()),
			}),
		},
		{
			name:        "should return an error if struct type is unsupported",
			data:        &(struct{ Unsupported struct{ name string } }{}),
			field:       "Unsupported",
			value:       "some value",
			expectedErr: newUnsupportedTypeError("Unsupported"),
		},
		{
			name:        "should return an error if int string is non numeric",
			data:        &(struct{ IntField int }{}),
			field:       "IntField",
			value:       "some value",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int string is fractional",
			data:        &(struct{ IntField int }{}),
			field:       "IntField",
			value:       "22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int string is fractional negative",
			data:        &(struct{ IntField int }{}),
			field:       "IntField",
			value:       "-22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int8 string is non numeric",
			data:        &(struct{ IntField int8 }{}),
			field:       "IntField",
			value:       "some value",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int8 string is fractional",
			data:        &(struct{ IntField int8 }{}),
			field:       "IntField",
			value:       "22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int8 string is fractional negative",
			data:        &(struct{ IntField int8 }{}),
			field:       "IntField",
			value:       "-22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int16 string is non numeric",
			data:        &(struct{ IntField int16 }{}),
			field:       "IntField",
			value:       "some value",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int16 string is fractional",
			data:        &(struct{ IntField int16 }{}),
			field:       "IntField",
			value:       "22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int16 string is fractional negative",
			data:        &(struct{ IntField int16 }{}),
			field:       "IntField",
			value:       "-22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int32 string is non numeric",
			data:        &(struct{ IntField int32 }{}),
			field:       "IntField",
			value:       "some value",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int32 string is fractional",
			data:        &(struct{ IntField int32 }{}),
			field:       "IntField",
			value:       "22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int32 string is fractional negative",
			data:        &(struct{ IntField int32 }{}),
			field:       "IntField",
			value:       "-22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int64 string is non numeric",
			data:        &(struct{ IntField int64 }{}),
			field:       "IntField",
			value:       "some value",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int64 string is fractional",
			data:        &(struct{ IntField int64 }{}),
			field:       "IntField",
			value:       "22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if int64 string is fractional negative",
			data:        &(struct{ IntField int }{}),
			field:       "IntField",
			value:       "-22.1",
			expectedErr: newInvalidIntError("IntField"),
		},
		{
			name:        "should return an error if uint string is non numeric",
			data:        &(struct{ UintField uint }{}),
			field:       "UintField",
			value:       "some value",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint string is negative",
			data:        &(struct{ UintField uint }{}),
			field:       "UintField",
			value:       "-22",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint string is fractional",
			data:        &(struct{ UintField uint }{}),
			field:       "UintField",
			value:       "22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint string is fractional negative",
			data:        &(struct{ UintField uint }{}),
			field:       "UintField",
			value:       "-22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint8 string is non numeric",
			data:        &(struct{ UintField uint8 }{}),
			field:       "UintField",
			value:       "some value",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint8 string is negative",
			data:        &(struct{ UintField uint8 }{}),
			field:       "UintField",
			value:       "-22",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint8 string is fractional",
			data:        &(struct{ UintField uint8 }{}),
			field:       "UintField",
			value:       "22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint8 string is fractional negative",
			data:        &(struct{ UintField uint8 }{}),
			field:       "UintField",
			value:       "-22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint16 string is non numeric",
			data:        &(struct{ UintField uint16 }{}),
			field:       "UintField",
			value:       "some value",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint16 string is negative",
			data:        &(struct{ UintField uint16 }{}),
			field:       "UintField",
			value:       "-22",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint16 string is fractional",
			data:        &(struct{ UintField uint16 }{}),
			field:       "UintField",
			value:       "22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint16 string is fractional negative",
			data:        &(struct{ UintField uint16 }{}),
			field:       "UintField",
			value:       "-22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint32 string is non numeric",
			data:        &(struct{ UintField uint32 }{}),
			field:       "UintField",
			value:       "some value",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint32 string is negative",
			data:        &(struct{ UintField uint32 }{}),
			field:       "UintField",
			value:       "-22",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint32 string is fractional",
			data:        &(struct{ UintField uint32 }{}),
			field:       "UintField",
			value:       "22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint32 string is fractional negative",
			data:        &(struct{ UintField uint32 }{}),
			field:       "UintField",
			value:       "-22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint64 string is non numeric",
			data:        &(struct{ UintField uint64 }{}),
			field:       "UintField",
			value:       "some value",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint64 string is negative",
			data:        &(struct{ UintField uint64 }{}),
			field:       "UintField",
			value:       "-22",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint64 string is fractional",
			data:        &(struct{ UintField uint64 }{}),
			field:       "UintField",
			value:       "22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if uint64 string is fractional negative",
			data:        &(struct{ UintField uint64 }{}),
			field:       "UintField",
			value:       "-22.1",
			expectedErr: newInvalidUintError("UintField"),
		},
		{
			name:        "should return an error if float32 string is non numeric",
			data:        &(struct{ FloatField float32 }{}),
			field:       "FloatField",
			value:       "some value",
			expectedErr: newInvalidFloatError("FloatField"),
		},
		{
			name:        "should return an error if float64 string is non numeric",
			data:        &(struct{ FloatField float64 }{}),
			field:       "FloatField",
			value:       "some value",
			expectedErr: newInvalidFloatError("FloatField"),
		},
		{
			name:        "should return an error if complex64 string is invalid",
			data:        &(struct{ ComplexField complex64 }{}),
			field:       "ComplexField",
			value:       "some value",
			expectedErr: newInvalidComplexError("ComplexField"),
		},
		{
			name:        "should return an error if complex128 string is invalid",
			data:        &(struct{ ComplexField complex128 }{}),
			field:       "ComplexField",
			value:       "some value",
			expectedErr: newInvalidComplexError("ComplexField"),
		},
		{
			name:        "should return an error if bool string is invalid",
			data:        &(struct{ BoolField bool }{}),
			field:       "BoolField",
			value:       "Not True",
			expectedErr: newInvalidBoolError("BoolField"),
		},
		{
			name:        "should return an error if time string is invalid",
			data:        &(struct{ TimeField time.Time }{}),
			field:       "TimeField",
			value:       "Not a datetime",
			expectedErr: newInvalidTimeError("TimeField"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			r, err := newReflector(tc.data)
			require.NoError(err, "Failed to create reflector")

			err = r.SetField(tc.field, tc.value)

			if tc.expectedErr == nil {
				require.NoError(err, "An unexpected error ocurred")
				assert.Equal(tc.expectedData, tc.data)
			} else {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

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
}

func Test_Reflect(t *testing.T) {
	testcases := []struct {
		name        string
		data        map[string]string
		s           any
		expectedS   any
		expectedErr *errors.GofidentialError
	}{
		{
			name:        "should return an error if data is nil",
			data:        nil,
			s:           nil,
			expectedErr: newInvalidEnvDataError(),
		},
		{
			name:        "should return an error if s is nil",
			data:        map[string]string{},
			s:           nil,
			expectedErr: newDataIsNilError(),
		},
		{
			name:        "should return an error if s is not a pointer",
			data:        map[string]string{},
			s:           1,
			expectedErr: newDataIsNotPtrError(),
		},
		{
			name: "should return an error if s is a nil pointer",
			data: map[string]string{},
			s: func() any {
				var data *struct{} = nil
				return data
			}(),
			expectedErr: newDataIsNilPtrError(),
		},
		{
			name:        "should return an error if s does not point to a struct",
			data:        map[string]string{},
			s:           &([]string{}),
			expectedErr: newDataIsNotStructError(),
		},
		{
			name:        "should return an error if s has unexported fields",
			data:        map[string]string{},
			s:           &(struct{ name string }{}),
			expectedErr: newUnexportedFieldError(),
		},
		{
			name: "should return an error if data has duplicate entries",
			data: map[string]string{
				"VALUE_": "value 1",
				"VALUE":  "value 1",
			},
			s:           &(struct{ Value string }{}),
			expectedErr: newDuplicateKeyError("Value"),
		},
		{
			name: "should return an error if s is missing a field",
			data: map[string]string{
				"NAME": "Felipe",
				"AGE":  "18",
			},
			s:           &(struct{ Name string }{}),
			expectedErr: newMissingFieldError("Age"),
		},
		{
			name: "should return an error if s has an unsupported type",
			data: map[string]string{
				"UNSUPPORTED": "unknown",
			},
			s:           &(struct{ Unsupported struct{ Name string } }{}),
			expectedErr: newUnsupportedTypeError("Unsupported"),
		},
		{
			name: "should return an error if data has a non numeric int",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int8",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int8",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int8",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int16",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int16",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int16",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int32",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int32",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int32",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int64",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int64",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int64",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},

		{
			name: "should return an error if data has a non numeric int",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int8",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int8",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int8",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int8 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int16",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int16",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int16",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int16 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int32",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int32",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int32",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int32 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric int64",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional int64",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative int64",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field int64 }{}),
			expectedErr: newInvalidIntError("Field"),
		},
		{
			name: "should return an error if data has a non numeric uint",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field uint }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a negative uint",
			data: map[string]string{
				"FIELD": "-22",
			},
			s:           &(struct{ Field uint }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional uint",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field uint }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative uint",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field uint }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a non numeric uint8",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field uint8 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a negative uint8",
			data: map[string]string{
				"FIELD": "-22",
			},
			s:           &(struct{ Field uint8 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional uint8",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field uint8 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative uint8",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field uint8 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a non numeric uint16",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field uint16 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a negative uint16",
			data: map[string]string{
				"FIELD": "-22",
			},
			s:           &(struct{ Field uint16 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional uint16",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field uint16 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative uint16",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field uint16 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a non numeric uint32",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field uint32 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a negative uint32",
			data: map[string]string{
				"FIELD": "-22",
			},
			s:           &(struct{ Field uint32 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional uint32",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field uint32 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative uint32",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field uint32 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a non numeric uint64",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field uint64 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a negative uint64",
			data: map[string]string{
				"FIELD": "-22",
			},
			s:           &(struct{ Field uint64 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional uint64",
			data: map[string]string{
				"FIELD": "22.1",
			},
			s:           &(struct{ Field uint64 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a fractional negative uint64",
			data: map[string]string{
				"FIELD": "-22.1",
			},
			s:           &(struct{ Field uint64 }{}),
			expectedErr: newInvalidUintError("Field"),
		},
		{
			name: "should return an error if data has a non numeric float32",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field float32 }{}),
			expectedErr: newInvalidFloatError("Field"),
		},
		{
			name: "should return an error if data has a non numeric float64",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field float64 }{}),
			expectedErr: newInvalidFloatError("Field"),
		},
		{
			name: "should return an error if data has an invalid complex64",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field complex64 }{}),
			expectedErr: newInvalidComplexError("Field"),
		},
		{
			name: "should return an error if data has an invalid complex128",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field complex128 }{}),
			expectedErr: newInvalidComplexError("Field"),
		},
		{
			name: "should return an error if data has an invalid bool",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field bool }{}),
			expectedErr: newInvalidBoolError("Field"),
		},
		{
			name: "should return an error if data has an invalid time",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field time.Time }{}),
			expectedErr: newInvalidTimeError("Field"),
		},
		{
			name: "should return an error if data has an invalid time",
			data: map[string]string{
				"FIELD": "some value",
			},
			s:           &(struct{ Field time.Time }{}),
			expectedErr: newInvalidTimeError("Field"),
		},
		{
			name: "should return an error if data has an invalid time",
			data: map[string]string{
				"APP_NAME":    "MyApp",
				"VERSION_NUM": "1",
				"DEBUG":       "true",
				"ENVIRONMENT": "production",
			},
			s: &(struct {
				AppName     string
				VersionNum  int
				Debug       bool
				Environment string
			}{}),
			expectedS: &(struct {
				AppName     string
				VersionNum  int
				Debug       bool
				Environment string
			}{
				AppName:     "MyApp",
				VersionNum:  1,
				Debug:       true,
				Environment: "production",
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			assert := assert.New(t)

			err := Reflect(tc.data, tc.s)

			if tc.expectedErr == nil {
				require.NoError(err, "An unexpected error ocurred")
				assert.Equal(tc.expectedS, tc.s)
			} else {
				require.Error(err, "An error was expected. But got none.")

				castErr, ok := err.(*errors.GofidentialError)
				require.True(ok, "Error is not of type GofidentialError.")

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
}
