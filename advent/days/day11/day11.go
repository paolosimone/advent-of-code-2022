package day11

import (
	"sort"
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type (
	ID        = int
	Item      = int
	Operation = func(Item) Item
)

type Monkey struct {
	ID        ID
	Items     []Item
	Operation Operation
	Throw     Throw
}

type Throw struct {
	Test  Item
	True  ID
	False ID
}

func (throw *Throw) Next(item Item) ID {
	return lo.Ternary(item%throw.Test == 0, throw.True, throw.False)
}

type Day11 struct {
	Monkeys []Monkey
}

func Load(input string) days.Day {
	blocks := strings.Split(input, "\n\n")
	monkeys := lo.Map(blocks, parseMonkey)
	return Day11{monkeys}
}

func parseMonkey(block string, index int) Monkey {
	lines := strings.Split(block, "\n")
	items := parseItems(lines[1])
	operation := parseOperation(lines[2])
	throw := parseThrow(lines[3:])
	return Monkey{index, items, operation, throw}
}

func parseItems(line string) []Item {
	right := strings.Split(line, ": ")[1]
	tokens := strings.Split(right, ", ")
	return lo.Map(tokens, func(token string, _ int) Item {
		item, _ := strconv.Atoi(token)
		return item
	})
}

func parseOperation(line string) Operation {
	right := strings.Split(line, "new = ")[1]
	tokens := strings.Split(right, " ")
	op := parseBinaryOp(tokens[1])
	value, _ := strconv.Atoi(tokens[2])
	return func(old Item) Item {
		right := lo.Ternary(tokens[2] == "old", old, value)
		return op(old, right)
	}
}

func parseBinaryOp(op string) func(Item, Item) Item {
	switch op {
	case "+":
		return func(a Item, b Item) Item { return a + b }
	case "-":
		return func(a Item, b Item) Item { return a - b }
	case "*":
		return func(a Item, b Item) Item { return a * b }
	case "/":
		return func(a Item, b Item) Item { return a / b }
	default:
		panic("unsupported operation")
	}
}

func parseThrow(lines []string) Throw {
	token := strings.Split(lines[0], "divisible by ")[1]
	test, _ := strconv.Atoi(token)

	token = strings.Split(lines[1], "throw to monkey ")[1]
	ifTrue, _ := strconv.Atoi(token)

	token = strings.Split(lines[2], "throw to monkey ")[1]
	ifFalse, _ := strconv.Atoi(token)

	return Throw{test, ifTrue, ifFalse}
}

func (day Day11) FirstChallenge() string {
	relief := func(item Item) Item {
		return item / 3
	}

	result := simulate(day.Monkeys, relief, 20)
	return strconv.Itoa(result)
}

func (day Day11) SecondChallenge() string {
	common := 1
	for _, monkey := range day.Monkeys {
		common *= monkey.Throw.Test
	}

	relief := func(item Item) Item {
		return item % common
	}

	result := simulate(day.Monkeys, relief, 10_000)
	return strconv.Itoa(result)
}

func simulate(monkeys []Monkey, relief Operation, rounds int) int {
	inspections := make([]int, len(monkeys))

	items := make([][]Item, len(monkeys))
	for id := 0; id < len(items); id++ {
		items[id] = append(items[id], monkeys[id].Items...)
	}

	for round := 0; round < rounds; round++ {
		for id, monkey := range monkeys {
			for _, item := range items[id] {
				inspections[id]++
				item = relief(monkey.Operation(item))
				next := monkey.Throw.Next(item)
				items[next] = append(items[next], item)
			}
			items[id] = make([]int, 0)
		}
	}

	sort.Ints(inspections)
	lo.Reverse(inspections)
	return inspections[0] * inspections[1]
}
