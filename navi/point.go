package salesman

import "math"

type Point struct {
	x    float64
	y    float64
	name string
}

func (point *Point) Name() string {
	return point.name
}

func (point *Point) X() float64 {
	return point.x
}

func (point *Point) Y() float64 {
	return point.y
}

func (point *Point) Distance(other *Point) float64 {
	xdiff := (other.x - point.x)
	ydiff := (other.y - point.y)
	return (math.Cbrt(xdiff*xdiff + ydiff*ydiff))
}
