package main

import (
	"errors"
	"fmt"
)

func main()  {
	e := errors.New("errorss")

	fmt.Println(e)
}
