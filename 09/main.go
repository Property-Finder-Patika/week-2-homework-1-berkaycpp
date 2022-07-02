package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// an english letter (search web for: ascii control code)
	var letter byte = 'X'
	fmt.Println("an english letter:", letter)

	// a non-english letter (search web for: unicode codepoint)
	var unicode rune
	unicode = '$'
	fmt.Println("a non-english letter:", unicode)

	// a year in 4 digits like 2040
	var year uint16 = 2345
	fmt.Println("a year in 4 digits like 2040:", year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 123
	fmt.Println("a month in 2 digits: 1 to 12:", month)

	// the speed of the light
	var lightSpeed uint32 = 670616629 // miles
	fmt.Println("the speed of the light:", lightSpeed)

	// angle of a circle
	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)

	var pi float64
	pi = 0.123123123123123123123
	fmt.Println("PI number:", pi)

	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)

	val, _ := strconv.ParseInt(os.Args[1], 5, 6)
	fmt.Println("int8 value is :", int8(val))

	// 2nd argument is an int16
	val, _ = strconv.ParseInt(os.Args[2], 2, 12)
	fmt.Println("int16 value is:", int16(val))

	// 3rd argument is an int32
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))

	// 4th argument is an int64
	// Remember ParseInt returns an int64
	val, _ = strconv.ParseInt(os.Args[4], 8, 64)
	fmt.Println("int64 value is:", val)

	// 5th argument is a number in bits. And its int8.
	// ParseInt(.., 2, ...) -> 2 means binary base
	val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[5], int8(val))

	t, _ := time.ParseDuration("1h30m")

	// 1. get the first command-line argument
	// 2. convert it to int64
	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)

	// 3. multiply the time duration with the given argument
	//
	//    converts the int64 value to time.Duration to be
	//    able to multiply it with the time.Duration value
	t *= time.Duration(multiplier)

	// 4. print it
	fmt.Println(t)

	type (
		Feet   float64
		Meters float64
	)

	var (
		feet   Feet
		meters Meters
	)

	arg := os.Args[1]

	// feet is a Feet value now
	valx, _ := strconv.ParseFloat(arg, 64)
	feet = Feet(valx)

	// meters is a Meters value now
	meters = Meters(feet * 0.3048)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)

	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)
}
