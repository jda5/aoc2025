package main

import (
	"fmt"

	puzzles "github.com/jda5/aoc2025/puzzles/04"
	"github.com/jda5/aoc2025/utils"
)

func main() {

	input := utils.ReadInputLines("04")

	fmt.Println(puzzles.RemoveRolls(input))
}
