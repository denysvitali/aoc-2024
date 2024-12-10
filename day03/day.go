package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(3, day{})
}

type day struct {
}

func mustParseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

type op struct {
	op   string
	args []int64
}

func parse(f *os.File) ([]op, error) {
	scanner := bufio.NewScanner(f)
	ops := make([]op, 0)
	re := regexp.MustCompile(`(mul|don't|do)\((?:(\d+),(\d+)|)\)`)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}

		matches := re.FindAllStringSubmatch(txt, -1)
		if len(matches) == 0 {
			return nil, fmt.Errorf("invalid input: %s", txt)
		}

		for _, m := range matches {
			currOp := op{
				op: m[1],
			}
			if m[1] == "mul" {
				if m[2] != "" && m[3] != "" {
					currOp.args = []int64{mustParseInt(m[2]), mustParseInt(m[3])}
				} else {
					continue
				}
			}
			ops = append(ops, currOp)
		}

	}
	return ops, nil
}

func (d day) Part1(f *os.File) error {
	mul, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sum := int64(0)
	for _, v := range mul {
		if v.op == "mul" {
			sum += v.args[0] * v.args[1]
		}
	}
	log.Infof("Sum: %d", sum)
	return nil
}

func (d day) Part2(f *os.File) error {
	ops, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sum := int64(0)
	active := true
	for _, v := range ops {
		switch v.op {
		case "don't":
			active = false
		case "do":
			active = true
		case "mul":
			if active {
				sum += v.args[0] * v.args[1]
			}
		}
	}
	log.Infof("Sum: %d", sum)
	return nil
}
