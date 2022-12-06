package day06_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day06"
)

func TestFirstChallenge(t *testing.T) {
	testCases := []struct {
		input  string
		result string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "7"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "5"},
		{"nppdvjthqldpwncqszvftbrmjlhg", "6"},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "10"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11"},
	}
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			day := Load(test.input)

			if result := day.FirstChallenge(); result != test.result {
				t.Fatal(result)
			}
		})
	}
}

func TestSecondChallenge(t *testing.T) {
	testCases := []struct {
		input  string
		result string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "23"},
		{"nppdvjthqldpwncqszvftbrmjlhg", "23"},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "29"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "26"},
	}
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			day := Load(test.input)

			if result := day.SecondChallenge(); result != test.result {
				t.Fatal(result)
			}
		})
	}
}
