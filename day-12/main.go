package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-12/input.txt")
	if err != nil {
		panic(err)
	}

	caveSystem := readCaveSystem(file)

	fmt.Println("Part 1")
	fmt.Printf("Number of paths through cave system: %d\n", numberOfPaths(caveSystem, 1))

	fmt.Println("Part 2")
	fmt.Printf("Number of paths through cave system: %d\n", numberOfPaths(caveSystem, 2))
}

func readCaveSystem(file *os.File) map[string][]string {
	caveSystem := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		startEnd := strings.Split(scanner.Text(), "-")
		a := startEnd[0]
		b := startEnd[1]
		if a == "start" || b == "end" {
			caveSystem[a] = append(caveSystem[a], b)
			continue
		}
		if b == "start" || a == "end" {
			caveSystem[b] = append(caveSystem[b], a)
			continue
		}
		caveSystem[a] = append(caveSystem[a], b)
		caveSystem[b] = append(caveSystem[b], a)
	}

	return caveSystem
}

func numberOfPaths(caveSystem map[string][]string, smallCaveMaxVisit int) int {
	paths := 0

	starts := caveSystem["start"]
	for _, start := range starts {
		if isSmallCave(start) {
			paths += findPathsToEnd(start, caveSystem, []string{start}, smallCaveMaxVisit)
		} else {
			paths += findPathsToEnd(start, caveSystem, []string{}, smallCaveMaxVisit)
		}

	}
	return paths
}

func findPathsToEnd(cave string, caveSystem map[string][]string, visitedSmallCaves []string, smallCaveMaxVisit int) int {
	paths := 0

	out := caveSystem[cave]
	for _, c := range out {
		if c == "end" {
			paths += 1
			continue
		}

		if wasCaveAlreadyVisited(c, visitedSmallCaves) {
			if wasAnyCaveAlreadyVisitedMoreThan(smallCaveMaxVisit, visitedSmallCaves) {
				continue
			}
		}

		if isSmallCave(c) {
			paths += findPathsToEnd(c, caveSystem, append(visitedSmallCaves, c), smallCaveMaxVisit)
		} else {
			paths += findPathsToEnd(c, caveSystem, visitedSmallCaves, smallCaveMaxVisit)
		}

	}

	return paths
}

var bigCaveRegExp = regexp.MustCompile(`[A-Z]+`)

func isSmallCave(cave string) bool {
	return !bigCaveRegExp.MatchString(cave)
}

func wasCaveAlreadyVisited(cave string, alreadyVisited []string) bool {
	for _, c := range alreadyVisited {
		if c == cave {
			return true
		}
	}
	return false
}

func wasAnyCaveAlreadyVisitedMoreThan(max int, alreadyVisited []string) bool {
	caveCount := make(map[string]int)
	for _, cave := range alreadyVisited {
		caveCount[cave] = caveCount[cave] + 1
	}
	for _, count := range caveCount {
		if count >= max {
			return true
		}
	}
	return false
}
