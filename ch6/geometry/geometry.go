package geometry

import "math"

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}


func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p *Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) InMap(m map[int]int) int {
	return 0
}

//func (p Point) inMap(q Point) float64 {
//	return math.Hypot(q.X - p.X, q.Y - p.Y)
//}

type Path []Point

func (path Path) Distance () float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (path Path) TranslateBy(offset Point, add bool)  {
	var op func(q, p Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int  {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
