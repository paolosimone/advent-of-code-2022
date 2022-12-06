package day06

import (
	"strconv"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
)

type Counter[T comparable] map[T]int

func (counter Counter[T]) KeysCount() int {
	return len(counter)
}

func (counter Counter[T]) Add(item T) {
	counter[item]++
}

func (counter Counter[T]) Remove(item T) {
	counter[item]--
	if counter[item] <= 0 {
		delete(counter, item)
	}
}

type Day06 struct {
	Buffer []rune
}

func Load(input string) days.Day {
	return Day06{[]rune(input)}
}

func (day Day06) FirstChallenge() string {
	return strconv.Itoa(findMarkerIndex(day.Buffer, 4))
}

func (day Day06) SecondChallenge() string {
	return strconv.Itoa(findMarkerIndex(day.Buffer, 14))
}

func findMarkerIndex(buffer []rune, target int) int {
	unique := make(Counter[rune])
	for i := 0; i < target; i++ {
		unique.Add(buffer[i])
	}
	if unique.KeysCount() == target {
		return target
	}
	for i := target; i < len(buffer); i++ {
		unique.Remove(buffer[i-target])
		unique.Add(buffer[i])
		if unique.KeysCount() == target {
			return i + 1
		}
	}
	return -1
}
