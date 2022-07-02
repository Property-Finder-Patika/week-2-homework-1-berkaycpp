package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("number of logical CPUs usable by the current process:", runtime.NumCPU())
}

//go doc runtime NumCPU
