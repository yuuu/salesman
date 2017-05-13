package main

import "testing"
import "fmt"

func Benchmark_all(b *testing.B) {
	csvMap := CsvMap{"map.csv"}
	navi, _ := csvMap.Build()
	route := navi.ShortestRoute()
	for i := 0; i < len(route.points); i++ {
		fmt.Println(route.points[i].Name())
	}
}
