package day00_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day00"
)

func TestLoad(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day, ok := day.(Day00); !ok || day.Input != "Merry Christmas" {
		t.FailNow()
	}
}

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day.FirstChallenge() != "Merry Christmas" {
		t.FailNow()
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day.SecondChallenge() != "Merry Christmas" {
		t.FailNow()
	}
}
