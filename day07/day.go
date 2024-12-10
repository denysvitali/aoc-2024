package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(7, day{})
}

type day struct{}

type operation struct {
	res      uint64
	elements []uint64
}

func parse(f *os.File) ([]operation, error) {
	var operations []operation
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		s := strings.Split(line, ":")
		if len(s) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		res, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse res: %w", err)
		}
		elementsStr := strings.Split(s[1], " ")
		var elements []uint64
		for _, e := range elementsStr {
			if e == "" {
				continue
			}
			n, err := strconv.ParseInt(e, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parse element: %w", err)
			}
			elements = append(elements, uint64(n))
		}
		operations = append(operations, operation{res: uint64(res), elements: elements})
	}

	return operations, nil
}

type opFunc func(uint64, uint64) uint64

var ops = map[rune]opFunc{
	'+': func(a, b uint64) uint64 {
		return a + b
	},
	'*': func(a, b uint64) uint64 {
		return a * b
	},
	'|': func(a, b uint64) uint64 {
		if b == 0 {
			return a
		}
		// Find the number of digits in b
		digits := uint64(1)
		tmp := b
		for tmp >= 10 {
			digits++
			tmp /= 10
		}
		// Shift a left by the number of digits in b and add b
		return a*pow10(digits) + b
	},
}

func pow10(n uint64) uint64 {
	result := uint64(1)
	for i := uint64(0); i < n; i++ {
		result *= 10
	}
	return result
}

func (d day) Part1(f *os.File) error {
	m, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sum := uint64(0)
	for _, o := range m {
		if recEvaluate(o.elements, o.res, 1, o.elements[0], []rune{'+', '*'}) {
			sum += o.res
		}
	}

	log.Infof("Sum: %d", sum)
	return nil
}

func recEvaluate(elements []uint64, target uint64, idx int, value uint64, operators []rune) bool {
	if idx == len(elements) {
		return value == target
	}

	for _, op := range operators {
		newVal := ops[op](value, elements[idx])
		if newVal > target {
			continue
		}
		if recEvaluate(elements, target, idx+1, newVal, operators) {
			return true
		}
	}
	return false
}

func (d day) Part2(f *os.File) error {
	m, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sum := uint64(0)
	for _, o := range m {
		if recEvaluate(o.elements, o.res, 1, o.elements[0], []rune{'+', '*', '|'}) {
			sum += o.res
		}
	}
	log.Infof("Sum: %d", sum)
	return nil
}
