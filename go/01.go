package main

import (
	"maps"
	"strconv"
	"strings"
)

var wordSymbols = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitSymbols = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func matchSymbol(symbols map[string]int, line string) int {
	for symbol, value := range symbols {
		if strings.HasPrefix(line, symbol) {
			return value
		}
	}
	return -1
}

func solveDay01(symbols map[string]int, inputs []string) string {
	sum := 0

	for _, line := range inputs {
		firstDigit := -1
		lastDigit := -1

		for i := 0; i < len(line); i++ {
			num := matchSymbol(symbols, line[i:])

			if num == -1 {
				continue
			}

			if firstDigit == -1 {
				firstDigit = num
			}

			lastDigit = num
		}

		sum += firstDigit*10 + lastDigit
	}

	return strconv.Itoa(sum)
}

func Day01Puzzle1(inputs []string) (string, error) {
	symbols := map[string]int{}
	maps.Copy(symbols, digitSymbols)

	return solveDay01(symbols, inputs), nil
}

func Day01Puzzle2(inputs []string) (string, error) {
	symbols := map[string]int{}
	maps.Copy(symbols, digitSymbols)
	maps.Copy(symbols, wordSymbols)

	return solveDay01(symbols, inputs), nil
}
