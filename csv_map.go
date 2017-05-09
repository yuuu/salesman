package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Map interface {
	Build()
}

type CsvMap struct {
	path string
}

func (self *CsvMap) buildPoints() ([]Point, error) {
	var fp *os.File
	var err error
	var points []Point
	var record []string

	fp, err = os.Open(self.path)
	if nil != err {
		return nil, err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true

	record, err = reader.Read()
	if nil != err {
		return nil, err
	}
	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		x, atoierr := strconv.Atoi(record[1])
		if nil != atoierr {
			return nil, atoierr
		}
		y, atoierr := strconv.Atoi(record[2])
		if nil != atoierr {
			return nil, atoierr
		}
		points = append(points, Point{float64(x), float64(y), record[0]})
	}

	return points, nil
}

func (self *CsvMap) buildPermutations(num int) [][]int {
	nums := make([]int, num)
	permutationNum := 1
	for i := 0; i < num; i++ {
		nums[i] = i + 1
		permutationNum *= (i + 1)
	}

	indexes := make([]int, num)
	permutations := make([][]int, permutationNum)
	for seed := 0; seed < permutationNum; seed++ {
		permutations[seed] = make([]int, num)
		tmpSeed := seed
		for i := 0; i < num; i++ {
			indexes[i] = tmpSeed % (i + 1)
			tmpSeed = tmpSeed / (i + 1)
		}
		tmpNums := make([]int, num)
		copy(tmpNums, nums)
		for i := num - 1; i >= 0; i-- {
			permutations[seed][num-1-i] = tmpNums[indexes[i]]
			if num > (indexes[i] + 1) {
				tmpNums = append(tmpNums[:indexes[i]], tmpNums[indexes[i]+1:]...)
			}
		}
	}
	return permutations
}

func (self *CsvMap) buildRoutes(points []Point, permutations [][]int) []Route {
	routes := make([]Route, len(permutations))
	for i := 0; i < len(permutations); i++ {
		routes[i].Add(&points[0])
		for j := 0; j < len(permutations[i]); j++ {
			routes[i].Add(&points[permutations[i][j]])
		}
		routes[i].Add(&points[0])
	}
	return routes
}

func (self *CsvMap) buildNavi(routes []Route) *Navi {
	navi := Navi{}
	for i := 0; i < len(routes); i++ {
		navi.Add(&routes[i])
	}
	return &navi
}

func (self *CsvMap) Build() (*Navi, error) {
	points, err := self.buildPoints()
	if nil != err {
		return nil, err
	}
	navi := self.buildNavi(self.buildRoutes(points, self.buildPermutations(len(points)-1)))

	return navi, nil
}
