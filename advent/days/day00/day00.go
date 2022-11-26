package day00

import (
	"github.com/paolosimone/advent-of-code-2022/advent/days"
)

type Day00 struct {
	Input string
}

func Load(input string) days.Day {
	return Day00{input}
}

func (day Day00) FirstChallenge() string {
	return day.Input
}

func (day Day00) SecondChallenge() string {
	return day.Input
}
