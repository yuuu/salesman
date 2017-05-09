package main

type Navi struct {
	//routes        []Route
	shortestRoute *Route
}

func (self *Navi) Add(route *Route) *Navi {
	//self.routes = append(self.routes, *route)
	if nil == self.shortestRoute {
		self.shortestRoute = route
	} else {
		if self.shortestRoute.Distance() > route.Distance() {
			self.shortestRoute = route
		}
	}
	return self
}
func (self *Navi) ShortestRoute() *Route {
	return self.shortestRoute
}
