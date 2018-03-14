package main

import (
	"fmt"
	"time"
)

func main()  {
	naturals := make(chan int)
	squares := make(chan int)
	wait := make(chan int)

	go func() {
		for x := 0; x < 4 ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break //
			}
			squares <- x * x
		}
		close(squares)
	}()

	go func() {
		for x := range squares {
			fmt.Println(x)
		}
		wait <- 1
	}()

	<- wait

}
