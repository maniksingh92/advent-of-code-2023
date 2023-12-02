package main

import (
	"fmt"
	"strconv"
	"unicode"
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

func Day01Puzzle1(inputs []string) error {
	sum := 0

	for _, line := range inputs {
		fmt.Println(line)
		firstDigit := -1
		lastDigit := -1

		for i := 0; i < len(line); i++ {
			char := line[i]
			if unicode.IsDigit(rune(char)) {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					return err
				}

				if firstDigit == -1 {
					firstDigit = num
				}

				lastDigit = num
			}

		}

		fmt.Printf("%d %d\n", firstDigit, lastDigit)

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println(sum)

	return nil
}
