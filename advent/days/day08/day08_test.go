package day08_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day08"
)

func TestFirstChallenge(t *testing.T) {
	day := Load(`30373
25512
65332
33549
35390`)

	if result := day.FirstChallenge(); result != "21" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	day := Load(`30373
25512
65332
33549
35390`)

	if result := day.SecondChallenge(); result != "8" {
		t.Fatal(result)
	}
}
