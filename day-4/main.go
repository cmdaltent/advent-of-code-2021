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
	file, err := lib.OpenFile("day-4/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	boards, numbers, err := readBoardsAndInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	fmt.Printf("Score of winning board %d\n", firstWinningBoardScore(boards, numbers))

	resetBoards(boards)

	fmt.Println("Part 2")
	fmt.Printf("Score of last winning board: %d\n", lastWinningBoardScore(boards, numbers))
}

func readBoardsAndInput(file *os.File) ([]board, []int64, error) {
	scanner := bufio.NewScanner(file)

	input := make([]int64, 0)
	_ = scanner.Scan()
	for _, number := range strings.Split(scanner.Text(), ",") {
		parsed, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number as input: %s", number)
		}
		input = append(input, parsed)
	}

	// skip over first empty line between input and boards
	_ = scanner.Scan()
	_ = scanner.Text()

	var boards []board

	var currentBoard board
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			boards = append(boards, currentBoard)
			currentBoard = board{}
			continue
		}
		var row []*number
		for _, n := range strings.Split(line, " ") {
			if n == "" {
				continue
			}
			parsed, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number on board: %s", n)
			}
			row = append(row, &number{
				value: parsed,
			})
		}
		currentBoard.rows = append(currentBoard.rows, row)
	}
	boards = append(boards, currentBoard)

	return boards, input, nil
}

func firstWinningBoardScore(boards []board, numbers []int64) int64 {
	var winningBoard *board
	var lastDrawnNumber int64
stopDrawing:
	for _, drawnNumber := range numbers {
		lastDrawnNumber = drawnNumber
		for _, board := range boards {
			board.markNumber(drawnNumber)
		}
		if wb := filterWinningBoard(boards); wb != nil {
			winningBoard = wb
			break stopDrawing
		}
	}
	return lastDrawnNumber * winningBoard.sumOfUnmarked()
}

func filterWinningBoard(boards []board) *board {
	for _, board := range boards {
		if board.hasWon() {
			return &board
		}
	}
	return nil
}

func lastWinningBoardScore(boards []board, numbers []int64) int64 {
	var lastNumber int64 = 0
search:
	for _, number := range numbers {
		for _, board := range boards {
			board.markNumber(number)
		}
		if b := keepLosingBoards(boards); len(b) > 0 {
			boards = b
			continue
		}
		lastNumber = number
		break search
	}
	return boards[0].sumOfUnmarked() * lastNumber
}

func keepLosingBoards(boards []board) []board {
	losingBoards := make([]board, 0)
	for _, board := range boards {
		if !board.hasWon() {
			losingBoards = append(losingBoards, board)
		}
	}
	return losingBoards
}

func resetBoards(boards []board) {
	for _, board := range boards {
		board.reset()
	}
}
