package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var wordSymbols = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var digitSymbols = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

func Day01Puzzle1(inputs []string) {
	sum := 0

	for _, input := range inputs {
		// find first digit
		var firstDigit string
		for i := 0; i < len(input); i++ {
			char := input[i]
			if unicode.IsDigit(rune(char)) {
				firstDigit = string(char)
				break
			}
		}

		// find last digit
		var lastDigit string
		for i := len(input) - 1; i >= 0; i-- {
			char := input[i]
			if unicode.IsDigit(rune(char)) {
				lastDigit = string(char)
				break
			}
		}

		num, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println(sum)
}