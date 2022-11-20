package days

type Day interface {
	FirstChallenge() string
	SecondChallenge() string
}

type DayLoader func(input string) Day
