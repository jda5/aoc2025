package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/02"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInput("02")

	fmt.Println(puzzles.SumInvalidID(input))
}
