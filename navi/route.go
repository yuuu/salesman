package salesman

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
