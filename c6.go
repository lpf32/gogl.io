package main

import (
	"fmt"
	"gopl.io/ch6/coloredpoint"
	"gopl.io/ch6/geometry"
	"image/color"
)

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint.ColoredPoint{geometry.Point{1, 1}, red}
	var q = coloredpoint.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)

	fmt.Println(p.Distance(q.Point))
}
