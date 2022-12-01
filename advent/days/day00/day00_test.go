package day00_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day00"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")

	if result := day.FirstChallenge(); result != "Merry Christmas" {
		t.Fatal(day.FirstChallenge())
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")

	if result := day.SecondChallenge(); result != "Merry Christmas" {
		t.Fatal(day.SecondChallenge())
	}
}
