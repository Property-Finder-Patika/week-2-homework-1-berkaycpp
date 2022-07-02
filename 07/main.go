package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("I'm %d years old.\n", 24)

	fmt.Printf("My name is %s and my lastname is %s.\n", "Berkay", "Sahin")
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Berkay", "Sahin")

	tf := false
	fmt.Printf("These are %t claims.\n", tf)

	fmt.Printf("Temperature is %.1f degrees.\n", 35.7)

	fmt.Printf("%q\n", "hello world")

	fmt.Printf("Type of %d is %[1]T\n", 3)

	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)

	fmt.Printf("Type of %s is %[1]T\n", "hello")

	fmt.Printf("Type of %t is %[1]T\n", true)

	name, lastname := os.Args[1], os.Args[2]
	msg = "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, name, lastname)
}
