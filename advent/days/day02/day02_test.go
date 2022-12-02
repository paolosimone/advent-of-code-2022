package day02_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day02"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`A Y
B X
C Z`)

	if result := day.FirstChallenge(); result != "15" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`A Y
B X
C Z`)

	if result := day.SecondChallenge(); result != "12" {
		t.Fatal(result)
	}
}
