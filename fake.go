package fake

import (
	"math"
	"math/rand"
	"reflect"
	"time"
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

var src = rand.NewSource(time.Now().UnixNano())

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setRandomInvalid(v reflect.Value) {
	panic(v.Kind())
}

func setRandomBool(v reflect.Value) {
	v.SetBool(true)
}

func setRandomInt(v reflect.Value) {
	v.SetInt(randomInt(1, math.MaxInt64))
}

func setRandomInt8(v reflect.Value) {
	v.SetInt(randomInt(1, math.MaxInt8))
}

func setRandomInt16(v reflect.Value) {
	v.SetInt(randomInt(1, math.MaxInt16))
}

func setRandomInt32(v reflect.Value) {
	v.SetInt(randomInt(1, math.MaxInt32))
}

func setRandomInt64(v reflect.Value) {
	v.SetInt(randomInt(1, math.MaxInt64))
}

func setRandomUint(v reflect.Value) {
	u := uint(randomInt(1, math.MaxInt64))
	v.Set(reflect.ValueOf(u))
}

func setRandomUint8(v reflect.Value) {
	u := uint8(randomInt(1, math.MaxUint8))
	v.Set(reflect.ValueOf(u))
}

func setRandomUint16(v reflect.Value) {
	u := uint16(randomInt(1, math.MaxUint16))
	v.Set(reflect.ValueOf(u))
}

func setRandomUint32(v reflect.Value) {
	u := uint32(randomInt(1, math.MaxUint32))
	v.Set(reflect.ValueOf(u))
}

func setRandomUint64(v reflect.Value) {
	u := uint64(randomInt(1, math.MaxInt64))
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
	n := 25
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	v.SetString(string(b))
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

func randomInt(min, max int) int64 {
	n := rand.Intn(max)
	//TODO fix this up
	if n <= min {
		randomInt(min, max)
	}
	return int64(n)
}
