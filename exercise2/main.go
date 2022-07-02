package main

import "fmt"

func main() {
	s := 100
	n := 2.5
	s = s * int(n)
	fmt.Println(s)

	a := 100
	b := 2.5

	fmt.Printf("speed     : %T\n", a)
	fmt.Printf("conversion: %T\n", float64(a))
	fmt.Printf("expression: %T\n", float64(a)*b)

	a = int(float64(a) * b)

	fmt.Println(a)

	var qwe int
	var asd int32

	qwe = int(asd)

	qwe = 65
	f := string(qwe)
	fmt.Println(f)

	fmt.Println(string([]byte{199, 105, 87, 34, 123}))

}
