package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)

	fmt.Println(os.Args[0])

	fmt.Println(os.Args)

	fmt.Println(&(os.Args))
}
