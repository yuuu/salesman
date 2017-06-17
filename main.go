package main

func main() {
	csvMap := CsvMap{"map.csv"}
	navi, _ := csvMap.Build()
	route := navi.ShortestRoute()
	route.Print()
}
