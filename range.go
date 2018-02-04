package main

import "fmt"

func main()  {
	ages := make(map[string]int)

	age, ok := ages["zhang"]
	fmt.Println(age, ok)
}
