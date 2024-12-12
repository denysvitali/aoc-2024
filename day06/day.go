package day06

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(6, day{})
}

type day struct{}

func parse(f *os.File) ([][]rune, error) {
	scanner := bufio.NewScanner(f)
	var m [][]rune
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

func (p point) Add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

func (p point) OutOfBounds(maxX int, maxY int) bool {
	return p.x >= maxX || p.x < 0 || p.y >= maxY || p.y < 0
}

func (d day) Part1(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}
	currGuardPos := getGuardPos(m)
	_, visited := run(currGuardPos, m)
	return int64(len(visited)), nil
}

func isObstacle(p point, m [][]rune) bool {
	return m[p.y][p.x] == '#'
}

func (d day) Part2(f *os.File) (int64, error) {
	m, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}
	return int64(d.part2(m)), nil
}

func (d day) part2(m [][]rune) int {
	// Do a first run
	_, visited := run(getGuardPos(m), m)
	currGuardPos := getGuardPos(m)
	return countLoops(currGuardPos, m, visited)
}

func getGuardPos(m [][]rune) point {
	for y, row := range m {
		for x, c := range row {
			if c == '^' {
				return point{x, y}
			}
		}
	}
	return point{}
}

func (p point) isValid(maxX, maxY int) bool {
	return p.x >= 0 && p.x < maxX && p.y >= 0 && p.y < maxY
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var directions = [4]point{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}

func encodeState(x, y, dir int, maxX int) int {
	return (y*maxX+x)*4 + dir
}

func run(currGuardPos point, m [][]rune) (loop bool, visited map[point]struct{}) {
	maxX, maxY := len(m[0]), len(m)
	initialCapacity := (maxX * maxY) / 4

	visited = make(map[point]struct{}, initialCapacity)
	visitedStates := make([]bool, maxX*maxY*4) // Encode position+direction

	currDir := UP
	visited[currGuardPos] = struct{}{}
	for {
		nextPos := currGuardPos
		for {
			nextPos = currGuardPos.Add(directions[currDir])
			if !nextPos.isValid(maxX, maxY) || !isObstacle(nextPos, m) {
				break
			}
			currDir = nextDirection(currDir)
		}

		if !nextPos.isValid(maxX, maxY) {
			break
		}

		state := encodeState(nextPos.x, nextPos.y, currDir, maxX)
		if visitedStates[state] {
			loop = true
			break
		}

		currGuardPos = nextPos
		visited[currGuardPos] = struct{}{}
		visitedStates[state] = true
	}

	return loop, visited
}

func nextDirection(curr int) int {
	return (curr + 1) & 3 // Faster than modulo for powers of 2
}

func countLoops(currGuardPos point, m [][]rune, visited map[point]struct{}) int {
	loops := 0
	for k, _ := range visited {
		if m[k.y][k.x] == '.' && k != currGuardPos {
			m[k.y][k.x] = '#'
			isLoop, _ := run(currGuardPos, m)
			if isLoop {
				loops++
			}
			m[k.y][k.x] = '.'
		}
	}
	return loops
}
