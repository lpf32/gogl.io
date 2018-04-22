package coloredpoint

import (
	"gopl.io/ch6/geometry"
	"image/color"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
