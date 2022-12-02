package day02

import (
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissor
)

const ShapeCount = 3

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

type Strategy int

type Match struct {
	Mine     Strategy
	Opponent Shape
}

type Score = int

const (
	LoseScore Score = 0
	DrawScore Score = 3
	WinScore  Score = 6
)

type Day02 struct {
	Matches []Match
}

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	matches := lo.Map(lines, func(line string, _ int) Match { return ParseLine(line) })
	return Day02{matches}
}

func ParseLine(line string) Match {
	chars := []rune(line)
	opponent := Shape(chars[0]) - Shape('A')
	mine := Strategy(chars[2]) - Strategy('X')
	return Match{mine, opponent}
}

func (day Day02) FirstChallenge() string {
	scores := lo.Map(day.Matches, func(match Match, _ int) Score {
		return scoreFirstChallenge(match)
	})
	totalScore := lo.Sum(scores)
	return strconv.Itoa(totalScore)
}

func scoreFirstChallenge(match Match) Score {
	shapeScore := Score(match.Mine) + 1
	return shapeScore + outcomeScoreFirstChallenge(match)
}

func outcomeScoreFirstChallenge(match Match) Score {
	switch {
	case Shape(match.Mine) == match.Opponent:
		return DrawScore
	case Shape(match.Mine) == winningShape(match.Opponent):
		return WinScore
	default:
		return LoseScore
	}
}

func (day Day02) SecondChallenge() string {
	scores := lo.Map(day.Matches, func(match Match, _ int) Score {
		return scoreSecondChallenge(match)
	})
	totalScore := lo.Sum(scores)
	return strconv.Itoa(totalScore)
}

func scoreSecondChallenge(match Match) Score {
	switch {
	case Outcome(match.Mine) == Draw:
		return Score(match.Opponent) + 1 + DrawScore
	case Outcome(match.Mine) == Win:
		return Score(winningShape(match.Opponent)) + 1 + WinScore
	default:
		return Score(losingShape(match.Opponent)) + 1 + LoseScore
	}
}

func winningShape(shape Shape) Shape {
	return (shape + 1) % ShapeCount
}

func losingShape(shape Shape) Shape {
	return (ShapeCount + shape - 1) % ShapeCount
}
