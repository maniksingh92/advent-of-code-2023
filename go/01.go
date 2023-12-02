package main

import (
	"fmt"
	"maps"
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

func Day01Puzzle1(inputs []string) error {
	sum := 0

	symbols := map[string]int{}
	maps.Copy(symbols, digitSymbols)

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

	fmt.Println(sum)

	return nil
}

func Day01Puzzle2(inputs []string) error {
	sum := 0

	symbols := map[string]int{}
	maps.Copy(symbols, wordSymbols)
	maps.Copy(symbols, digitSymbols)

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

	fmt.Println(sum)

	return nil
}
