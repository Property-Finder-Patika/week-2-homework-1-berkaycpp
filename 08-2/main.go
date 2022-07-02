package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`
	fmt.Println(json)

	name := os.Args[1]
	msg := `hi ` + name + `!
how are you?`
	fmt.Println(msg)

	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)

	msg = os.Args[1]
	l := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", l)
	fmt.Println(s)

	fmt.Println(strings.ToLower(os.Args[1]))

	msg = `
	The weather looks good.




					I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))

	name = "berkay           "

	name = strings.TrimRight(name, " ")
	l = utf8.RuneCountInString(name)
	fmt.Println(l)
}
