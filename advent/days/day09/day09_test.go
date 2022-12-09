package day09_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day09"
)

func TestFirstChallenge(t *testing.T) {
	day := Load(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

	if result := day.FirstChallenge(); result != "13" {
		t.Fatal(result)
	}
}

func TestSecondChallengeSimple(t *testing.T) {
	day := Load(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

	if result := day.SecondChallenge(); result != "1" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	day := Load(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)

	if result := day.SecondChallenge(); result != "36" {
		t.Fatal(result)
	}
}
