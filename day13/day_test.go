package day13

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	exampleF, err := os.Open("example.txt")
	if err != nil {
		t.Fatalf("TestPart1: %v", err)
	}
	p1Res, err := day{}.Part1(exampleF)
	if err != nil {
		t.Fatalf("Part1: %v", err)
	}

	assert.Equal(t, int64(480), p1Res)
}

func TestPart1Input(t *testing.T) {
	exampleF, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	p1Res, err := day{}.Part1(exampleF)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(36758), p1Res)
}

func TestPart2Example(t *testing.T) {
	exampleF, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	p2Res, err := day{}.Part2(exampleF)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1206, p2Res)
}

func TestPart2Input(t *testing.T) {
	inputF, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	p2Res, err := day{}.Part2(inputF)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(-1), p2Res)
}
