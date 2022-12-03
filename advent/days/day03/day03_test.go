package day03_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day03"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	if result := day.FirstChallenge(); result != "157" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	if result := day.SecondChallenge(); result != "70" {
		t.Fatal(result)
	}
}
