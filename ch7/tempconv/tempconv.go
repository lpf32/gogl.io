package tempconv

import (
	ch2temp "gopl.io/ch2/tempconv"
	"fmt"
	"flag"
)

type celsiusFlag struct { ch2temp.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = ch2temp.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = ch2temp.Celsius(ch2temp.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}


func CelsiusFlag(name string, value ch2temp.Celsius, usage string) *ch2temp.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
