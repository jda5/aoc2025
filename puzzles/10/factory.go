package puzzles

import (
	"fmt"
	"strconv"
)

// -------------------------------------------------------------------------------- state bit operations

type State int8

// func (s *State) toggle(i int) {
// 	*s ^= (1 << i)
// }

func (s *State) toggle(b *State) {
	*s ^= *b
}

func (s *State) print() {
	fmt.Println(strconv.FormatInt(int64(*s), 2))
}

// -------------------------------------------------------------------------------- helpers

// func parseInput(row string) (State, []State, []int) {
// 	components := strings.Split(row, " ")
// 	target := components[0]
// 	buttons := components[1 : len(components)-1]
// 	jolage := components[len(components)-1]

// }

// -------------------------------------------------------------------------------- puzzle one

func FewestButtonPresses(input []string) int {
	res := 0
	s := State(10)
	s.print()

	b := State(100)
	b.print()

	s.toggle(&b)
	s.print()

	return res
}
