package main

import "fmt"

func main() {
	var ratio float64 = 3 / 2
	fmt.Println(ratio)
	ratio = float64(int(7) / int(3))
	fmt.Println(ratio)

	ratio = float64(8) / 3
	fmt.Println(ratio)

	ratio = 3.0 / 2
	fmt.Println(ratio)
	fmt.Println("sum :", 5+1)
	fmt.Println("sum :", 22+5.6231)
	fmt.Println("dif :", -1-0.5)
	fmt.Println("prod:", 4*-5)
	fmt.Println("prod:", 0.65*2.5)
	fmt.Println("quot:", 4/2)
	fmt.Println("quot:", 8/1.5)
	fmt.Println("remainder :", 8%3)

	f := 8.0
	fmt.Println("remainder :", int(f)%3)
}
