package day05_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day05"
)

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	if result := day.FirstChallenge(); result != "CMZ" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load(`    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	if result := day.SecondChallenge(); result != "MCD" {
		t.Fatal(result)
	}
}
