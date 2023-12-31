package main

import (
	"fmt"

	"github.com/maniksingh92/advent-of-code-2023/day01"
	"github.com/maniksingh92/advent-of-code-2023/day02"
	"github.com/maniksingh92/advent-of-code-2023/day03"
	"github.com/maniksingh92/advent-of-code-2023/day04"
	"github.com/maniksingh92/advent-of-code-2023/day05"
	"github.com/maniksingh92/advent-of-code-2023/utils"
)

var functionsByPuzzle = map[string]func(inputs []string) int{
	"01_1": day01.Day01Puzzle1,
	"01_2": day01.Day01Puzzle2,
	"02_1": day02.Day02Puzzle1,
	"02_2": day02.Day02Puzzle2,
	"03_1": day03.Day03Puzzle1,
	"03_2": day03.Day03Puzzle2,
	"04_1": day04.Day04Puzzle1,
	"04_2": day04.Day04Puzzle2,
	"05_1": day05.Day05Puzzle1,
}

func main() {
	puzzle, inputs, err := utils.GetInputsForPuzzle()
	if err != nil {
		panic(err)
	}

	f, ok := functionsByPuzzle[puzzle]
	if !ok {
		panic("Unknown puzzle.")
	}

	fmt.Println(f(inputs))
}
