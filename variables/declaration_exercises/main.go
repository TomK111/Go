package main

import "fmt"

func main() {
	var (
		speed     = 100
		prevSpeed = 50
	)

	speed, prevSpeed = prevSpeed, speed
	fmt.Println(speed, prevSpeed)

	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// float types
	var f32 float32
	var f64 float64

	// complex types
	var c64 complex64
	var c128 complex128

	// bool
	var b bool

	//string types
	var s string
	var r rune
	var by byte

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	fmt.Printf("%q\n", s)
}
