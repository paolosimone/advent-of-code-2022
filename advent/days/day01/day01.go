package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type (
	Calories = int
)

type Day01 struct {
	Elfs []Calories
}

func Load(input string) days.Day {
	sections := strings.Split(input, "\n\n")
	elfs := lo.Map(sections, func(section string, _ int) Calories { return ParseCalories(section) })
	return Day01{elfs}
}

func ParseCalories(section string) Calories {
	lines := strings.Split(section, "\n")
	calories := lo.Map(lines, func(line string, _ int) Calories {
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		return calories
	})
	return lo.Sum(calories)
}

func (day Day01) FirstChallenge() string {
	maximum := lo.Max(day.Elfs)
	return strconv.Itoa(maximum)
}

func (day Day01) SecondChallenge() string {
	sort.Ints(day.Elfs)
	top3Sum := lo.Sum(day.Elfs[len(day.Elfs)-3:])
	return strconv.Itoa(top3Sum)
}
