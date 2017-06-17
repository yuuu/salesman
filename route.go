package main

import "fmt"

type Route struct {
	points   []Point
	distance float64
}

func (self *Route) Distance() float64 {
	return self.distance
}

func (self *Route) Add(point *Point) *Route {
	len := len(self.points)
	self.points = append(self.points, *point)
	if 0 < len {
		self.distance += self.points[len-1].Distance(point)
	}
	return self
}

func (self *Route) Print() {
	fmt.Println("Shortest Route is...")
	for i := 0 ; i < len(self.points) - 1 ; i++ {
		fmt.Print(self.points[i].Name())
		fmt.Print(" -> ")
	}
	fmt.Println(self.points[len(self.points) - 1].Name())
}
