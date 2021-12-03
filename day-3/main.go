package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-3/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	report := readDiagnosticReport(file)

	powerConsumption, err := calculatePowerConsumption(report)
	if err != nil {
		panic(err)
	}

	lifeSupportRate, err := calculateLifeSupportRate(report)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	fmt.Printf("Power Consumption: %d\n", powerConsumption)
	fmt.Println("Part 2")
	fmt.Printf("Life Support Rate: %d\n", lifeSupportRate)
}

func readDiagnosticReport(file *os.File) []string {
	var report []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}
	return report
}

func calculatePowerConsumption(report []string) (int, error) {
	if len(report) == 0 {
		return 0, fmt.Errorf("empty report")
	}

	wordSize := len(report[0])

	gamma := 0
	epsilon := 0

	for idx := 0; idx < wordSize; idx++ {
		zeroBits := 0
		for _, s := range report {
			if s[idx] == '0' {
				zeroBits++
			}
		}
		if zeroBits > (len(report) / 2) {
			epsilon = epsilon | (1 << (wordSize - 1 - idx))
		} else {
			gamma = gamma | (1 << (wordSize - 1 - idx))
		}
	}

	return gamma * epsilon, nil
}

func calculateLifeSupportRate(report []string) (int64, error) {
	oxygenGeneratorRating, err := calculateRating(report, 0, oxygenGeneratorFilter)
	if err != nil {
		return 0, err
	}
	co2ScrubberRating, err := calculateRating(report, 0, co2ScrubberFilter)
	if err != nil {
		return 0, err
	}
	return oxygenGeneratorRating * co2ScrubberRating, nil
}

func oxygenGeneratorFilter(report []string, oneBits int) uint8 {
	var mostCommon uint8 = '0'
	if oneBits >= (len(report)+1)/2 {
		mostCommon = '1'
	}
	return mostCommon
}

func co2ScrubberFilter(report []string, oneBits int) uint8 {
	var leastCommon uint8 = '0'
	if oneBits < (len(report)+1)/2 {
		leastCommon = '1'
	}
	return leastCommon
}

type reportFilter func(report []string, oneBits int) uint8

func calculateRating(report []string, pos int, filter reportFilter) (int64, error) {
	if len(report) == 1 {
		return toNumber(report[0])
	}
	if pos > len(report[0]) {
		return 0, fmt.Errorf("bit position of %d exceeds word size of %d", pos, len(report[0]))
	}

	oneBits := 0
	for _, s := range report {
		if s[pos] == '1' {
			oneBits++
		}
	}

	return calculateRating(filterByValueOnPos(report, filter(report, oneBits), pos), pos+1, filter)
}

func toNumber(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
}

func filterByValueOnPos(report []string, value uint8, pos int) []string {
	var filtered []string

	for _, s := range report {
		if s[pos] == value {
			filtered = append(filtered, s)
		}
	}

	return filtered
}
