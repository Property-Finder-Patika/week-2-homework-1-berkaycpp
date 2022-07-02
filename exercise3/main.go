package main

import (
	"fmt"
	"strconv"
)

func main() {

	b, err1 := strconv.ParseBool("true")
	f, err2 := strconv.ParseFloat("3.1415", 64)
	i, err3 := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u, err1, err2, err3, err)

	q := strconv.Quote("Hello, 世界")
	fmt.Println(q)
	q = strconv.QuoteToASCII("Hello, 世asd界")
	fmt.Println(q)

	x := "245"
	y, e := strconv.Atoi(x)
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
}
