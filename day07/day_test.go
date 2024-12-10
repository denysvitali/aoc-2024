package day07

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConcat(t *testing.T) {
	got := concat(element{Value: 123}, element{Value: 456})
	want := element{Value: 123456}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatSingleDigit(t *testing.T) {
	got := concat(element{Value: 1}, element{Value: 2})
	want := element{Value: 12}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatLeadingZeros(t *testing.T) {
	got := concat(element{Value: 100}, element{Value: 2})
	want := element{Value: 1002}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatZero(t *testing.T) {
	got := concat(element{Value: 0}, element{Value: 123})
	want := element{Value: 123}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatBothZero(t *testing.T) {
	got := concat(element{Value: 0}, element{Value: 0})
	want := element{Value: 0}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatLargeNumbers(t *testing.T) {
	got := concat(element{Value: 123456789}, element{Value: 987654321})
	want := element{Value: 123456789987654321}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConcatWithZeroResult(t *testing.T) {
	got := concat(element{Value: 0}, element{Value: 0})
	want := element{Value: 0}
	if cmp.Diff(got, want) != "" {
		t.Errorf("got %v want %v", got, want)
	}
}
