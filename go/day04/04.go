package day04

import (
	"math"
	"strconv"
	"strings"
)

func parseLine(line string) (int, []string, []string) {
	a := strings.Split(line, ":")
	cardMeta := a[0]
	numbers := a[1]

	b := strings.Fields(cardMeta)
	cardNumber, err := strconv.Atoi(b[1])
	if err != nil {
		panic(err)
	}

	c := strings.Split(numbers, "|")
	winningNumbers := strings.Fields(strings.TrimSpace(c[0]))
	ticketNumbers := strings.Fields(strings.TrimSpace(c[1]))

	return cardNumber, winningNumbers, ticketNumbers
}

func Day04Puzzle1(lines []string) int {
	sum := 0

	for _, line := range lines {
		_, winningNumbers, ticketNumbers := parseLine(line)

		winningNumbersMap := map[string]bool{}
		for _, n := range winningNumbers {
			winningNumbersMap[n] = true
		}

		found := []string{}

		for _, n := range ticketNumbers {
			if winningNumbersMap[n] {
				found = append(found, n)

				if len(found) == len(winningNumbers) {
					break
				}
			}
		}

		if len(found) != 0 {
			sum += int(math.Pow(2, float64(len(found)-1)))
		}
	}

	return sum
}

func Day04Puzzle2(lines []string) int {
	sum := 0

	cardCount := map[int]int{}

	for _, line := range lines {
		cardNumber, winningNumbers, ticketNumbers := parseLine(line)
		cardCount[cardNumber] += 1

		winningNumbersMap := map[string]bool{}
		for _, n := range winningNumbers {
			winningNumbersMap[n] = true
		}

		found := []string{}

		for _, n := range ticketNumbers {
			if winningNumbersMap[n] {
				found = append(found, n)

				if len(found) == len(winningNumbers) {
					break
				}
			}
		}

		if len(found) != 0 {
			for i := 1; i <= len(found); i++ {
				cardCount[cardNumber+i] += cardCount[cardNumber]
			}
		}
	}

	for _, count := range cardCount {
		sum += count
	}

	return sum
}
