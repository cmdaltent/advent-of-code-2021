package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-7/input.txt")
	if err != nil {
		panic(err)
	}

	positions, err := parsePositions(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	alignmentConstantFuelConsumption, leastFuelAtConstantConsumption := optimalAlignment(positions, constantFuelConsumption)
	fmt.Printf("Least fuel at position %d with fuel %0.f\n", alignmentConstantFuelConsumption, leastFuelAtConstantConsumption)

	fmt.Println("Part 2")
	alignmentLinearFuelConsumption, leasFuelAtLinearConsumption := optimalAlignment(positions, linearFuelConsumption)
	fmt.Printf("Least fuel at position %d with fuel %0.f.\n", alignmentLinearFuelConsumption, leasFuelAtLinearConsumption)
}

func parsePositions(file *os.File) ([]int64, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	rawPositions := strings.Split(string(content), ",")

	positions := make([]int64, len(rawPositions))
	for i, age := range rawPositions {
		parsed, err := strconv.ParseInt(age, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid fish age: %w", err)
		}
		positions[i] = parsed
	}
	return positions, nil
}

type fuelConsumption func(startPos, endPos int64) float64

func constantFuelConsumption(startPos, endPos int64) float64 {
	return math.Abs(float64(startPos - endPos))
}

func linearFuelConsumption(startPos, endPos int64) float64 {
	// fuel consumption grows linearly with distance
	// sum it up
	// 1 + 2 + 3 + 4 + ... + n = .5n(n+1)
	n := math.Abs(float64(startPos - endPos))
	return 0.5 * n * (n + 1)
}

func optimalAlignment(positions []int64, consumption fuelConsumption) (int64, float64) {
	leastFuel := math.MaxFloat64
	var horizontalAlignment int64 = 0

	minPosition, maxPosition := minMaxPositions(positions)

	for i := minPosition; i <= maxPosition; i++ {
		fuel := 0.0
		for _, position := range positions {
			fuel += consumption(i, position)
		}
		if fuel < leastFuel {
			leastFuel = fuel
			horizontalAlignment = i
		}
	}
	return horizontalAlignment, leastFuel
}

func minMaxPositions(positions []int64) (int64, int64) {
	var minPosition int64 = math.MaxInt64
	var maxPosition int64 = 0

	for _, position := range positions {
		if position < minPosition {
			minPosition = position
		}
		if position > maxPosition {
			maxPosition = position
		}
	}

	return minPosition, maxPosition
}
