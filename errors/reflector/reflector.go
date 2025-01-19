package reflector

const (
	// Occurs when the 'data' parameter is nil.
	DataIsNilCode string = "DATA_IS_NIL"

	// Occurs when the .env file data is not passed correctly.
	InvalidEnvDataCode string = "INVALID_ENV_DATA"

	// Occurs when the provided 'data' parameter is not a pointer.
	DataIsNotPtrCode string = "DATA_IS_NOT_PTR"

	// Occurs when the provided 'data' parameter is a nil pointer.
	DataIsNilPtrCode string = "DATA_IS_NIL_PTR"

	// Occurs when the provided 'data' pointer does not point to a struct.
	DataIsNotStructCode string = "DATA_IS_NOT_STRUCT"

	// Occurs when the .env file has a duplicate key (maybe due to a bug).
	DuplicateKeyCode string = "DUPLICATE_KEY"

	// Occurs when the provided struct is missing a require field.
	MissingFieldCode string = "MISSING_FIELD"

	// Occurs when a field of the provided struct has an unsupported type.
	UnsupportedTypeCode string = "UNSUPPORTED_TYPE"

	// Occurs when the field of the provided struct is not settable.
	UnsettableFieldCode string = "UNSETTABLE_FIELD"

	// Occurs when the field of the provided struct is not exported.
	UnexportedFieldCode string = "UNEXPORTED_FIELD"

	// Occurs when the value is not a valid int in a .env file.
	InvalidIntCode string = "INVALID_INT"

	// Occurs when the value is not a valid uint in a .env file.
	InvalidUintCode string = "INVALID_UINT"

	// Occurs when the value is not a valid float in a .env file.
	InvalidFloatCode string = "INVALID_FLOAT"

	// Occurs when the value is not a valid bool in a .env file.
	InvalidBoolCode string = "INVALID_BOOL"

	// Occurs when the value is not a valid complex number in a .env file.
	InvalidComplexCode string = "INVALID_COMPLEX"

	// Occurs when the value is not a valid time in a .env file.
	InvalidTimeCode string = "INVALID_TIME"
)
