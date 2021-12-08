package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-8/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	inputs := parseInputs(file)

	fmt.Println("Part 1")
	fmt.Printf("Occurrences of 1, 4, 7, 8: %d\n", count1478(inputs))
}

func parseInputs(file *os.File) []input {
	inputs := make([]input, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		byDelimiter := strings.Split(line, " | ")

		signalPatters := strings.Split(byDelimiter[0], " ")
		digitValues := strings.Split(byDelimiter[1], " ")

		inputs = append(inputs, input{
			signalPattern: signalPatters,
			outputValue:   digitValues,
		})
	}

	return inputs
}

func count1478(in []input) (count int) {
	for _, i := range in {
		for _, dv := range i.outputValue {
			l := len(dv)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	return
}

type input struct {
	signalPattern []string
	outputValue   []string
}
