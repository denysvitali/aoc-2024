package main

import (
	"fmt"
	"os"
	"time"

	"github.com/denysvitali/aoc-2024/framework"
)

func runDay(day int, part int) {
	d := framework.Registry.Get(day)
	if d == nil {
		log.Fatalf("day %d not found", args.Day)
		return
	}
	run(d, day, part, "example")
	run(d, day, part, "input")
}

func run(d framework.Day, day int, part int, s string) {
	f, err := os.Open(fmt.Sprintf("day%02d/%s.txt", day, s))
	if err != nil {
		log.Fatalf("open %s file: %v", s, err)
	}
	if part == 1 || part == 0 {
		benchmarkRun(f, d.Part1, day, s, 1)
	}
	if part == 0 {
		_, _ = f.Seek(0, 0)
	}
	if part == 2 || part == 0 {
		benchmarkRun(f, d.Part2, day, s, 2)
	}
}

func benchmarkRun(file *os.File, fn func(file *os.File) (int64, error), day int, s string, part int) {
	startTime := time.Now()
	v, err := fn(file)
	took := time.Since(startTime)
	if err != nil {
		log.Fatalf("error running part %d: %v", part, err)
	}
	log.Infof("Day %02d - Part %d - %s\tResult: %d\t\tTook: %s",
		day, part, s, v, took.String(),
	)
}
