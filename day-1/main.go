package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := openFile("day-1/input.txt")
	if err != nil {
		panic(err)
	}

	numbers, err := readNumbers(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	fmt.Printf("Number of increases: %d\n", increasesWithSlidingWindow(numbers, 1))

	fmt.Println("Part 2")
	fmt.Printf("Number of increases: %d\n", increasesWithSlidingWindow(numbers, 3))
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readNumbers(inputFile *os.File) ([]int, error) {
	scanner := bufio.NewScanner(inputFile)
	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func increasesWithSlidingWindow(numbers []int, windowSize int) int {
	increases := 0
	for idx := 0; idx+windowSize < len(numbers); idx++ {
		aWindow := 0
		bWindow := 0
		for i := 0; i < windowSize; i++ {
			aWindow += numbers[idx+i]
			bWindow += numbers[idx+i+1]
		}

		if bWindow > aWindow {
			increases++
		}
	}
	return increases
}
