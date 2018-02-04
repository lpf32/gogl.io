package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n*n
}

func add1(r rune) rune {
	return r + 1
}

func product(m, n int) int  {
	return m * n
}
func main()  {

	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "abc"))
}
