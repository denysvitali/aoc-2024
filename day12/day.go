package day12

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(12, day{})
}

type day struct{}

func parse(f *os.File) ([][]rune, error) {
	var m [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		m = append(m, []rune(line))
	}
	return m, nil
}

type point struct {
	x, y int
}

func (p point) oob(m [][]rune) bool {
	return p.x < 0 || p.y < 0 || p.x >= len(m[0]) || p.y >= len(m)
}

func (p point) top() point {
	return point{p.x, p.y - 1}
}

func (p point) bottom() point {
	return point{p.x, p.y + 1}
}

func (p point) left() point {
	return point{p.x - 1, p.y}
}

func (p point) right() point {
	return point{p.x + 1, p.y}
}

func (p point) add(o point) point {
	return point{p.x + o.x, p.y + o.y}
}

func (d day) Part1(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	visited := make(map[rune]map[point]struct{})
	t := 0
	for y, row := range m {
		for x, cell := range row {
			p := point{x, y}
			if v, ok := visited[cell]; ok {
				if _, ok := v[p]; ok {
					continue
				}
			} else {
				visited[cell] = make(map[point]struct{})
			}
			a, f := checkRegions(m, p, visited[cell], cell, 0, 0)
			t += len(a) * f
		}
	}
	return int64(t), nil
}

func checkRegions(m [][]rune, p point, visited map[point]struct{}, cell rune, a int, f int) (map[point]struct{}, int) {
	if _, ok := visited[p]; ok {
		return map[point]struct{}{}, 0
	}
	if p.oob(m) {
		return map[point]struct{}{}, 1
	}
	if m[p.y][p.x] != cell {
		return map[point]struct{}{}, 1
	}
	outMap := map[point]struct{}{}
	outMap[p] = struct{}{}
	visited[p] = struct{}{}
	a++
	al, fl := checkRegions(m, p.left(), visited, cell, 0, 0)
	ar, fr := checkRegions(m, p.right(), visited, cell, 0, 0)
	at, ft := checkRegions(m, p.top(), visited, cell, 0, 0)
	ab, fb := checkRegions(m, p.bottom(), visited, cell, 0, 0)
	for _, v := range []map[point]struct{}{al, ar, at, ab} {
		for k, _ := range v {
			outMap[k] = struct{}{}
		}
	}
	f += fl + fr + ft + fb
	return outMap, f
}

func (d day) Part2(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	visited := make(map[rune]map[point]struct{})
	t := 0
	for y, row := range m {
		for x, cell := range row {
			p := point{x, y}
			if v, ok := visited[cell]; ok {
				if _, ok := v[p]; ok {
					continue
				}
			} else {
				visited[cell] = make(map[point]struct{})
			}
			a, _ := checkRegions(m, p, visited[cell], cell, 0, 0)
			corners := findCorners(m, a)
			t += len(a) * corners
		}
	}
	return int64(t), nil
}

func rAt(m [][]rune, p point) rune {
	if p.oob(m) {
		return -1
	}
	return m[p.y][p.x]
}

var cornerToOrtho = map[point][]point{
	{-1, -1}: {{0, -1}, {-1, 0}}, // Top Left
	{1, -1}:  {{1, 0}, {0, -1}},  // Top Right
	{1, 1}:   {{1, 0}, {0, 1}},   // Bottom Right
	{-1, 1}:  {{-1, 0}, {0, 1}},  // Bottom Left
}

func findCorners(m [][]rune, region map[point]struct{}) (corners int) {
	for p, _ := range region {
		x := rAt(m, p)
		for k, mCorners := range cornerToOrtho {
			c := rAt(m, p.add(k))
			p1 := rAt(m, p.add(mCorners[0]))
			p2 := rAt(m, p.add(mCorners[1]))

			// Interior corner
			if x != p1 && x != p2 {
				corners++
			}

			// Exterior corner
			if x == p1 && x == p2 && x != c {
				corners++
			}
		}
	}
	return
}
