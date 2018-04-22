package main

import "fmt"

func p(i int) {
	fmt.Println(i)
}

func reentrant(c chan<- int) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- i
		}(i)
	}
}

func variable() (i int) {
	i = 1
	return i
}

func main() {
	c := make(chan int, 10)
	reentrant(c)
	for i := 0; i < 10; i++ {
		<-c
	}
	//c <- 1
	//c <- 2
	//c <- 3
	//
	//close(c)
	//
	//for i := range c {
	//	fmt.Println(i)
	//}
	//
	//fmt.Println(<-c)
}
