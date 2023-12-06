package day03

import (
	"strconv"
	"unicode"
)

type SchematicQueue [][]rune

func (sch SchematicQueue) Append(s string) SchematicQueue {
	sch = append(sch, []rune(s))
	// maintain the original length of the queue during initialization
	// e.g. if the queue was initialize with a length 3, this would ensure
	// the slice remains of length 3 after append.
	return sch[1:]
}

func (sch SchematicQueue) GetSumOfValidPartNumbersFromCurrentLine() int {
	validPositions := make([]bool, len(sch[1]))
	for _, line := range sch {
		for i, char := range line {
			validPositions[i] = validPositions[i] || !unicode.IsDigit(char) && string(char) != "."
		}
	}

	var validPartNumbers []int

	currentLine := sch[1]
	i := 0
	for i < len(currentLine) {
		if !unicode.IsDigit(currentLine[i]) {
			i += 1
			continue
		}

		start := i

		isValid := false

		// concatenate a numerical string
		numStr := []rune{}
		for i < len(currentLine) && unicode.IsDigit(currentLine[i]) {
			numStr = append(numStr, currentLine[i])
			// identify if the numerical string is valid
			isValid = isValid || validPositions[i]
			i += 1
		}

		// if valid position was not found at the indices of the numerical string,
		// check if valid positions exist left or right of the numerical string.
		if start != 0 && validPositions[start-1] {
			isValid = true
		}

		if i < len(currentLine) && validPositions[i] {
			isValid = true
		}

		// does numerical string belong to valid part number?
		if isValid {
			num, err := strconv.Atoi(string(numStr))
			if err != nil {
				panic(err)
			}

			validPartNumbers = append(validPartNumbers, num)
		}
	}

	sum := 0
	for _, num := range validPartNumbers {
		sum += num
	}

	return sum
}

func (sch SchematicQueue) GetSumOfValidGearRatiosOfCurrentLine() int {
	gears := make([]bool, len(sch[1]))
	for i, char := range sch[1] {
		if string(char) == "*" {
			gears[i] = true
		}
	}

	var validGearRatios []int

	for gearIdx, value := range gears {
		if value != true {
			continue
		}

		adjacentNumbers := []int{}

		for _, currentLine := range sch {
			i := 0
			for i < len(currentLine) {
				if !unicode.IsDigit(currentLine[i]) {
					i += 1
					continue
				}

				start := i

				isAdjacent := false

				// concatenate a numerical string
				numStr := []rune{}
				for i < len(currentLine) && unicode.IsDigit(currentLine[i]) {
					numStr = append(numStr, currentLine[i])
					// identify if the numerical string is valid
					isAdjacent = isAdjacent || i == gearIdx
					i += 1
				}

				if start != 0 && start-1 == gearIdx {
					isAdjacent = true
				}

				if i < len(currentLine) && i == gearIdx {
					isAdjacent = true
				}

				// is number adjacent to a gear
				if isAdjacent {
					num, err := strconv.Atoi(string(numStr))
					if err != nil {
						panic(err)
					}
					adjacentNumbers = append(adjacentNumbers, num)
				}
			}

		}

		if len(adjacentNumbers) == 2 {
			validGearRatios = append(validGearRatios, adjacentNumbers[0]*adjacentNumbers[1])
		}
	}

	sum := 0
	for _, num := range validGearRatios {
		sum += num
	}

	return sum
}

func Day03Puzzle1(inputs []string) int {
	if len(inputs) < 1 {
		panic("Inputs must be provided")
	}

	schematicQueue := make(SchematicQueue, 3)
	sum := 0

	// init by adding the first input line
	schematicQueue = schematicQueue.Append(inputs[0])

	// iterate from the next input line onwards
	for _, input := range inputs[1:] {
		schematicQueue = schematicQueue.Append(input)
		sumOfCurrentLine := schematicQueue.GetSumOfValidPartNumbersFromCurrentLine()
		sum += sumOfCurrentLine
	}

	// process last input line
	schematicQueue = schematicQueue.Append("")
	sumOfCurrentLine := schematicQueue.GetSumOfValidPartNumbersFromCurrentLine()
	sum += sumOfCurrentLine

	return sum
}

func Day03Puzzle2(inputs []string) int {
	if len(inputs) < 1 {
		panic("Inputs must be provided")
	}

	schematicQueue := make(SchematicQueue, 3)
	sum := 0

	// init by adding the first input line
	schematicQueue = schematicQueue.Append(inputs[0])

	// iterate from the next input line onwards
	for _, input := range inputs[1:] {
		schematicQueue = schematicQueue.Append(input)
		sumOfCurrentLine := schematicQueue.GetSumOfValidGearRatiosOfCurrentLine()
		sum += sumOfCurrentLine
	}

	// process last input line
	schematicQueue = schematicQueue.Append("")
	sumOfCurrentLine := schematicQueue.GetSumOfValidGearRatiosOfCurrentLine()
	sum += sumOfCurrentLine

	return sum
}
