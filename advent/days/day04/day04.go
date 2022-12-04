package day04

import (
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type Range struct {
	From int
	To   int
}

func (r Range) Contains(other Range) bool {
	return r.From <= other.From && other.To <= r.To
}

func (r Range) Overlaps(other Range) bool {
	return r.From <= other.To && r.To >= other.From
}

type Pair[T any, U any] struct {
	First  T
	Second U
}

type Day04 struct {
	Assignments []Pair[Range, Range]
}

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	assignments := lo.Map(lines, func(line string, _ int) Pair[Range, Range] {
		return parseLine(line)
	})
	return Day04{assignments}
}

func parseLine(line string) Pair[Range, Range] {
	ranges := strings.Split(line, ",")
	first := parseRange(ranges[0])
	second := parseRange(ranges[1])
	return Pair[Range, Range]{first, second}
}

func parseRange(str string) Range {
	boundaries := strings.Split(str, "-")
	from, _ := strconv.Atoi(boundaries[0])
	to, _ := strconv.Atoi(boundaries[1])
	return Range{from, to}
}

func (day Day04) FirstChallenge() string {
	count := lo.CountBy(day.Assignments, func(pair Pair[Range, Range]) bool {
		return pair.First.Contains(pair.Second) || pair.Second.Contains(pair.First)
	})
	return strconv.Itoa(count)
}

func (day Day04) SecondChallenge() string {
	count := lo.CountBy(day.Assignments, func(pair Pair[Range, Range]) bool {
		return pair.First.Overlaps(pair.Second)
	})
	return strconv.Itoa(count)
}
