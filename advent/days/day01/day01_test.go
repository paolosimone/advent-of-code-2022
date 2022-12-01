package day01_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day01"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()

	day := Load(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

	if result := day.FirstChallenge(); result != "24000" {
		t.Fatal(day.FirstChallenge())
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
	if result := day.SecondChallenge(); result != "45000" {
		t.Fatal(day.SecondChallenge())
	}
}
