package day07

import (
	"strconv"
	"strings"

	"github.com/paolosimone/advent-of-code-2022/advent/days"
	"github.com/samber/lo"
)

type NodeType int

const (
	Directory NodeType = iota
	File
)

type Node struct {
	Name     string
	Type     NodeType
	Parent   *Node
	Children map[string]*Node
	Size     int
}

func NewDir(name string) Node {
	return Node{
		Name:     name,
		Type:     Directory,
		Parent:   nil,
		Children: make(map[string]*Node),
		Size:     0,
	}
}

func NewFile(name string, size int) Node {
	return Node{
		Name:     name,
		Type:     File,
		Parent:   nil,
		Children: nil,
		Size:     size,
	}
}

type Day07 struct {
	Root Node
}

func Load(input string) days.Day {
	root := NewDir("/")

	current := &root
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); {
		current, i = parseCommand(&root, current, lines, i)
	}

	computeSizes(&root)
	return Day07{root}
}

func parseCommand(root *Node, node *Node, lines []string, start int) (*Node, int) {
	switch lines[start][2:4] {
	case "cd":
		next := cd(root, node, lines[start])
		return next, start + 1
	case "ls":
		return ls(node, lines, start)
	default:
		panic("unknown command")
	}
}

func cd(root *Node, node *Node, line string) *Node {
	tokens := strings.Split(line, " ")
	directory := tokens[len(tokens)-1]
	switch directory {
	case "/":
		return root
	case "..":
		return node.Parent
	default:
		return node.Children[directory]
	}
}

func ls(node *Node, lines []string, start int) (*Node, int) {
	index := start + 1
	for ; index < len(lines) && !strings.HasPrefix(lines[index], "$"); index++ {
		child := parseNode(lines[index])
		child.Parent = node
		node.Children[child.Name] = &child
	}
	return node, index
}

func parseNode(line string) Node {
	tokens := strings.Split(line, " ")
	name := tokens[1]
	switch tokens[0] {
	case "dir":
		return NewDir(name)
	default:
		size, _ := strconv.Atoi(tokens[0])
		return NewFile(name, size)
	}
}

func computeSizes(node *Node) {
	if node.Type == File {
		return
	}

	for _, child := range node.Children {
		computeSizes(child)
	}

	node.Size = lo.SumBy(lo.Values(node.Children), func(child *Node) int {
		return child.Size
	})
}

func (day Day07) FirstChallenge() string {
	result := sumSmallDirectoriesSize(&day.Root)
	return strconv.Itoa(result)
}

func sumSmallDirectoriesSize(node *Node) int {
	if node.Type == File {
		return 0
	}

	sum := lo.SumBy(lo.Values(node.Children), sumSmallDirectoriesSize)

	if node.Size <= 100_000 {
		sum += node.Size
	}

	return sum
}

func (day Day07) SecondChallenge() string {
	unused := 70_000_000 - day.Root.Size
	target := 30_000_000 - unused
	result, _ := findSmallestDeletableSize(&day.Root, target)
	return strconv.Itoa(result)
}

func findSmallestDeletableSize(node *Node, target int) (int, bool) {
	if node.Type == File || node.Size < target {
		return 0, false
	}

	candidates := lo.FilterMap(lo.Values(node.Children), func(child *Node, _ int) (int, bool) {
		return findSmallestDeletableSize(child, target)
	})

	candidates = append(candidates, node.Size)

	return lo.Min(candidates), true
}
