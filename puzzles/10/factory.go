package puzzles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jda5/aoc2025/utils"
)

// -------------------------------------------------------------------------------- state bit operations

type State int16

func (s *State) toggleBit(i int) {
	*s ^= (1 << i)
}

// useful for debugging
func (s *State) print(size int) {
	strconv.FormatInt(int64(*s), 2)
	binStr := strconv.FormatInt(int64(*s), 2)
	for len(binStr) < size {
		binStr = "0" + binStr
	}
	fmt.Println(binStr)
}

func toggle(s State, b *State) State {
	return s ^ *b
}

type StatePresses struct {
	state   State
	presses int
}

// -------------------------------------------------------------------------------- queue data structure
// https://medium.com/@danielabatibabatunde1/mastering-queues-in-golang-be77414abe9e

type Queue[T any] []T

// Enqueue adds an element to the rear of the queue
func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes and returns an element from the front of the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// -------------------------------------------------------------------------------- helpers

func parseInput(row string) (State, []State, []int) {
	components := strings.Split(row, " ")

	targetString := components[0]
	targetLength := len(targetString[1 : len(targetString)-1])
	var targetBinary int16
	for _, char := range targetString[1 : len(targetString)-1] {
		targetBinary <<= 1 // Shift left by 1 bit
		if char == '#' {
			targetBinary |= 1 // Set the rightmost bit to 1
		}
		// If char is '.', we leave the bit as 0 (do nothing)
	}
	target := State(targetBinary)

	buttons := make([]State, 0)
	for _, buttonString := range components[1 : len(components)-1] {
		indexes := strings.Split(buttonString[1:len(buttonString)-1], ",")
		button := State(0)
		for _, idxString := range indexes {
			i, err := strconv.Atoi(idxString)
			utils.Check(err)
			button.toggleBit(targetLength - i - 1)
		}
		buttons = append(buttons, button)
	}

	joltages := make([]int, 0)
	jolageString := components[len(components)-1]
	for numString := range strings.SplitSeq(jolageString[1:len(jolageString)-1], ",") {
		num, err := strconv.Atoi(numString)
		utils.Check(err)
		joltages = append(joltages, num)
	}

	return target, buttons, joltages
}

// -------------------------------------------------------------------------------- puzzle one

func calcualtePresses(target State, buttons []State) int {

	q := Queue[StatePresses]{}
	states := make(map[State]struct{}, 0)

	initial := StatePresses{state: State(0), presses: 0}

	q.Enqueue(initial)
	states[initial.state] = struct{}{}

	for !q.IsEmpty() {
		current, err := q.Dequeue()
		utils.Check(err)

		if current.state == target {
			return current.presses
		}

		for _, button := range buttons {
			newState := toggle(current.state, &button)
			if _, ok := states[newState]; !ok {
				states[newState] = struct{}{}
				q.Enqueue(StatePresses{state: newState, presses: current.presses + 1})
			}
		}
	}
	return -1
}

func FewestButtonPresses(input []string) int {
	res := 0
	for _, row := range input {
		target, buttons, _ := parseInput(row)
		presses := calcualtePresses(target, buttons)
		res += presses
	}
	return res
}

// -------------------------------------------------------------------------------- puzzle two
// doesn't work

type StatePressesJoltage struct {
	state   State
	presses int
	joltage []int
}

func (s *StatePressesJoltage) toString() string {
	return fmt.Sprintf("%v-%v-%v", s.state, s.presses, s.joltage)
}

func compareJoltage(currJoltage []int, targetJoltage []int) int {
	equal := true
	for i := range currJoltage {
		if currJoltage[i] != targetJoltage[i] {
			equal = false
			if currJoltage[i] > targetJoltage[i] {
				return 1
			}
		}
	}
	if equal {
		return 2
	}
	return 0
}

func incrementJoltage(joltage []int, button *State) []int {

	result := make([]int, len(joltage))
	copy(result, joltage)

	for i := range joltage {
		// Check if the bit at position i (from the right) is set
		if (*button)&(1<<i) != 0 {
			result[len(result)-i-1]++
		}
	}
	return result
}

func calcualtePressesWithJoltage(targetLights State, buttons []State, targetJoltage []int) int {

	q := Queue[StatePressesJoltage]{}
	states := make(map[string]struct{}, 0)

	initial := StatePressesJoltage{state: State(0), presses: 0, joltage: make([]int, len(targetJoltage))}

	q.Enqueue(initial)
	states[initial.toString()] = struct{}{}

	for !q.IsEmpty() {
		current, err := q.Dequeue()
		utils.Check(err)

		joltageState := compareJoltage(current.joltage, targetJoltage)
		if joltageState == 2 {
			if current.state == targetLights {
				return current.presses
			}
			// since joltages are equal, any further button presses will result in unequal joltages
			continue
		}

		if joltageState == 1 {
			// joltage state has exceeded the alloted max
			continue
		}

		for _, button := range buttons {
			buttonPtr := &button

			newState := StatePressesJoltage{
				state:   toggle(current.state, buttonPtr),
				presses: current.presses + 1,
				joltage: incrementJoltage(current.joltage, buttonPtr),
			}

			newStateString := newState.toString()

			if _, ok := states[newStateString]; !ok {
				states[newStateString] = struct{}{}
				q.Enqueue(newState)
			}
		}
	}
	return -1
}

func FewestButtonPressesWithJoltage(input []string) int {
	res := 0
	for _, row := range input {
		target, buttons, joltage := parseInput(row)
		presses := calcualtePressesWithJoltage(target, buttons, joltage)
		res += presses
	}
	return res
}
