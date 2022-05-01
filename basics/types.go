package main

func types() {
	// Boolean
	var _ bool // a boolean, true or false

	// String
	var _ string // a string, UTF-8 encoded by default

	// Numeric types
	var _ uint    // either 32 or 64 bits unsigned integer
	var _ uint8   // (0 to 255)
	var _ uint16  // (0 to 65535)
	var _ uint32  // (0 to 4294967295)
	var _ uint64  // (0 to 18446744073709551615)
	var _ uintptr // an unsigned integer large enough to store the uninterpreted bits of a pointer value

	var _ int   // either 32 or 64 bits integer
	var _ int8  // (-128 to 127)
	var _ int16 // (-32768 to 32767)
	var _ int32 // (-2147483648 to 2147483647)
	var _ int64 // (-9223372036854775808 to 9223372036854775807)

	var _ float32 // IEEE-754 32-bit floating-point numbers
	var _ float64 // IEEE-754 64-bit floating-point numbers

	var _ complex64  // complex numbers with float32 real and imaginary parts
	var _ complex128 // complex numbers with float64 real and imaginary parts

	var _ byte // alias for uint8
	var _ rune // alias for int32, used to emphasize than an integer represents a Unicode code point.

	// Arrays
	var _ [5]float32 // array of 5 float32s
	var _ [32]byte
	var _ [2 * 5]struct{ x, y int32 }
	var _ [1000]*float64
	var _ [3][5]int
	var _ [2][2][2]float64 // same as [2]([2]([2]float64))

	// Slice type
	var _ []float32 // descriptor for a contiguous segment of an underlying array

	// Struct
	var _ struct{}
	var _ struct {
		x, y int
		u    float32
		_    float32 // padding
		A    *[]int
		F    func()
	}

	// Struct with `embedded fields`
	type T1 string
	type T2 int
	var _ struct {
		x, y int
		T1   // field name is T1
		*T2  // field name is T2
	}

	// Pointer
	type Point struct{ x, y int }
	var _ *Point
	var _ *[4]int

	//  The value of an uninitialized variable of function type is nil

}
