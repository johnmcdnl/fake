package fake

import "testing"
import "github.com/stretchr/testify/assert"

type TestingType struct {
	Child
	ChildNoPtr Child
	ChildPtr   *Child
}

type Child struct {
	Bool          bool
	Int           int
	Int8          int8
	Int16         int16
	Int32         int32
	Int64         int64
	Uint          uint
	Uint8         uint8
	Uint16        uint16
	Uint32        uint32
	Uint64        uint64
	Uintptr       uintptr
	Float32       float32
	Float64       float64
	Complex64     complex64
	Complex128    complex128
	String        string
	BoolPtr       *bool
	IntPtr        *int
	Int8Ptr       *int8
	Int16Ptr      *int16
	Int32Ptr      *int32
	Int64Ptr      *int64
	UintPtr       *uint
	Uint8Ptr      *uint8
	Uint16Ptr     *uint16
	Uint32Ptr     *uint32
	Uint64Ptr     *uint64
	UintptrPtr    *uintptr
	Float32Ptr    *float32
	Float64Ptr    *float64
	Complex64Ptr  *complex64
	Complex128Ptr *complex128
	StringPtr     *string
}

func TestFake(t *testing.T) {
	var testingType TestingType

	Fake(&testingType)
	assert.NotEmpty(t, testingType.Bool)
	assert.NotEmpty(t, testingType.Int)
	assert.NotEmpty(t, testingType.Int8)
	assert.NotEmpty(t, testingType.Int16)
	assert.NotEmpty(t, testingType.Int32)
	assert.NotEmpty(t, testingType.Int64)
	assert.NotEmpty(t, testingType.Uint)
	assert.NotEmpty(t, testingType.Uint8)
	assert.NotEmpty(t, testingType.Uint16)
	assert.NotEmpty(t, testingType.Uint32)
	assert.NotEmpty(t, testingType.Uint64)
	assert.NotEmpty(t, testingType.Uintptr)
	assert.NotEmpty(t, testingType.Float32)
	assert.NotEmpty(t, testingType.Float64)
	assert.NotEmpty(t, testingType.Complex64)
	assert.NotEmpty(t, testingType.Complex128)
	assert.NotEmpty(t, testingType.String)
	assert.NotEmpty(t, testingType.BoolPtr)
	assert.NotEmpty(t, testingType.IntPtr)
	assert.NotEmpty(t, testingType.Int8Ptr)
	assert.NotEmpty(t, testingType.Int16Ptr)
	assert.NotEmpty(t, testingType.Int32Ptr)
	assert.NotEmpty(t, testingType.Int64Ptr)
	assert.NotEmpty(t, testingType.UintPtr)
	assert.NotEmpty(t, testingType.Uint8Ptr)
	assert.NotEmpty(t, testingType.Uint16Ptr)
	assert.NotEmpty(t, testingType.Uint32Ptr)
	assert.NotEmpty(t, testingType.Uint64Ptr)
	assert.NotEmpty(t, testingType.UintptrPtr)
	assert.NotEmpty(t, testingType.Float32Ptr)
	assert.NotEmpty(t, testingType.Float64Ptr)
	assert.NotEmpty(t, testingType.Complex64Ptr)
	assert.NotEmpty(t, testingType.Complex128Ptr)
	assert.NotEmpty(t, testingType.StringPtr)

	assert.NotEmpty(t, testingType.ChildNoPtr.Bool)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int8)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int16)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int32)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int64)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint8)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint16)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint32)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint64)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uintptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Float32)
	assert.NotEmpty(t, testingType.ChildNoPtr.Float64)
	assert.NotEmpty(t, testingType.ChildNoPtr.Complex64)
	assert.NotEmpty(t, testingType.ChildNoPtr.Complex128)
	assert.NotEmpty(t, testingType.ChildNoPtr.String)
	assert.NotEmpty(t, testingType.ChildNoPtr.BoolPtr)
	assert.NotEmpty(t, testingType.ChildNoPtr.IntPtr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int8Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int16Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int32Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Int64Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.UintPtr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint8Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint16Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint32Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Uint64Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.UintptrPtr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Float32Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Float64Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Complex64Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.Complex128Ptr)
	assert.NotEmpty(t, testingType.ChildNoPtr.StringPtr)

	// assert.NotEmpty(t, testingType.ChildPtr.Bool)
	// assert.NotEmpty(t, testingType.ChildPtr.Int)
	// assert.NotEmpty(t, testingType.ChildPtr.Int8)
	// assert.NotEmpty(t, testingType.ChildPtr.Int16)
	// assert.NotEmpty(t, testingType.ChildPtr.Int32)
	// assert.NotEmpty(t, testingType.ChildPtr.Int64)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint8)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint16)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint32)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint64)
	// assert.NotEmpty(t, testingType.ChildPtr.Uintptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Float32)
	// assert.NotEmpty(t, testingType.ChildPtr.Float64)
	// assert.NotEmpty(t, testingType.ChildPtr.Complex64)
	// assert.NotEmpty(t, testingType.ChildPtr.Complex128)
	// assert.NotEmpty(t, testingType.ChildPtr.String)
	// assert.NotEmpty(t, testingType.ChildPtr.BoolPtr)
	// assert.NotEmpty(t, testingType.ChildPtr.IntPtr)
	// assert.NotEmpty(t, testingType.ChildPtr.Int8Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Int16Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Int32Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Int64Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.UintPtr)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint8Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint16Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint32Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Uint64Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.UintptrPtr)
	// assert.NotEmpty(t, testingType.ChildPtr.Float32Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Float64Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Complex64Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.Complex128Ptr)
	// assert.NotEmpty(t, testingType.ChildPtr.StringPtr)
}
