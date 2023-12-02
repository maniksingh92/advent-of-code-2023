package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const inputsDir string = "../inputs"

func getFilePath() (string, error) {
	// accept first argument as file path if it exists
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	// attempt to check if inputs exist in an expected path
	program := strings.Split(os.Args[0], "/")
	expectedInputFileName := program[len(program)-1] + ".txt"
	expectedInputFileFullPath := filepath.Join(inputsDir, expectedInputFileName)
	_, err := os.Stat(expectedInputFileFullPath)
	if !errors.Is(err, os.ErrNotExist) {
		return expectedInputFileFullPath, nil
	}

	// could not automatically find file, ask user to provide one through the shell input
	br := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	line, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func ProcessInputFilePathArg() ([]string, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	// open file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// deserialize to string slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
