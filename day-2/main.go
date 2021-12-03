package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

type direction uint8

const (
	forward direction = iota
	up
	down
)

type command struct {
	direction direction
	value     int64
}

func main() {
	file, err := lib.OpenFile("day-2/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	commands, err := readCommands(file)
	if err != nil {
		panic(err)
	}

	var horizontalPos int64
	var depth int64

	fmt.Println("Part 1")
	horizontalPos, depth = followCourseWithoutAim(commands)
	fmt.Printf("horizontal position * depth: %d\n", horizontalPos*depth)

	fmt.Println("Part 1")
	horizontalPos, depth = followCourse(commands)
	fmt.Printf("horizontal position * depth: %d\n", horizontalPos*depth)
}

func readCommands(file *os.File) ([]command, error) {
	var commands []command

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return nil, fmt.Errorf("malformed line: %s", line)
		}

		direction, err := parseDirection(split[0])
		if err != nil {
			return nil, err
		}
		value, err := parseValue(split[1])
		if err != nil {
			return nil, err
		}

		command := command{
			direction: direction,
			value:     value,
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func parseDirection(in string) (direction, error) {
	var direction direction
	switch in {
	case "forward":
		direction = forward
	case "up":
		direction = up
	case "down":
		direction = down
	default:
		return direction, fmt.Errorf("invalid direction: %s", in)
	}
	return direction, nil
}

func parseValue(in string) (int64, error) {
	value, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value for direction: %s", in)
	}
	return value, nil
}

func followCourseWithoutAim(commands []command) (horizontalPosition int64, depth int64) {
	for _, command := range commands {
		switch command.direction {
		case forward:
			horizontalPosition += command.value
		case up:
			depth -= command.value
		case down:
			depth += command.value
		}
	}
	return
}

func followCourse(commands []command) (horizontalPosition int64, depth int64) {
	var aim int64
	for _, command := range commands {
		switch command.direction {
		case forward:
			horizontalPosition += command.value
			depth += command.value * aim
		case up:
			aim -= command.value
		case down:
			aim += command.value
		}
	}
	return
}
