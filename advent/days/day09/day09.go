package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

type Position struct {
	X int
	Y int
}

func (pos Position) Hash() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}

type Move struct {
	X int
	Y int
}

func (move Move) Apply(pos Position) Position {
	return Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

var (
	LEFT  = Move{X: -1, Y: 0}
	RIGHT = Move{X: 1, Y: 0}
	UP    = Move{X: 0, Y: -1}
	DOWN  = Move{X: 0, Y: 1}
)

type StraightMove struct {
	Step  Move
	Count int
}

type Day09 struct {
	Moves []StraightMove
}

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	moves := lo.Map(lines, func(line string, _ int) StraightMove {
		return parseMove(line)
	})
	return Day09{moves}
}

func parseMove(line string) StraightMove {
	tokens := strings.Split(line, " ")
	direction := parseDirection(tokens[0])
	steps, _ := strconv.Atoi(tokens[1])
	return StraightMove{direction, steps}
}

func parseDirection(direction string) Move {
	switch direction {
	case "L":
		return LEFT
	case "R":
		return RIGHT
	case "D":
		return DOWN
	case "U":
		return UP
	default:
		panic("invalid direction")
	}
}

func (day Day09) FirstChallenge() string {
	result := simulate(day.Moves, 2)
	return strconv.Itoa(result)
}

func (day Day09) SecondChallenge() string {
	result := simulate(day.Moves, 10)
	return strconv.Itoa(result)
}

func simulate(moves []StraightMove, length int) int {
	knots := make([]Position, length)

	visited := make(map[string]bool)
	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			// head
			knots[0] = move.Step.Apply(knots[0])

			// tail(s)
			for knot := 1; knot < length; knot++ {
				follow := keepUp(knots[knot], knots[knot-1])
				knots[knot] = follow.Apply(knots[knot])
			}

			visited[knots[length-1].Hash()] = true
		}
	}

	return len(visited)
}

func keepUp(from Position, to Position) Move {
	delta := distance(from, to)
	if isClose(delta) {
		return Move{0, 0}
	}
	return Move{Sign(delta.X), Sign(delta.Y)}
}

func distance(from Position, to Position) Move {
	return Move{
		X: to.X - from.X,
		Y: to.Y - from.Y,
	}
}

func isClose(move Move) bool {
	return Abs(move.X) <= 1 && Abs(move.Y) <= 1
}

func Abs[T constraints.Signed](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

func Sign[T constraints.Signed](value T) T {
	switch {
	case value > 0:
		return 1
	case value < 0:
		return -1
	default:
		return 0
	}
}
