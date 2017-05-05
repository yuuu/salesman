package salesman

import (
	"math"
	"testing"
)

func TestPoint_Property(t *testing.T) {
	name := "Start"
	x, y := 1.0, 2.0
	point := Point{x, y, name}
	if name != point.Name() {
		t.Fatal("point.Name() is not \"Start\".")
	}
	if x != point.X() {
		t.Fatal("point.X() is not", x, ".")
	}
	if y != point.Y() {
		t.Fatal("point.Y() is not", y, ".")
	}
}

func TestPoint_Distance(t *testing.T) {
	point1 := Point{1, 2, "point1"}
	point2 := Point{3, 3, "point2"}
	distance := math.Cbrt((3-1)*(3-1) + (3-2)*(3-2))
	if distance != point1.Distance(&point2) {
		t.Fatal("Distance from point1 to point2 is incorrect.")
	}
}
