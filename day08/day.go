package day08

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(8, day{})
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

type pos struct {
	x, y int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p *pos) Vector(q pos) pos {
	dx := p.x - q.x
	dy := p.y - q.y
	return pos{dx, dy}
}

func (p *pos) Magnitude() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p *pos) Plus(o pos) pos {
	return pos{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

func (p *pos) Times(s int) pos {
	return pos{
		x: p.x * s,
		y: p.y * s,
	}
}

func (p *pos) Minus(o pos) pos {
	return p.Plus(pos{
		x: -o.x,
		y: -o.y,
	})
}

func (d day) Part1(f *os.File) error {
	m, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	var antennas map[rune][]pos
	maxY := len(m)
	maxX := len(m[0])

	antennas = printMap(m)

	// Find pairs of antennas that are close nearby
	for _, a := range antennas {
		for i, _ := range a {
			for j := i + 1; j < len(a); j++ {
				// Check distance
				v1 := a[i].Vector(a[j])

				for k := 0; k < 1000; k++ {
					newPos := a[i].Plus(v1.Times(k))
					newPos2 := a[j].Minus(v1.Times(k))
					drawPos(m, newPos, maxX, maxY)
					drawPos(m, newPos2, maxX, maxY)
					// TODO: find a better way to break instead of hardcoding 1000 iterations
				}
			}
		}
	}

	unique := 0
	for _, r := range m {
		for _, c := range r {
			if c == '#' {
				unique++
			}
		}
	}

	log.Infof("Unique: %d", unique)
	return nil
}

func drawPos(m [][]rune, newPos pos, maxX int, maxY int) {
	if newPos.x >= maxX || newPos.y >= maxY || newPos.x < 0 || newPos.y < 0 {
		return
	}
	m[newPos.y][newPos.x] = '#'
}

func printMap(m [][]rune) map[rune][]pos {
	var antennas map[rune][]pos
	for y, row := range m {
		fmt.Printf("%02d: ", y)
		for x, c := range row {
			fmt.Printf("%c", c)
			if c != '.' {
				if antennas == nil {
					antennas = make(map[rune][]pos)
				}
				antennas[c] = append(antennas[c], pos{x, y})
			}
		}
		fmt.Println()
	}
	return antennas
}

func (d day) Part2(f *os.File) error {
	// m, err := parse(f)
	return nil
}
