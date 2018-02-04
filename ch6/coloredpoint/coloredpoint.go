package coloredpoint

import (
	"image/color"
	"awesomeProject/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
