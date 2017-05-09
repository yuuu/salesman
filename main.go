package main

import "fmt"

func main() {
	csvMap := CsvMap{"map.csv"}
	navi, _ := csvMap.Build()
	route := navi.ShortestRoute()
	for i := 0; i < len(route.points); i++ {
		fmt.Println(route.points[i].Name())
	}
}
