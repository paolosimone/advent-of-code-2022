package day07_test

import (
	"testing"

	. "github.com/paolosimone/advent-of-code-2022/advent/days/day07"
)

func TestFirstChallenge(t *testing.T) {
	day := Load(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

	if result := day.FirstChallenge(); result != "95437" {
		t.Fatal(result)
	}
}

func TestSecondChallenge(t *testing.T) {
	day := Load(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

	if result := day.SecondChallenge(); result != "24933642" {
		t.Fatal(result)
	}
}
