package advent

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/paolosimone/advent-of-code-2022/advent/days/day01"
	"github.com/samber/lo"
)

var Days = [...]days.DayLoader{day01.Load}

type DayResult struct {
	Day           int
	LoadElapsed   time.Duration
	FirstResult   string
	FirstElapsed  time.Duration
	SecondResult  string
	SecondElapsed time.Duration
}

func RunDay(number int) DayResult {
	input := readDayInput(number)
	loadDay := Days[number-1]

	day, loadElapsed := timed(func() days.Day { return loadDay(input) })
	firstResult, firstElapsed := timed(func() string { return day.FirstChallenge() })
	secondResult, secondElapsed := timed(func() string { return day.SecondChallenge() })

	return DayResult{
		number,
		loadElapsed,
		firstResult,
		firstElapsed,
		secondResult,
		secondElapsed,
	}
}

func RunAllDays() []DayResult {
	return lo.Times(len(Days), func(index int) DayResult { return RunDay(index + 1) })
}

// because f*ck portability, amiright?
func readDayInput(day int) string {
	dayFolder := fmt.Sprintf("day%02d", day)
	inputFile := path.Join("advent", "days", dayFolder, "input.txt")
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(bytes)
}

func timed[R any](task func() R) (R, time.Duration) {
	start := time.Now()
	result := task()
	elapsed := time.Since(start)
	return result, elapsed
}
