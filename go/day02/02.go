package day02

import (
	"strconv"
	"strings"
)

type gameSet map[string]int

func parseGameSetStr(gameSetStr string) gameSet {
	cubes := strings.Split(gameSetStr, ", ")

	data := gameSet{}

	for _, cubeStr := range cubes {
		cube := strings.Split(cubeStr, " ")
		color := cube[1]
		count, err := strconv.Atoi(cube[0])
		if err != nil {
			panic(err)
		}

		data[color] = count
	}

	return data
}

func parseLine(line string) (int, []gameSet) {
	record := strings.Split(line, ": ")
	game, err := strconv.Atoi(strings.Split(record[0], " ")[1:][0])
	if err != nil {
		panic(err)
	}

	gameSetStrs := strings.Split(record[1], "; ")
	gameSets := []gameSet{}

	for _, gameSetStr := range gameSetStrs {
		gameSets = append(gameSets, parseGameSetStr(gameSetStr))
	}

	return game, gameSets
}

var maximumAllowed = gameSet{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isGameSetValid(gameSet gameSet) bool {
	for color, value := range gameSet {
		if value > maximumAllowed[color] {
			return false
		}
	}

	return true
}

func isGameValid(gameSets []gameSet) bool {
	for _, gameSet := range gameSets {
		if !isGameSetValid(gameSet) {
			return false
		}
	}
	return true
}

func getPowerForGame(gameSets []gameSet) int {
	minimum := gameSet{}

	for _, gameSet := range gameSets {
		for color, value := range gameSet {
			if value > minimum[color] {
				minimum[color] = value
			}
		}
	}

	power := 1
	for _, value := range minimum {
		power *= value
	}

	return power
}

func Day02Puzzle1(inputs []string) int {
	sum := 0

	for _, line := range inputs {
		game, gameSets := parseLine(line)

		if isGameValid(gameSets) {
			sum += game
		}
	}

	return sum
}

func Day02Puzzle2(inputs []string) int {
	sum := 0

	for _, line := range inputs {
		_, gameSets := parseLine(line)
		sum += getPowerForGame(gameSets)
	}

	return sum
}
