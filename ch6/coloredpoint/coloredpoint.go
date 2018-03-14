package coloredpoint

import (
	"image/color"
	"gopl.io/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
