package day00

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day, ok := day.(Day00); !ok || day.input != "Merry Christmas" {
		t.FailNow()
	}
}

func TestFirstChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day.FirstChallenge() != "Merry Christmas" {
		t.FailNow()
	}
}

func TestSecondChallenge(t *testing.T) {
	t.Parallel()
	day := Load("Merry Christmas")
	if day.SecondChallenge() != "Merry Christmas" {
		t.FailNow()
	}
}
