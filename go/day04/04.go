package day04

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func buildUnknownLineFormatError(line string) error {
	return fmt.Errorf("Unknown line format: %s", line)
}

func parseLine(line string) (int, []string, []string, error) {
	a := strings.Split(line, ":")
	if len(a) != 2 {
		return 0, nil, nil, buildUnknownLineFormatError(line)
	}
	cardMeta := a[0]
	numbers := a[1]

	b := strings.Fields(cardMeta)
	if len(b) != 2 {
		return 0, nil, nil, buildUnknownLineFormatError(line)
	}
	cardNumber, err := strconv.Atoi(b[1])
	if err != nil {
		return 0, nil, nil, err
	}

	c := strings.Split(numbers, "|")
	if len(c) != 2 {
		return 0, nil, nil, buildUnknownLineFormatError(line)
	}
	winningNumbers := strings.Fields(strings.TrimSpace(c[0]))
	ticketNumbers := strings.Fields(strings.TrimSpace(c[1]))

	return cardNumber, winningNumbers, ticketNumbers, nil
}

func Day04Puzzle1(lines []string) (string, error) {
	sum := 0

	for _, line := range lines {
		_, winningNumbers, ticketNumbers, err := parseLine(line)
		if err != nil {
			return "", err
		}

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

	return strconv.Itoa(sum), nil
}

func Day04Puzzle2(lines []string) (string, error) {
	sum := 0

	cardCount := map[int]int{}

	for _, line := range lines {
		cardNumber, winningNumbers, ticketNumbers, err := parseLine(line)
		if err != nil {
			return "", err
		}
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

	return strconv.Itoa(sum), nil
}
