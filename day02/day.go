package day02

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
	framework.Registry.Register(2, day02{})
}

type day02 struct {
}

func parse(f *os.File) ([][]int64, error) {
	scanner := bufio.NewScanner(f)
	ints := make([][]int64, 0)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		intsString := strings.Split(txt, " ")
		intsInt := make([]int64, 0)
		for _, v := range intsString {
			parsedInt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("int: %w", err)
			}
			intsInt = append(intsInt, parsedInt)
		}
		ints = append(ints, intsInt)
	}
	return ints, nil
}

func (d day02) Part1(f *os.File) (int64, error) {
	ints, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	safe := 0
	for _, v := range ints {
		if isSafe(v) {
			safe++
		}
	}
	return int64(safe), nil
}

func isSafe(v []int64) bool {
	dir := v[1]-v[0] > 0
	for i := 1; i < len(v); i++ {
		prev := v[i-1]
		curr := v[i]

		diff := curr - prev
		thisDir := diff > 0
		if thisDir != dir {
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

func (d day02) Part2(f *os.File) (int64, error) {
	ints, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	safe := 0
	for _, v := range ints {
		if isSafe(v) {
			safe++
		} else {
			for i := 0; i < len(v); i++ {
				// Try to delete each element one by one and check if it becomes safe
				newArr := make([]int64, 0)
				newArr = append(newArr, v[:i]...)
				newArr = append(newArr, v[i+1:]...)
				if isSafe(newArr) {
					safe++
					break
				}
			}
		}
	}

	log.Infof("Safe: %d", safe)
	return int64(safe), nil
}
