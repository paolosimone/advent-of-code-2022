package day01

import (
	"github.com/paolosimone/advent-of-code-2022/advent/days"
)

type Day01 struct {
	Input string
}

func Load(input string) days.Day {
	return Day01{input}
}

func (day Day01) FirstChallenge() string {
	return day.Input
}

func (day Day01) SecondChallenge() string {
	return day.Input
}
