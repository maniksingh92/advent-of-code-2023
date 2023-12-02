package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var symbols = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
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
