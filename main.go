package main

import (
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"

	// <AUTOMATIC-IMPORT>
	_ "github.com/denysvitali/aoc-2024/day01"
	_ "github.com/denysvitali/aoc-2024/day02"
	_ "github.com/denysvitali/aoc-2024/day03"
	_ "github.com/denysvitali/aoc-2024/day04"
	_ "github.com/denysvitali/aoc-2024/day05"
	_ "github.com/denysvitali/aoc-2024/day06"
	_ "github.com/denysvitali/aoc-2024/day07"
	_ "github.com/denysvitali/aoc-2024/day08"
	_ "github.com/denysvitali/aoc-2024/day09"
	_ "github.com/denysvitali/aoc-2024/day10"
	_ "github.com/denysvitali/aoc-2024/day11"
	_ "github.com/denysvitali/aoc-2024/day12"
	_ "github.com/denysvitali/aoc-2024/day13"
	// </AUTOMATIC-IMPORT>
)

type DayCmd struct {
	Day  int `arg:"positional,required"`
	Part int `arg:"-p,--part"`
}

type GenerateCmd struct {
	Day int `arg:"positional,required"`
}

var args struct {
	Day      *DayCmd      `arg:"subcommand:day"`
	Generate *GenerateCmd `arg:"subcommand:generate"`
}

var log = logrus.StandardLogger()

func main() {
	arg.MustParse(&args)
	log.SetLevel(logrus.DebugLevel)

	if args.Generate != nil {
		generate(args.Generate.Day)
		return
	}
	if args.Day != nil {
		runDay(args.Day.Day, args.Day.Part)
		return
	}
	log.Fatalf("no command specified")
}

