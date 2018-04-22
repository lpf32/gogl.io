package main

import "fmt"

func change(s string) {
	s = "345"
}

func main() {
	s := "123"

	change(s[:])
	fmt.Println(s[1:2])
}
