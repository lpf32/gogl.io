package main

import (
	"flag"
	"fmt"
	"time"
)

type Rocket struct{}

func (r *Rocket) Launch() {
	fmt.Println("hello")
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
