package reflector

import "reflect"

type Reflector struct {
	ptrval reflect.Value
}

func (r *Reflector) IsEmpty() bool {
	return false
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

	// Dereference pointer.
	ptrval = ptrval.Elem()

	// Check if the data pointer value is not nil.
	if !ptrval.IsValid() {
		return nil, newDataIsNilPtrError()
	}

	// Check if the data pointer value is not a struct.
	if ptrval.Kind() != reflect.Struct {
		return nil, newDataIsNotStructError()
	}

	return &Reflector{
		ptrval: ptrval,
	}, nil
}
