package main

import "fmt"

func main() {
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)
	fmt.Println(C0, C1, C2)

	const (
		x1 = iota*3 + 10
		_
		x12
		x13
	)
	fmt.Println(x1, x12, x13)

}
