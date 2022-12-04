package day04_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day04"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	if result := day.FirstChallenge(); result != "2" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	if result := day.SecondChallenge(); result != "4" {
		t.Fatal(result)
	}
}
