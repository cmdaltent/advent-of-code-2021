package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"cmdaltent/advent-of-code-2021/lib"
)

func main() {
	file, err := lib.OpenFile("day-10/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	lines := readLines(file)

	fmt.Println("Part 1")
	fmt.Printf("Score of illegal closing characters: %d\n", corruptLineScore(findCorruptedLines(lines)))

	fmt.Println("Part 2")
	fmt.Printf("Middle score of incomplete lines: %d\n", incompleteLinesScore(missingClosingSequences(lines)))
}

func readLines(file *os.File) []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func corruptLineScore(invalidClosings []rune) int {
	score := 0
	for _, c := range invalidClosings {
		score += scorePointsIllegalClosing[c]
	}
	return score
}

func findCorruptedLines(lines []string) []rune {
	illegalClosings := make([]rune, 0)

	for _, line := range lines {
		expectedClosingCharacters := ""
		for _, c := range line {
			if isOpeningCharacter(c) {
				expectedClosingCharacters += string(closingCharacters[c])
				continue
			}
			if len(expectedClosingCharacters) == 0 {
				illegalClosings = append(illegalClosings, c)
				break
			}
			if rune(expectedClosingCharacters[len(expectedClosingCharacters)-1]) == c {
				expectedClosingCharacters = expectedClosingCharacters[:len(expectedClosingCharacters)-1]
				continue
			}
			illegalClosings = append(illegalClosings, c)
			break
		}
	}

	return illegalClosings
}

func incompleteLinesScore(missingClosings []string) int {
	closingsScore := make([]int, len(missingClosings))
	for i, closing := range missingClosings {
		sum := 0
		for c := len(closing) - 1; c >= 0; c-- {
			sum = (sum * 5) + scorePointsIncompleteLine[rune(closing[c])]
		}
		closingsScore[i] = sum
	}
	sort.Slice(closingsScore, func(i, j int) bool {
		return closingsScore[i] < closingsScore[j]
	})
	return closingsScore[(len(closingsScore) / 2)] // as there is always an odd number of scores
}

func missingClosingSequences(lines []string) []string {
	closings := make([]string, 0)

	for _, line := range lines {
		missingClosingCharacters := ""
		lineWasLegal := true
		for _, c := range line {
			if isOpeningCharacter(c) {
				missingClosingCharacters += string(closingCharacters[c])
				continue
			}
			if len(missingClosingCharacters) == 0 {
				lineWasLegal = false
				break
			}
			if rune(missingClosingCharacters[len(missingClosingCharacters)-1]) == c {
				missingClosingCharacters = missingClosingCharacters[:len(missingClosingCharacters)-1]
				continue
			}
			lineWasLegal = false
			break
		}
		if lineWasLegal {
			closings = append(closings, missingClosingCharacters)
		}
	}

	return closings
}

func isOpeningCharacter(c rune) bool {
	return c == '(' || c == '[' || c == '{' || c == '<'
}

var closingCharacters = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scorePointsIllegalClosing = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var scorePointsIncompleteLine = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
