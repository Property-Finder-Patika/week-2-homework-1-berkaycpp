package main

import (
	"fmt"
	"path"
)

func main() {
	color := "green"
	color = "blue"
	fmt.Println(color)

	color = "dark " + color
	fmt.Println(color)

	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	_, b := multi()
	fmt.Println(b)

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

func multi() (int, int) {
	return 5, 4
}
