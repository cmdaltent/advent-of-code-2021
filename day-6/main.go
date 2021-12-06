package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-6/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	swarm, err := parseSwarm(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	fmt.Printf("Fish after 80 days: %d\n", simulateReproduction(swarm, 80))

	fmt.Println("Part 2")
	fmt.Printf("Fish after 256 days: %d\n", simulateReproduction(swarm, 256))
}

func parseSwarm(file *os.File) ([]int8, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	rawFishAges := strings.Split(string(content), ",")

	fishAges := make([]int8, len(rawFishAges))
	for i, age := range rawFishAges {
		parsed, err := strconv.ParseInt(age, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("invalid fish age: %w", err)
		}
		fishAges[i] = int8(parsed)
	}
	return fishAges, nil
}

func simulateReproduction(initial []int8, days int) int64 {
	ageDistribution := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 7; i++ {
		var ages int64 = 0
		for _, v := range initial {
			if v == int8(i) {
				ages++
			}
		}
		ageDistribution[i] = ages
	}

	for days > 0 {
		newFish := ageDistribution[0]

		ageDistribution[0] = ageDistribution[1]
		ageDistribution[1] = ageDistribution[2]
		ageDistribution[2] = ageDistribution[3]
		ageDistribution[3] = ageDistribution[4]
		ageDistribution[4] = ageDistribution[5]
		ageDistribution[5] = ageDistribution[6]
		ageDistribution[6] = ageDistribution[7] + newFish
		ageDistribution[7] = ageDistribution[8]
		ageDistribution[8] = newFish

		days--
	}

	var fishes int64 = 0
	for _, sameAge := range ageDistribution {
		fishes += sameAge
	}
	return fishes
}
