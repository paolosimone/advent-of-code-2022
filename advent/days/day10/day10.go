package day10

import (
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

type Register = int

type Op interface {
	Apply(register Register) Register
}

type Noop struct{}

func (noop *Noop) Apply(register Register) Register {
	return register
}

type AddX struct {
	Value int
}

func (addx *AddX) Apply(register Register) Register {
	return register + addx.Value
}

type Day10 struct {
	Program []Op
}

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	program := lo.FlatMap(lines, func(line string, _ int) []Op {
		return parseOp(line)
	})
	return Day10{program}
}

func parseOp(line string) []Op {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "noop":
		return []Op{&Noop{}}
	case "addx":
		value, _ := strconv.Atoi(tokens[1])
		return []Op{&Noop{}, &AddX{Value: value}}
	default:
		panic("unknown operation")
	}
}

func (day Day10) FirstChallenge() string {
	values := make([]Register, len(day.Program) + 1)

	values[0] = 1
	for tick := 0; tick < len(day.Program); tick++ {
		values[tick+1] = day.Program[tick].Apply(values[tick])
	}

	result := 0
	for cycle := 20; cycle <= 220; cycle += 40 {
		result += cycle * values[cycle-1]
	}

	return strconv.Itoa(result)
}

func (day Day10) SecondChallenge() string {
	rows, cols := 6, 40
	crt := MakeCRT(rows, cols)

	for tick, x := 0, 1; tick < len(day.Program); tick++ {
		row := tick / cols
		col := tick % cols
		if Abs(x-col) <= 1 {
			crt[row][col] = true
		}
		x = day.Program[tick].Apply(x)
	}

	return crt.Sprint()
}

type Pixel = bool

type (
	Row []Pixel
	CRT []Row
)

func (crt CRT) Sprint() string {
	lines := lo.Map(crt, func(line Row, _ int) string {
		pixels := lo.Map(line, func(lit Pixel, _ int) string {
			return lo.Ternary(lit, "#", ".")
		})
		return strings.Join(pixels, "")
	})
	return strings.Join(lines, "\n")
}

func MakeCRT(rows int, cols int) CRT {
	crt := make(CRT, rows)
	for i := 0; i < len(crt); i++ {
		crt[i] = make(Row, cols)
	}
	return crt
}

func Abs[T constraints.Signed](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
