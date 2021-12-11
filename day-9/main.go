package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-9/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	floorMap := readFloorMap(file)

	fmt.Println("Part 1")
	fmt.Printf("Sum of low point risk levels: %d\n", sumOfRiskLevels(floorMap.lowPoints()))

	fmt.Println("Part 2")
	fmt.Printf("Multiplied size of three largest basins: %d\n", productOfNLargestBasins(floorMap.basins(), 3))
}

type floorMap [][]int

type lowPoint struct {
	caveIndex     int
	locationIndex int
	value         int
}

func (fm floorMap) lowPoints() []lowPoint {
	lowPoints := make([]lowPoint, 0)
	for c, cave := range fm {
		for l, height := range cave {
			if isLowest(height, fm.adjacentLocations(c, l)) {
				lowPoints = append(lowPoints, lowPoint{
					caveIndex:     c,
					locationIndex: l,
					value:         height,
				})
			}
		}
	}
	return lowPoints
}

func (fm floorMap) adjacentLocations(caveIndex, locationIndex int) []int {
	adjacent := make([]int, 0)

	if caveAboveIndex := caveIndex - 1; caveAboveIndex >= 0 {
		adjacent = append(adjacent, fm[caveAboveIndex][locationIndex])
	}
	if caveBelowIndex := caveIndex + 1; caveBelowIndex < len(fm) {
		adjacent = append(adjacent, fm[caveBelowIndex][locationIndex])
	}

	if locationLeftIndex := locationIndex - 1; locationLeftIndex >= 0 {
		adjacent = append(adjacent, fm[caveIndex][locationLeftIndex])
	}
	if locationRightIndex := locationIndex + 1; locationRightIndex < len(fm[caveIndex]) {
		adjacent = append(adjacent, fm[caveIndex][locationRightIndex])
	}

	return adjacent
}

func (fm floorMap) basins() [][]int {
	basins := make([][]int, 0)
	for _, lowPoint := range fm.lowPoints() {
		alreadyVisited := make(map[position]bool)
		alreadyVisited[position{
			caveIndex:     lowPoint.caveIndex,
			locationIndex: lowPoint.locationIndex,
		}] = true
		basins = append(basins, fm.expandBasin(lowPoint.caveIndex, lowPoint.locationIndex, alreadyVisited))
	}
	return basins
}

type position struct {
	caveIndex     int
	locationIndex int
}

func (fm floorMap) expandBasin(caveIndex, locationIndex int, alreadyVisited map[position]bool) []int {
	if caveIndex < 0 || locationIndex < 0 || caveIndex >= len(fm) || locationIndex >= len(fm[caveIndex]) {
		return []int{}
	}
	current := fm[caveIndex][locationIndex]
	if current == 9 {
		return []int{}
	}
	basin := []int{current}

	expand := []position{
		{
			caveIndex:     caveIndex - 1,
			locationIndex: locationIndex,
		},
		{
			caveIndex:     caveIndex + 1,
			locationIndex: locationIndex,
		},
		{
			caveIndex:     caveIndex,
			locationIndex: locationIndex - 1,
		},
		{
			caveIndex:     caveIndex,
			locationIndex: locationIndex + 1,
		},
	}

	for _, p := range expand {
		if !alreadyVisited[p] {
			alreadyVisited[p] = true
			basin = append(basin, fm.expandBasin(p.caveIndex, p.locationIndex, alreadyVisited)...)
		}
	}

	return basin
}

func readFloorMap(file *os.File) floorMap {
	floorMap := make(floorMap, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawNumbers := scanner.Text()
		cave := make([]int, len(rawNumbers))
		for i, rawNumber := range rawNumbers {
			cave[i] = int(rawNumber - 48) // -48 for conversion to actual number of the numeric character
		}
		floorMap = append(floorMap, cave)
	}

	return floorMap
}

func isLowest(value int, adjacent []int) bool {
	for _, a := range adjacent {
		if a <= value {
			return false
		}
	}
	return true
}

func sumOfRiskLevels(lowPoints []lowPoint) int {
	sum := 0
	for _, point := range lowPoints {
		sum += point.value + 1
	}
	return sum
}

func productOfNLargestBasins(basins [][]int, n int) int {
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	basinSizeProduct := 1
	for i := 0; i < n; i++ {
		basinSizeProduct *= len(basins[i])
	}
	return basinSizeProduct
}
