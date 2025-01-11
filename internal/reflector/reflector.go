package reflector

import "reflect"

type Reflector struct {
	ptrval reflect.Value
}

func New(data any) (*Reflector, error) {
	ptrval := reflect.ValueOf(data)

	// Checks if data is nil.
	if !ptrval.IsValid() {
		return nil, newDataIsNilError()
	}

	// Check if data is not a pointer.
	if ptrval.Kind() != reflect.Pointer {
		return nil, newDataIsNotPtrError()
	}

	// Check if the data pointer value is not nil.
	if !ptrval.Elem().IsValid() {
		return nil, newDataIsNilPtrError()
	}

	// Check if the data pointer value is not a struct.
	if ptrval.Elem().Kind() != reflect.Struct {
		return nil, newDataIsNotStructError()
	}

	return &Reflector{
		ptrval: ptrval,
	}, nil
}
