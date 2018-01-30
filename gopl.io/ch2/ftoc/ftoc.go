package main

import (
	"go-recipes/gopl.io/ch2/scope"
	"fmt"
	"os"
)


func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))
	scope.Print()
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
