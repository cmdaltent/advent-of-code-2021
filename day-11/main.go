package main

import (
	"bufio"
	"fmt"
	"os"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-11/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	fmt.Println("Part 1")
	fmt.Printf("Number of flashes after 100 steps: %d\n", play(readInitialOctopusConfiguration(file), 100))

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2")
	fmt.Printf("First synchronized flash after round: %d\n", playUntilFirstFlashSynchronization(readInitialOctopusConfiguration(file)))
}

func readInitialOctopusConfiguration(file *os.File) [][]int {
	grid := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		energyLevels := make([]int, len(line))
		for i, v := range line {
			energyLevels[i] = int(v - 48) // convert read char to numeric representation
		}
		grid = append(grid, energyLevels)
	}
	return grid
}

func play(configuration [][]int, numberOfSteps int) int {
	numberOfFlashes := 0
	for i := 0; i < numberOfSteps; i++ {
		numberOfFlashes += playStep(configuration)
	}
	return numberOfFlashes
}

func playUntilFirstFlashSynchronization(configuration [][]int) int {
	steps := 0
	for {
		steps++
		_ = playStep(configuration)
		if isSynchronizedFlash(configuration) {
			return steps
		}
	}
}

func playStep(configuration [][]int) int {
	// increment all by one
	for _, row := range configuration {
		for col := range row {
			row[col] += 1
		}
	}

	// perform flashes
	totalNumberOfFlashes := 0
	numberOfFlashes := flashOctopuses(configuration)
	for numberOfFlashes > 0 {
		totalNumberOfFlashes += numberOfFlashes
		numberOfFlashes = flashOctopuses(configuration)
	}
	return totalNumberOfFlashes
}

func flashOctopuses(configuration [][]int) int {
	numberOfFlashes := 0
	for rowIndex, row := range configuration {
		for colIndex := range row {
			if row[colIndex] > 9 {
				row[colIndex] = 0
				increaseUnflashedAdjacent(rowIndex, colIndex, configuration)
				numberOfFlashes++
			}
		}
	}
	return numberOfFlashes
}

func increaseUnflashedAdjacent(rowIndex, colIndex int, configuration [][]int) {
	// above 1 2 3
	// row   4 X 5
	// below 6 7 8

	// above
	if rowIndex-1 >= 0 {
		// above 2
		if configuration[rowIndex-1][colIndex] != 0 {
			configuration[rowIndex-1][colIndex] += 1
		}
		// above 1
		if colIndex-1 >= 0 && configuration[rowIndex-1][colIndex-1] != 0 {
			configuration[rowIndex-1][colIndex-1] += 1
		}
		//above 3
		if colIndex+1 < len(configuration[rowIndex-1]) && configuration[rowIndex-1][colIndex+1] != 0 {
			configuration[rowIndex-1][colIndex+1] += 1
		}
	}

	// below
	if rowIndex+1 < len(configuration) {
		// below 7
		if configuration[rowIndex+1][colIndex] != 0 {
			configuration[rowIndex+1][colIndex] += 1
		}
		// below 6
		if colIndex-1 >= 0 && configuration[rowIndex+1][colIndex-1] != 0 {
			configuration[rowIndex+1][colIndex-1] += 1
		}
		//above 8
		if colIndex+1 < len(configuration[rowIndex+1]) && configuration[rowIndex+1][colIndex+1] != 0 {
			configuration[rowIndex+1][colIndex+1] += 1
		}
	}

	// row 4
	if colIndex-1 >= 0 {
		if configuration[rowIndex][colIndex-1] != 0 {
			configuration[rowIndex][colIndex-1] += 1
		}
	}

	// row 5
	if colIndex+1 < len(configuration[rowIndex]) {
		if configuration[rowIndex][colIndex+1] != 0 {
			configuration[rowIndex][colIndex+1] += 1
		}
	}
}

func isSynchronizedFlash(configuration [][]int) bool {
	for _, row := range configuration {
		for _, col := range row {
			if col != 0 {
				return false
			}
		}
	}
	return true
}
