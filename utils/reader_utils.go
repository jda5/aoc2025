package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func loadFile(puzzleNumber string) []byte {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not find project.")
	}

	root := filepath.Join(filepath.Dir(file), "..")

	inputPath := filepath.Join(root, "puzzles", puzzleNumber, "input.txt")

	data, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return data
}

func ReadInput(puzzleNumber string) string {
	return string(loadFile(puzzleNumber))
}

func ReadInputLines(puzzleNumber string) []string {
	input := ReadInput(puzzleNumber)
	return strings.Split(input, "\n")
}
