package reflects

import "reflect"

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	return IsNilSafe(&v)
}

func IsNilN(i any) bool {
	if i == nil {
		return true
	}
	return IsNilS(&i)
}

func IsNilSafe(value *reflect.Value) bool {
	if value.IsValid() {
		return true
	}

	switch value.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}

func IsNilS(i any) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Array, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}