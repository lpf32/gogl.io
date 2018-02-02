package main

import (
	"fmt"
	"math"
)

func main()  {
	nan := math.NaN()

	fmt.Println(math.IsNaN(nan))
}
