package reflector

import (
	"fmt"

	"github.com/fueripe-desu/gofidential/errors"
	errorCode "github.com/fueripe-desu/gofidential/errors/reflector"
)

func newSIsNilError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.SIsNil,
		Message: "The parameter 's' must not be nil.",
		Hint:    "Ensure that a valid struct pointer is passed as the 's' parameter to the Load() function.",
	}
}

func newSIsNotPtrError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.SIsNotPtr,
		Message: "The parameter 's' must be a pointer.",
		Hint:    "Ensure that a valid struct pointer is passed as the 's' parameter to the Load() function.",
	}
}

func newNilStructPtrError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.NilStructPtr,
		Message: "The pointer passed to the 's' parameter must not point nil.",
		Hint:    "Ensure that a valid struct pointer is passed as the 's' parameter to the Load() function.",
	}
}

func newInvalidStructPtrError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidStructPtr,
		Message: "The pointer passed to the 's' parameter must point to a valid struct.",
		Hint:    "Ensure that a valid struct pointer is passed as the 's' parameter to the Load() function.",
	}
}

func newDuplicateKeysError(key1 string, key2 string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.DuplicateKeys,
		Message: fmt.Sprintf("The key '%s' conflicts with the key '%s'. Please use a different name.", key1, key2),
		Hint:    "This is an internal error. If you can’t find a solution it, please open an issue at https://github.com/fueripe-desu/gofidential.",
	}
}

func newMissingFieldError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.MissingField,
		Message: fmt.Sprintf("The expected field '%s' is missing in the struct.", field),
		Hint:    "Add the missing field in the struct, or remove the unused key from the .env file.",
	}
}

func newUnsupportedTypeError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnsupportedType,
		Message: fmt.Sprintf("The field '%s' defined in the struct has an unsupported type.", field),
		Hint:    "Check the type of the '%s' field and replace it with a supported type, or remove the field if it’s not required.",
	}
}

func newInvalidIntError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidInt,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid integer string.", key),
		Hint:    "Ensure the value is a valid integer string. Examples: '22', '-13', '58'.",
	}
}

func newInvalidUintError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidUint,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid unsigned integer string.", key),
		Hint:    "Ensure the value is a valid unsigned integer string. Examples: '95', '16', '868' (negative values are not allowed).",
	}
}

func newInvalidFloatError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidFloat,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid float string.", key),
		Hint:    "Ensure the value is a valid float string. Examples: '3.1415', '2.1', '-1.7'.",
	}
}

func newInvalidComplexError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidComplex,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid complex number string.", key),
		Hint:    "Ensure the value is a valid complex number string. Example: '3+4i' (real and imaginary parts, separated by a plus or minus sign).",
	}
}

func newInvalidBoolError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidBool,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid boolean string.", key),
		Hint:    "Ensure the value is a valid boolean string. Examples: 'true' and 'false'.",
	}
}

func newInvalidTimeError(key string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidTime,
		Message: fmt.Sprintf("The key '%s' defined in the .env file does not contain a valid datetime string.", key),
		Hint:    "Ensure the value is a valid datetime string. Example: '2023-01-18T15:04:05Z'.",
	}
}

func newUnsettableFieldError(field string) *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnsettableField,
		Message: fmt.Sprintf("The field '%s' defined in the struct is not settable.", field),
		Hint:    "This is an internal error. If you can’t find a solution it, please open an issue at https://github.com/fueripe-desu/gofidential.",
	}
}

func newUnexportedFieldError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.UnexportedField,
		Message: "The provided struct contains unexported fields.",
		Hint:    "Ensure all fields in the struct start with an uppercase letter to make them exported.",
	}
}

func newInvalidEnvDataError() *errors.GofidentialError {
	return &errors.GofidentialError{
		Code:    errorCode.InvalidEnvData,
		Message: "Failed to read .env file data.",
		Hint:    "This is an internal error. If you can’t find a solution it, please open an issue at https://github.com/fueripe-desu/gofidential.",
	}
}
