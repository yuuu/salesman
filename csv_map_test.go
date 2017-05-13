package main

import (
	"testing"
)

func TestCsvMap_buildPoints(t *testing.T) {
	var csvMap = CsvMap{"map.csv"}
	testpoints := []Point{
		Point{939.0, 256.0, "start"},
		Point{251.0, 702.0, "1"},
		Point{143.0, 625.0, "2"},
		Point{313.0, 264.0, "3"},
		Point{487.0, 905.0, "4"},
		Point{493.0, 357.0, "5"},
		Point{410.0, 834.0, "6"},
		Point{572.0, 726.0, "7"},
		Point{281.0, 700.0, "8"},
		Point{638.0, 345.0, "9"},
		Point{964.0, 241.0, "10"},
	}

	points, err := csvMap.buildPoints()
	if nil != err {
		t.Fatal("csvMap.buildPoint() is failed.")
	}
	for i := 0; i < len(testpoints); i++ {
		if testpoints[i].Name() != points[i].Name() {
			t.Fatal("points[", i, "].Name() is incorrect.")
		}
		if testpoints[i].X() != points[i].X() {
			t.Fatal("points[", i, "].X() is incorrect.")
		}
		if testpoints[i].Y() != points[i].Y() {
			t.Fatal("points[", i, "].Y() is incorrect.")
		}
	}
}

func TestCsvMap_buildPermutations(t *testing.T) {
	var csvMap = CsvMap{"map.csv"}
	testPermutations := [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 2, 3},
		{1, 4, 3, 2},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 1, 3},
		{2, 4, 3, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 1, 2, 3},
		{4, 1, 3, 2},
		{4, 2, 1, 3},
		{4, 2, 3, 1},
		{4, 3, 1, 2},
		{4, 3, 2, 1},
	}

	permutations := csvMap.buildPermutations(4)
	for i := 0; i < len(permutations); i++ {
		for j := 0; j < len(permutations[i]); j++ {
			if testPermutations[i][j] != permutations[i][j] {
				t.Fatal("permutations[", i, "][", j, "] is incorrect.")
			}
		}
	}
}
