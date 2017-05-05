package salesman

type Navi struct {
	routes []Route
}

func (self *Navi) Add(route *Route) *Navi {

}
func (self *Navi) ShortestRoute() *Route {
	shortestRoute := &self.routes[0]
	for i := 1; i < len(self.routes); i++ {
		if shortestRoute.Distance() < self.routes[i].Distance() {
			shortestRoute = &self.routes[i]
		}
	}
	return shortestRoute
}
