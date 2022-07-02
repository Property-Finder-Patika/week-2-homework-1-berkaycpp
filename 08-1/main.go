package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))

	x := 5. / 2
	fmt.Println(x)

	fmt.Println(10 + 5 - (5 - 10))
	fmt.Println(-10 + 0.5 - (1 + 5.5))
	fmt.Println(5 + 10*(2-5))
	fmt.Println(0.5 * (2 - 1))
	fmt.Println((3+1)/2*10 + 4)
	fmt.Println(10 / 2 * (10 % 7))
	fmt.Println(100 / (5. / 2))

	counter, factor := 45, 0.5
	counter++
	counter++
	counter++
	counter++
	counter++
	factor--
	factor--
	fmt.Println(float64(counter) * factor)
	counter++
	counter--
	counter += 5
	counter *= 10
	counter /= 5
	fmt.Println(counter)

	width, height := 10, 2
	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5
	fmt.Println(width)

	var (
		radius = 10.
		area   float64
	)
	area = math.Pi * radius * radius
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	area = 4 * math.Pi * math.Pow(radius, 2)
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	var vol float64
	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
