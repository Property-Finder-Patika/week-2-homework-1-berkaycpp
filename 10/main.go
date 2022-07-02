package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)

	fmt.Printf("There are %d minutes in %d weeks.\n",
		minsPerDay*weekDays*2, 2)

	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)

	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)

	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n",
		length, home)

	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)

	fmt.Printf("tau = %g\n", tau)

	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)

	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)

	const (
		Nov = 11 - iota // 11 - 0 = 11
		Oct             // 11 - 1 = 10
		Sep             // 11 - 2 = 9
	)

	fmt.Println(Sep, Oct, Nov)

	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
