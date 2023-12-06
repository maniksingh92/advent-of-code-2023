package day05

import (
	"slices"
	"strconv"
	"strings"
)

func parseSeedsLine(line string) []int {
	a := strings.Split(line, ": ")
	s := strings.Split(a[1], " ")

	seeds := []int{}
	for _, seed := range s {
		n, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, n)
	}

	return seeds
}

func parseMapNameLine(line string) [2]int {
	a := strings.Split(line, " ")
	b := strings.Split(a[0], "-to-")

	source, err := strconv.Atoi(b[0])
	if err != nil {
		panic(err)
	}

	destination, err := strconv.Atoi(b[1])
	if err != nil {
		panic(err)
	}

	return [...]int{source, destination}
}

func parseMapDataLine(line string) [3]int {
	a := strings.Split(line, " ")

	destination, err := strconv.Atoi(a[0])
	if err != nil {
		panic(err)
	}

	source, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}

	length, err := strconv.Atoi(a[2])
	if err != nil {
		panic(err)
	}

	return [...]int{destination, source, length}
}

func findSeedLocation(seed int, dataMaps [][][3]int) int {
	curr := seed

	for _, dataMap := range dataMaps {
		for _, data := range dataMap {
			destination := data[0]
			source := data[1]
			length := data[2]

			if source <= curr && curr <= source+length {
				curr = destination + (curr - source)
				break
			}
		}
	}

	return curr
}

func Day05Puzzle1(lines []string) int {
	seeds := parseSeedsLine(lines[0])

	dataMaps := [][][3]int{}

	i := 2
	for i < len(lines) {
		dataMaps = append(dataMaps, [][3]int{})

		i += 1

		for i < len(lines) && len(lines[i]) != 0 {
			dataMaps[len(dataMaps)-1] = append(dataMaps[len(dataMaps)-1], parseMapDataLine(lines[i]))
			i += 1
		}

		i += 1
	}

	locations := []int{}
	for _, seed := range seeds {
		locations = append(locations, findSeedLocation(seed, dataMaps))
	}

	return slices.Min(locations)
}
