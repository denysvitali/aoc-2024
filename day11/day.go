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
	framework.Registry.Register(11, &day{})
}

type day struct {
	memoiz map[mem]int
}

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

func (d *day) Part1(f *os.File) error {
	d.memoiz = make(map[mem]int)
	stones, err := parse(f)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	sum := 0
	for _, s := range stones {
		sum += d.blink(s, 0, 25)
	}
	log.Infof("Final: %v", sum)
	return nil
}

func (d *day) singleBlink(stone int) (res []int) {
	if stone == 0 {
		return []int{1}
	}
	digits := getDigits(stone)
	if len(digits)%2 == 0 {
		return []int{
			toNumber(digits[0 : len(digits)/2]),
			toNumber(digits[len(digits)/2:]),
		}
	}
	return []int{stone * 2024}
}

type mem struct {
	stone  int
	amount int
}

func (d *day) blink(stone int, totalLen int, amount int) int {
	if amount == 0 {
		return 1
	}
	if v, ok := d.memoiz[mem{stone: stone, amount: amount}]; ok {
		return v
	}
	mLen := 0
	for _, s := range d.singleBlink(stone) {
		mLen += d.blink(s, totalLen, amount-1)
	}
	d.memoiz[mem{stone: stone, amount: amount}] = mLen
	return mLen
}

func (d *day) Part2(f *os.File) error {
	d.memoiz = make(map[mem]int)
	stones, err := parse(f)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	sum := 0
	for _, s := range stones {
		sum += d.blink(s, 0, 75)
	}
	log.Infof("Final: %v", sum)
	return nil
}
