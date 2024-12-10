package day06

import (
	"os"
	"runtime"
	"testing"
)

func clone2DArray(in [][]rune) [][]rune {
	out := make([][]rune, len(in))
	for i := range in {
		out[i] = make([]rune, len(in[i]))
		copy(out[i], in[i])
	}
	return out
}

func BenchmarkDay_Part2(b *testing.B) {
	f, err := os.Open("./input.txt")
	if err != nil {
		b.Fatalf("open input: %v", err)
	}
	m, err := parse(f)
	if err != nil {
		b.Fatalf("parse: %v", err)
	}
	d := day{}
	for i := 0; i < b.N; i++ {
		var sink interface{}
		b.ReportAllocs()
		clonedInput := clone2DArray(m)
		b.Run("Part2", func(b *testing.B) {
			result := d.part2(clonedInput)
			sink = result
		})
		runtime.KeepAlive(sink)
	}
}
