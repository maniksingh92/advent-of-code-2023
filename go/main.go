package main

import (
	"fmt"

	"github.com/maniksingh92/advent-of-code-2023/utils"
)

var functionsByPuzzle = map[string]func(inputs []string) (string, error){
	"01_1": Day01Puzzle1,
	"01_2": Day01Puzzle2,
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

	answer, err := f(inputs)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer)
}
