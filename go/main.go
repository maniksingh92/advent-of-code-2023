package main

import (
	"github.com/maniksingh92/advent-of-code-2023/utils"
)

var functionsByPuzzle = map[string]func(inputs []string){
	"01_1": Day01Puzzle1,
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

	f(inputs)
}
