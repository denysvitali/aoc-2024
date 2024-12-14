package day14

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestGetQuadrants(t *testing.T) {
	want := []quad{
		{0, 0, 5, 3},
		{6, 0, 11, 3},
		{0, 4, 5, 7},
		{6, 4, 11, 7},
	}
	got := getQuadrants(11, 7)
	if diff := cmp.Diff(want, got, cmpopts.EquateComparable()); diff != "" {
		t.Fatalf("unexpected quadrants (-want +got):\n%s", diff)
	}
}
func TestPart1Example(t *testing.T) {
	exampleF, err := os.Open("example.txt")
	if err != nil {
		t.Fatalf("TestPart1: %v", err)
	}
	p1Res, err := day{}.Part1(exampleF)
	if err != nil {
		t.Fatalf("Part1: %v", err)
	}

	assert.Equal(t, int64(12), p1Res)
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

	assert.Equal(t, int64(215987200), p1Res)
}

func TestPart2Input(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	p1Res, err := day{}.Part2(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(215987200), p1Res)
}
