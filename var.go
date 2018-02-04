package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main () {
	var a Celsius = 1

	fmt.Println(a.String())
	fmt.Printf("%v\n", a)
	fmt.Printf("%s\n", a)
	fmt.Println(a)
	fmt.Println(float64(a))
}


func (c Celsius) String() string { return fmt.Sprintf("%g OC", c)}