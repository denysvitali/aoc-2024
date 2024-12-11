package day11

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(11, day{})
}

type day struct{}

func parse(f *os.File) ([]int, error) {
	var ints []int
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	intsString := strings.Split(lines[0], " ")
	for _, s := range intsString {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing line: %w", err)
		}
		ints = append(ints, int(i))
	}
	return ints, nil
}

func reverse(s []int) []int {
	var reversed []int
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}

func getDigits(i int) []int {
	var digits []int
	for i > 0 {
		digits = append(digits, i%10)
		i /= 10
	}
	return reverse(digits)
}

func toNumber(digits []int) int {
	var n int
	for i, d := range reverse(digits) {
		n += d * int(math.Pow10(i))
	}
	return n
}

func (d day) Part1(f *os.File) error {
	stones, err := parse(f)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones, i)
	}
	log.Infof("Final: %v", len(stones))
	return nil
}

func blink(stones []int, i int) []int {
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}
		d := getDigits(stone)
		if len(d)%2 == 0 {
			newStones = append(newStones, toNumber(d[0:len(d)/2]))
			newStones = append(newStones, toNumber(d[len(d)/2:]))
			continue
		}
		newStones = append(newStones, stone*2024)
	}
	stones = newStones
	if i < 3 {
		log.Infof("[%d]: %v", i+1, stones)
	}
	return stones
}

func (d day) Part2(f *os.File) error {
	stones, err := parse(f)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	for i := 0; i < 75; i++ {
		stones = blink(stones, i)
	}
	log.Infof("Final: %v", len(stones))
	return nil

}
