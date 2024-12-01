package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/aoc-2024/framework"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/denysvitali/aoc-2024/day01"
)

var args struct {
	Day int `arg:"positional,required"`
}

var log = logrus.StandardLogger()

func main() {
	arg.MustParse(&args)

	d := framework.Registry.Get(args.Day)
	if d == nil {
		log.Fatalf("day %d not found", args.Day)
		return
	}

	run(d, "example")
	run(d, "input")
}

func run(d framework.Day, s string) {
	f, err := os.Open(fmt.Sprintf("day%02d/%s.txt", args.Day, s))
	if err != nil {
		log.Fatalf("open %s file: %v", s, err)
	}
	if err := d.Part1(f); err != nil {
		log.Fatalf("day %d - part 1 %s: %v", args.Day, s, err)
	}
	f.Seek(0, 0)
	if err := d.Part2(f); err != nil {
		log.Fatalf("day %d - part 2 %s: %v", args.Day, s, err)
	}
}

func getFile(name string) *os.File {
	f, err := os.Open(fmt.Sprintf("day%02d/%s.txt", args.Day, name))
	if err != nil {
		log.Fatalf("open example file: %v", err)
	}
	return f
}
