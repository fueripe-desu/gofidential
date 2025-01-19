package reflector

const (
	// Occurs when the 'data' parameter is nil.
	SIsNil string = "REFLECTOR_S_IS_NIL"

	// Occurs when the provided 'data' parameter is not a pointer.
	SIsNotPtr string = "REFLECTOR_S_IS_NOT_PTR"

	// Occurs when the provided 'data' parameter is a nil pointer.
	NilStructPtr string = "REFLECTOR_NIL_STRUCT_PTR"

	// Occurs when the .env file data is not passed correctly.
	InvalidEnvData string = "REFLECTOR_INVALID_ENV_DATA"

	// Occurs when the provided 'data' pointer does not point to a struct.
	InvalidStructPtr string = "REFLECTOR_INVALID_STRUCT_PTR"

	// Occurs when the .env file has a duplicate key (maybe due to a bug).
	DuplicateKeys string = "REFLECTOR_DUPLICATE_KEYS"

	// Occurs when the provided struct is missing a require field.
	MissingField string = "REFLECTOR_MISSING_FIELD"

	// Occurs when a field of the provided struct has an unsupported type.
	UnsupportedType string = "REFLECTOR_UNSUPPORTED_TYPE"

	// Occurs when the field of the provided struct is not settable.
	UnsettableField string = "REFLECTOR_UNSETTABLE_FIELD"

	// Occurs when the field of the provided struct is not exported.
	UnexportedField string = "REFLECTOR_UNEXPORTED_FIELD"

	// Occurs when the value is not a valid int in a .env file.
	InvalidInt string = "REFLECTOR_INVALID_INT"

	// Occurs when the value is not a valid uint in a .env file.
	InvalidUint string = "REFLECTOR_INVALID_UINT"

	// Occurs when the value is not a valid float in a .env file.
	InvalidFloat string = "REFLECTOR_INVALID_FLOAT"

	// Occurs when the value is not a valid bool in a .env file.
	InvalidBool string = "REFLECTOR_INVALID_BOOL"

	// Occurs when the value is not a valid complex number in a .env file.
	InvalidComplex string = "REFLECTOR_INVALID_COMPLEX"

	// Occurs when the value is not a valid time in a .env file.
	InvalidTime string = "REFLECTOR_INVALID_TIME"
)
