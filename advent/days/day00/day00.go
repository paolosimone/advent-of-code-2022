package day00

import (
	"github.com/paolosimone/advent-of-code-2022/advent/days"
)

type Day00 struct {
	input string
}

func Load(input string) days.Day {
	return Day00{input}
}

func (day Day00) FirstChallenge() string {
	return day.input
}

func (day Day00) SecondChallenge() string {
	return day.input
}
