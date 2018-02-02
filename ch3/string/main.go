package main

import (
	"fmt"
	"unicode/utf8"

)

func main()  {
	s := "hello, 世界"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Printf("% x\n", s)

	r := []rune(s)
	fmt.Println(len(r))
	fmt.Println(string(r[8]))
}
