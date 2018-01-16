package fake

import (
	"reflect"
)

func Fake(i interface{}) {
	v := reflect.ValueOf(&i)
	fake(v)

}

func fake(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	default:
		panic(v.Kind())
	case reflect.Invalid:
		setRandomInvalid(v)
	case reflect.Bool:
		setRandomBool(v)
	case reflect.Int:
		setRandomInt(v)
	case reflect.Int8:
		setRandomInt8(v)
	case reflect.Int16:
		setRandomInt16(v)
	case reflect.Int32:
		setRandomInt32(v)
	case reflect.Int64:
		setRandomInt64(v)
	case reflect.Uint:
		setRandomUint(v)
	case reflect.Uint8:
		setRandomUint8(v)
	case reflect.Uint16:
		setRandomUint16(v)
	case reflect.Uint32:
		setRandomUint32(v)
	case reflect.Uint64:
		setRandomUint64(v)
	case reflect.Uintptr:
		setRandomUintptr(v)
	case reflect.Float32:
		setRandomFloat32(v)
	case reflect.Float64:
		setRandomFloat64(v)
	case reflect.Complex64:
		setRandomComplex64(v)
	case reflect.Complex128:
		setRandomComplex128(v)
	case reflect.Array:
		setRandomArray(v)
	case reflect.Chan:
		setRandomChan(v)
	case reflect.Func:
		setRandomFunc(v)
	case reflect.Interface:
		setRandomInterface(v)
	case reflect.Map:
		setRandomMap(v)
	case reflect.Ptr:
		setRandomPtr(v)
	case reflect.Slice:
		setRandomSlice(v)
	case reflect.String:
		setRandomString(v)
	case reflect.Struct:
		setRandomStruct(v)
	case reflect.UnsafePointer:
		setRandomUnsafePointer(v)
	}

}

func setRandomInvalid(v reflect.Value) {
	panic(v.Kind())
}

func setRandomBool(v reflect.Value) {
	v.SetBool(true)
}

func setRandomInt(v reflect.Value) {
	v.SetInt(1234)
}

func setRandomInt8(v reflect.Value) {
	v.SetInt(1234)
}

func setRandomInt16(v reflect.Value) {
	v.SetInt(1234)
}

func setRandomInt32(v reflect.Value) {
	v.SetInt(1234)
}

func setRandomInt64(v reflect.Value) {
	v.SetInt(1234)
}

func setRandomUint(v reflect.Value) {
	var u uint = 1234
	v.Set(reflect.ValueOf(u))
}

func setRandomUint8(v reflect.Value) {
	var u uint8 = 123
	v.Set(reflect.ValueOf(u))
}

func setRandomUint16(v reflect.Value) {
	var u uint16 = 1234
	v.Set(reflect.ValueOf(u))
}

func setRandomUint32(v reflect.Value) {
	var u uint32 = 1234
	v.Set(reflect.ValueOf(u))
}

func setRandomUint64(v reflect.Value) {
	var u uint64 = 1234
	v.Set(reflect.ValueOf(u))
}

func setRandomUintptr(v reflect.Value) {
	var u uintptr = 8
	v.Set(reflect.ValueOf(u))
}

func setRandomFloat32(v reflect.Value) {
	v.SetFloat(1.234)
}

func setRandomFloat64(v reflect.Value) {
	v.SetFloat(1.234)
}

func setRandomComplex64(v reflect.Value) {
	v.SetComplex(64)
}

func setRandomComplex128(v reflect.Value) {
	v.SetComplex(128)
}

func setRandomArray(v reflect.Value) {

}

func setRandomChan(v reflect.Value) {

}

func setRandomFunc(v reflect.Value) {

}

func setRandomInterface(v reflect.Value) {
	fake(v.Elem())
}

func setRandomMap(v reflect.Value) {

}

func setRandomPtr(v reflect.Value) {
	v.Set(reflect.New(v.Type().Elem()))
	fake(v.Elem())
}

func setRandomSlice(v reflect.Value) {

}

func setRandomString(v reflect.Value) {
	v.SetString("hello world")
}

func setRandomStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr {
			setRandomPtr(field)
			continue
		}
		fake(field)
	}
}

func setRandomUnsafePointer(v reflect.Value) {

}
