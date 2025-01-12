package errors

const (
	// Occurs when the 'data' parameter is nil.
	DataIsNilCode string = "DATA_IS_NIL"

	// Occurs when the provided 'data' parameter is not a pointer.
	DataIsNotPtrCode string = "DATA_IS_NOT_PTR"

	// Occurs when the provided 'data' parameter is a nil pointer.
	DataIsNilPtrCode string = "DATA_IS_NIL_PTR"

	// Occurs when the provided 'data' pointer does not point to a struct.
	DataIsNotStructCode string = "DATA_IS_NOT_STRUCT"
)
