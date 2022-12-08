package day08

import (
	"math"
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type (
	Forest = [][]int
)

func makeMatrix[T any](rows int, cols int) [][]T {
	matrix := make([][]T, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]T, cols)
	}
	return matrix
}

func columnToArray[T any](matrix [][]T, col int) []T {
	length := len(matrix)
	array := make([]T, length)
	for i := 0; i < length; i++ {
		array[i] = matrix[i][col]
	}
	return array
}

func copyReverse[T any](array []T) []T {
	result := make([]T, len(array))
	copy(result, array)
	return lo.Reverse(result)
}

type Day08 struct {
	Forest Forest
}

func Load(input string) days.Day {
	lines := strings.Split(input, "\n")
	forest := lo.Map(lines, func(line string, _ int) []int {
		trees := strings.Split(line, "")
		return lo.Map(trees, func(tree string, _ int) int {
			height, _ := strconv.Atoi(tree)
			return height
		})
	})
	return Day08{forest}
}

func (day Day08) FirstChallenge() string {
	rows, cols := len(day.Forest), len(day.Forest[0])

	visible := makeMatrix[bool](rows, cols)

	// rows
	for i := 1; i < rows-1; i++ {
		left := computeVisibility(day.Forest[i])
		right := lo.Reverse(computeVisibility(copyReverse(day.Forest[i])))
		for j := 1; j < cols-1; j++ {
			visible[i][j] = left[j] || right[j]
		}
	}

	// columns
	for j := 1; j < cols-1; j++ {
		column := columnToArray(day.Forest, j)
		up := computeVisibility(column)
		down := lo.Reverse(computeVisibility(lo.Reverse(column)))
		for i := 1; i < rows-1; i++ {
			visible[i][j] = visible[i][j] || up[i] || down[i]
		}
	}

	internal := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if visible[i][j] {
				internal++
			}
		}
	}

	borders := 2*(rows+cols) - 4
	result := borders + internal

	return strconv.Itoa(result)
}

func computeVisibility(line []int) []bool {
	length := len(line)
	visible := make([]bool, length)

	visible[0] = true
	max := line[0]
	for i := 1; i < len(line)-1; i++ {
		if line[i] > max {
			visible[i] = true
			max = line[i]
		}
	}

	return visible
}

func (day Day08) SecondChallenge() string {
	rows, cols := len(day.Forest), len(day.Forest[0])
	views := makeMatrix[View](rows, cols)

	// rows
	for i := 1; i < rows-1; i++ {
		left := computeViewDistance(day.Forest[i])
		right := lo.Reverse(computeViewDistance(copyReverse(day.Forest[i])))
		for j := 1; j < cols-1; j++ {
			views[i][j].Left = left[j]
			views[i][j].Right = right[j]
		}
	}

	// columns
	for j := 1; j < cols-1; j++ {
		column := columnToArray(day.Forest, j)
		up := computeViewDistance(column)
		down := lo.Reverse(computeViewDistance(lo.Reverse(column)))
		for i := 1; i < rows-1; i++ {
			views[i][j].Up = up[i]
			views[i][j].Down = down[i]
		}
	}

	result := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if score := views[i][j].Score(); score > result {
				result = score
			}
		}
	}

	return strconv.Itoa(result)
}

type View struct {
	Left  int
	Right int
	Up    int
	Down  int
}

func (view *View) Score() int {
	return view.Left * view.Right * view.Up * view.Down
}

type Tallest struct {
	Index  int
	Height int
}

type Stack[T any] []T

func (stack *Stack[T]) Peek() T {
	return (*stack)[len(*stack)-1]
}

func (stack *Stack[T]) Pop() T {
	top := stack.Peek()
	*stack = (*stack)[:len(*stack)-1]
	return top
}

func (stack *Stack[T]) Push(item T) {
	*stack = append(*stack, item)
}

func computeViewDistance(line []int) []int {
	length := len(line)
	view := make([]int, length)

	stack := make(Stack[Tallest], 0)
	stack.Push(Tallest{Index: 0, Height: math.MaxInt})

	for i := 1; i < length-1; i++ {
		height := line[i]
		for stack.Peek().Height < height {
			stack.Pop()
		}
		view[i] = i - stack.Peek().Index
		stack.Push(Tallest{Index: i, Height: height})
	}

	return view
}
