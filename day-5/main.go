package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-5/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	lines, err := readLines(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	horizontalVerticalLines := filterLines(lines, horizontalLineFilter, verticalLineFilter)
	fmt.Printf("Points covering more than 1 line: %d\n", countIntersections(generateIntersections(horizontalVerticalLines)))

	fmt.Println("Part 2")
	fmt.Printf("Points covering more than 1 line: %d\n", countIntersections(generateIntersections(lines)))

}

func readLines(file *os.File) ([]line, error) {
	lines := make([]line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linePoints := strings.Split(scanner.Text(), " -> ")
		start, err := parsePoint(linePoints[0])
		if err != nil {
			return nil, err
		}
		end, err := parsePoint(linePoints[1])
		if err != nil {
			return nil, err
		}
		lines = append(lines, line{
			start: *start,
			end:   *end,
		})
	}
	return lines, nil
}

func parsePoint(input string) (*point, error) {
	xy := strings.Split(input, ",")
	x, xErr := strconv.ParseInt(xy[0], 10, 64)
	if xErr != nil {
		return nil, fmt.Errorf("error parsing x value of point: %w", xErr)
	}
	y, yErr := strconv.ParseInt(xy[1], 10, 64)
	if yErr != nil {
		return nil, fmt.Errorf("error parsing y value of point: %w", yErr)
	}
	return &point{
		x: x,
		y: y,
	}, nil
}

type lineFilter func(line line) bool

func horizontalLineFilter(line line) bool {
	return line.start.x == line.end.x
}

func verticalLineFilter(line line) bool {
	return line.start.y == line.end.y
}

func filterLines(lines []line, lineFilter ...lineFilter) []line {
	survivingLines := make([]line, 0)
	for _, line := range lines {
		for _, filter := range lineFilter {
			if filter(line) {
				survivingLines = append(survivingLines, line)
				break
			}
		}
	}
	return survivingLines
}

func generateIntersections(lines []line) map[point]int {
	intersections := make(map[point]int)
	for _, line := range lines {
		for _, point := range line.points() {
			intersections[point] += 1
		}
	}
	return intersections
}

func countIntersections(intersections map[point]int) int {
	moreThanOne := 0
	for _, v := range intersections {
		if v > 1 {
			moreThanOne++
		}
	}
	return moreThanOne
}
