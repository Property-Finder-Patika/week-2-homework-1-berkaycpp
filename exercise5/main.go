package main

import "fmt"

var asd = 10

const constvar = true

func main() {
	var hello = "Hello"

	fmt.Println(hello, constvar)

	fmt.Println("inside main:", asd)

	nested()

	fmt.Println("inside main:", asd)

}

func nested() {
	var asd = 5
	fmt.Println("\nnested:", asd)

}
