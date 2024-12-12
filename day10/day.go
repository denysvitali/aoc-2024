package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(10, day{})
}

type day struct{}

func parse(f *os.File) ([][]int, error) {
	var m [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var ints []int
		for _, c := range line {
			i, err := strconv.ParseInt(string(c), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing line: %w", err)
			}
			ints = append(ints, int(i))
		}
		m = append(m, ints)
	}
	return m, nil
}

func (d day) Part1(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}

	var sumScore int
	for y, row := range m {
		for x, c := range row {
			if c == 0 {
				score, _ := findPath(m, map[pos]struct{}{}, pos{x, y}, -1)
				sumScore += score
			}
		}
	}

	log.Infof("Sum score: %d", sumScore)
	return int64(sumScore), nil
}

type pos struct {
	x int
	y int
}

func (p pos) left() pos {
	return pos{p.x - 1, p.y}
}

func (p pos) right() pos {
	return pos{p.x + 1, p.y}
}

func (p pos) up() pos {
	return pos{p.x, p.y - 1}
}

func (p pos) down() pos {
	return pos{p.x, p.y + 1}
}

func (p pos) oob(m [][]int) bool {
	return p.x < 0 || p.x >= len(m[0]) || p.y < 0 || p.y >= len(m)
}

// findPath finds the linear path from the starting point to 9
func findPath(m [][]int, visited map[pos]struct{}, p pos, lastNumber int) (int, int) {
	if p.oob(m) {
		return 0, 0
	}
	curr := m[p.y][p.x]
	if curr != lastNumber+1 {
		// Invalid position
		return 0, 0
	}
	if curr == 9 {
		// Found the top!
		if _, ok := visited[p]; ok {
			// Already reached this top
			return 0, 1
		}
		visited[p] = struct{}{}
		return 1, 1
	}

	leftS, leftR := findPath(m, visited, p.left(), curr)
	rightS, rightR := findPath(m, visited, p.right(), curr)
	upS, upR := findPath(m, visited, p.up(), curr)
	downS, downR := findPath(m, visited, p.down(), curr)

	s := leftS + rightS + upS + downS
	r := leftR + rightR + upR + downR
	return s, r
}

func (d day) Part2(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}

	var sumRating int
	for y, row := range m {
		for x, c := range row {
			if c == 0 {
				_, r := findPath(m, map[pos]struct{}{}, pos{x, y}, -1)
				sumRating += r
			}
		}
	}
	return int64(sumRating), nil
}
