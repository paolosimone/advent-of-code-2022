package day03

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type (
	Item = rune
	Sack = []Item
)

type Day03 struct {
	Rucksacks []Sack
}

type Priority = int

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	rucksacks := lo.Map(lines, func(line string, _ int) Sack {
		return []Item(line)
	})
	return Day03{rucksacks}
}

func (day Day03) FirstChallenge() string {
	priorities := lo.Map(day.Rucksacks, func(items Sack, _ int) Priority {
		middle := len(items) / 2
		shared := lo.Intersect(items[:middle], items[middle:])
		return priority(shared[0])
	})
	total := lo.Sum(priorities)
	return strconv.Itoa(total)
}

func (day Day03) SecondChallenge() string {
	groups := lo.Chunk(day.Rucksacks, 3)
	priorities := lo.Map(groups, func(sacks []Sack, _ int) Priority {
		shared := lo.Reduce(sacks[1:], func(shared Sack, sack Sack, _ int) Sack {
			return lo.Intersect(shared, sack)
		}, sacks[0])
		return priority(shared[0])
	})
	total := lo.Sum(priorities)
	return strconv.Itoa(total)
}

func priority(item Item) Priority {
	if unicode.IsLower(item) {
		return int(item) - int('a') + 1
	}
	return int(item) - int('A') + 27
}
