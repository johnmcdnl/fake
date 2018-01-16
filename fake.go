package fake

import (
	"reflect"
)

func Fake(i interface{}) {

}

func fake(v reflect.Value) {
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

}

func setRandomBool(v reflect.Value) {

}

func setRandomInt(v reflect.Value) {

}

func setRandomInt8(v reflect.Value) {

}

func setRandomInt16(v reflect.Value) {

}

func setRandomInt32(v reflect.Value) {

}

func setRandomInt64(v reflect.Value) {

}

func setRandomUint(v reflect.Value) {

}

func setRandomUint8(v reflect.Value) {

}

func setRandomUint16(v reflect.Value) {

}

func setRandomUint32(v reflect.Value) {

}

func setRandomUint64(v reflect.Value) {

}

func setRandomUintptr(v reflect.Value) {

}

func setRandomFloat32(v reflect.Value) {

}

func setRandomFloat64(v reflect.Value) {

}

func setRandomComplex64(v reflect.Value) {

}

func setRandomComplex128(v reflect.Value) {

}

func setRandomArray(v reflect.Value) {

}

func setRandomChan(v reflect.Value) {

}

func setRandomFunc(v reflect.Value) {

}

func setRandomInterface(v reflect.Value) {

}

func setRandomMap(v reflect.Value) {

}

func setRandomPtr(v reflect.Value) {

}

func setRandomSlice(v reflect.Value) {

}

func setRandomString(v reflect.Value) {

}

func setRandomStruct(v reflect.Value) {

}

func setRandomUnsafePointer(v reflect.Value) {

}
