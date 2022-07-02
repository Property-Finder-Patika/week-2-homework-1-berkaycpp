package main

import (
	"fmt"
)

func main() {
	var height int
	fmt.Println(height)

	var isOn bool
	fmt.Println(isOn)

	var brightness float64
	fmt.Println(brightness)

	var s string
	fmt.Printf("s (%T): %q\n", s, s)

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var b bool
	var s2 string
	var r rune
	var by byte
	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)
	fmt.Printf("%q\n", s2)

	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	var firstName, lastName string = "asd", "xyz"
	fmt.Printf("%q %q\n", firstName, lastName)
}
