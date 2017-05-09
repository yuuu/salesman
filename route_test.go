package main

import "testing"

func TestRoute_Distance(t *testing.T) {
	var route Route

	if 0.0 != route.Distance() {
		t.Fatal("route.distance is not initialized.")
	}
}

func TestRoute_Add(t *testing.T) {
	var route Route

	point1 := Point{1.0, 2.0, "Start"}
	point2 := Point{3.0, 4.0, "Point"}
	point3 := Point{5.0, 6.0, "Goal"}
	route.Add(&point1).Add(&point2).Add(&point3)

	if route.points[0] != point1 {
		t.Fatal("route.points[0] is not point1.")
	}
	if route.points[1] != point2 {
		t.Fatal("route.points[1] is not point2.")
	}
	if route.points[2] != point3 {
		t.Fatal("route.points[2] is not point3.")
	}
	if 4.0 != route.Distance() {
		t.Fatal("route.distance is incorrect.")
	}
}
