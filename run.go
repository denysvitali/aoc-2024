package main

import (
	"fmt"
	"os"
	"time"

	"github.com/denysvitali/aoc-2024/framework"
)

func runDay(day int) {
	d := framework.Registry.Get(day)
	if d == nil {
		log.Fatalf("day %d not found", args.Day)
		return
	}
	run(d, day, "example")
	run(d, day, "input")
}

func run(d framework.Day, day int, s string) {
	f, err := os.Open(fmt.Sprintf("day%02d/%s.txt", day, s))
	if err != nil {
		log.Fatalf("open %s file: %v", s, err)
	}
	benchmarkRun(func() {
		if err := d.Part1(f); err != nil {
			log.Fatalf("day %d - part 1 %s: %v", args.Day, s, err)
		}
	})
	_, _ = f.Seek(0, 0)
	benchmarkRun(func() {
		if err := d.Part2(f); err != nil {
			log.Fatalf("day %d - part 2 %s: %v", args.Day, s, err)
		}
	})
}

func benchmarkRun(f func()) {
	startTime := time.Now()
	f()
	took := time.Since(startTime)
	log.Infof("Took %s", took)
}
