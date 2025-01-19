package reflector

import (
	"reflect"
	"strconv"
	"time"
)

type reflector struct {
	ptrval reflect.Value
}

func (r *reflector) SetField(name string, value string) error {
	field := r.ptrval.FieldByName(name)

	if !field.IsValid() {
		return newMissingFieldError(name)
	}

	if !field.CanSet() {
		return newUnsettableFieldError(name)
	}

	fieldType := field.Type()

	switch fieldType.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return newInvalidIntError(name)
		}

		field.SetInt(intVal)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)

		if err != nil {
			return newInvalidUintError(name)
		}

		field.SetUint(uintVal)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return newInvalidFloatError(name)
		}

		field.SetFloat(floatVal)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		complexVal, err := strconv.ParseComplex(value, 128)

		if err != nil {
			return newInvalidComplexError(name)
		}

		field.SetComplex(complexVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)

		if err != nil {
			return newInvalidBoolError(name)
		}

		field.SetBool(boolVal)
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			parsedTime, err := time.Parse(time.RFC3339Nano, value)
			if err != nil {
				return newInvalidTimeError(name)
			}

			field.Set(reflect.ValueOf(parsedTime))
			break
		}

		fallthrough
	default:
		return newUnsupportedTypeError(name)
	}

	return nil
}

func (r *reflector) AllFields() (map[string]any, error) {
	// Get the struct type
	typ := r.ptrval.Type()

	// Get the reflect.Value of the struct
	val := r.ptrval

	fields := map[string]any{}

	for i := 0; i < typ.NumField(); i++ {
		// Get field metadata
		field := typ.Field(i)

		if !field.IsExported() {
			return nil, newUnexportedFieldError()
		}

		value := val.Field(i)
		fields[field.Name] = value.Interface()
	}

	return fields, nil
}

func newReflector(data any) (*reflector, error) {
	ptrval := reflect.ValueOf(data)

	// Checks if data is nil.
	if !ptrval.IsValid() {
		return nil, newSIsNilError()
	}

	// Check if data is not a pointer.
	if ptrval.Kind() != reflect.Pointer {
		return nil, newSIsNotPtrError()
	}

	// Dereference pointer.
	ptrval = ptrval.Elem()

	// Check if the data pointer value is not nil.
	if !ptrval.IsValid() {
		return nil, newNilStructPtrError()
	}

	// Check if the data pointer value is not a struct.
	if ptrval.Kind() != reflect.Struct {
		return nil, newInvalidStructPtrError()
	}

	return &reflector{
		ptrval: ptrval,
	}, nil
}

func Reflect(data map[string]string, s any) error {
	if data == nil {
		return newInvalidEnvDataError()
	}

	r, err := newReflector(s)

	if err != nil {
		return err
	}

	fields, err := r.AllFields()

	if err != nil {
		return err
	}

	newData := map[string]string{}
	originalKey := map[string]string{}

	for k, v := range data {
		newKey := upperToPascal(k)

		if _, ok := newData[newKey]; ok {
			return newDuplicateKeysError(k, originalKey[newKey])
		}

		if _, ok := fields[newKey]; !ok {
			return newMissingFieldError(newKey)
		}

		newData[newKey] = v
		originalKey[newKey] = k
	}

	for k, v := range newData {
		if err := r.SetField(k, v); err != nil {
			return err
		}
	}

	return nil
}
