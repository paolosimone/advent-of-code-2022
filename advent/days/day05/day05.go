package day05

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type (
	Item      = string
	ItemStack []Item
)

func (stack *ItemStack) Push(item Item) {
	*stack = append(*stack, item)
}

func (stack ItemStack) Peek() Item {
	return stack[len(stack)-1]
}

func (stack *ItemStack) Pop() Item {
	top := stack.Peek()
	*stack = (*stack)[:len(*stack)-1]
	return top
}

type Move struct {
	Count int
	From  int
	To    int
}

type Day05 struct {
	Stacks []ItemStack
	Moves  []Move
}

func (day Day05) Clone() Day05 {
	return Day05{
		Stacks: lo.Map(day.Stacks, func(stack ItemStack, _ int) ItemStack {
			return append(ItemStack(nil), stack...)
		}),
		Moves: append([]Move(nil), day.Moves...),
	}
}

func Load(input string) days.Day {
	sections := strings.Split(input, "\n\n")
	stacks := parseStacks(sections[0])
	moves := parseMoves(sections[1])
	return Day05{stacks, moves}
}

func parseStacks(input string) []ItemStack {
	lines := lo.Reverse(strings.Split(input, "\n"))

	indexes := strings.Split(lines[0], "")
	count, _ := strconv.Atoi(indexes[len(indexes)-2])

	stacks := make([]ItemStack, count)
	for _, line := range lines[1:] {
		chars := strings.Split(line, "")
		for i := 0; i < count; i++ {
			char := chars[1+i*4]
			if char != " " {
				stacks[i].Push(char)
			}
		}
	}
	return stacks
}

func parseMoves(input string) []Move {
	lines := strings.Split(input, "\n")
	return lo.Map(lines, func(line string, _ int) Move {
		return parseMove(line)
	})
}

var moveRegexp = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseMove(line string) Move {
	matches := moveRegexp.FindAllStringSubmatch(line, -1)[0][1:]
	count, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])
	return Move{count, from - 1, to - 1}
}

func (day Day05) FirstChallenge() string {
	day = day.Clone()
	for _, move := range day.Moves {
		moveFirstChallenge(&day, move)
	}
	return readTopItems(day.Stacks)
}

func moveFirstChallenge(day *Day05, move Move) {
	for i := 0; i < move.Count; i++ {
		item := day.Stacks[move.From].Pop()
		day.Stacks[move.To].Push(item)
	}
}

func (day Day05) SecondChallenge() string {
	day = day.Clone()
	for _, move := range day.Moves {
		moveSecondChallenge(&day, move)
	}
	return readTopItems(day.Stacks)
}

func moveSecondChallenge(day *Day05, move Move) {
	length := len(day.Stacks[move.From])
	items := day.Stacks[move.From][length-move.Count:]
	day.Stacks[move.From] = day.Stacks[move.From][:length-move.Count]
	day.Stacks[move.To] = append(day.Stacks[move.To], items...)
}

func readTopItems(stacks []ItemStack) string {
	tops := lo.Map(stacks, func(stack ItemStack, _ int) string {
		return stack.Peek()
	})

	return strings.Join(tops, "")
}
