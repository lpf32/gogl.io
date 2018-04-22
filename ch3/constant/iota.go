package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
)

func main() {
	fmt.Print(Sunday)
}
