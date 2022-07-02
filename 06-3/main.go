package main

import (
	"fmt"
)

func main() {

	i := 314
	f := 3.14
	s := "hello"
	b := true
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)

	a, b := 14, true
	fmt.Println(a, b)

	sum := 27 + 3.5

	fmt.Println(sum)

	/*
		b, _:= 10, 20
		fmt.Println(b)
	*/

	age, yourAge := 10, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)

}
