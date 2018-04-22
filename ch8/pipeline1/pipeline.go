package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	wait := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	go func() {
		for {
			fmt.Println(<-squares)
		}
	}()

	<-wait

}
