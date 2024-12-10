package day04

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(4, day{})
}

type day struct{}

func parse(f *os.File) ([][]rune, error) {
	var matrix [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, c := range line {
			row = append(row, c)
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}

func (d day) Part1(f *os.File) error {
	matrix, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sizeY := len(matrix)
	if sizeY == 0 {
		return fmt.Errorf("empty matrix")
	}
	sizeX := len(matrix[0])
	if sizeX == 0 {
		return fmt.Errorf("empty matrix")
	}

	log.Infof("Size: %dx%d", sizeX, sizeY)

	var sum int
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			sum += find(matrix, "XMAS", x, y)
			sum += find(matrix, "SAMX", x, y)
		}
	}
	log.Infof("Sum: %d", sum)

	return nil
}

func find(matrix [][]rune, needle string, x int, y int) int {
	var sum = 0
	sizeX := len(matrix[0])
	sizeY := len(matrix)

	// Horizontal
	if x+len(needle)-1 < sizeX {
		ss := string(matrix[y][x : x+len(needle)])
		if ss == needle {
			sum++
		}
	}

	// Vertical
	if y+len(needle)-1 < sizeY {
		ss := ""
		for i := 0; i < len(needle); i++ {
			ss += string(matrix[y+i][x])
		}
		if ss == needle {
			sum++
		}
	}

	// Diagonal
	if x+len(needle) <= sizeX && y+len(needle) <= sizeY {
		ss := ""
		for i := 0; i < len(needle); i++ {
			ss += string(matrix[y+i][x+i])
		}
		if ss == needle {
			sum++
		}
	}

	// Anti-Diagonal
	if x-len(needle)+1 >= 0 && y+len(needle) <= sizeX {
		ss := ""
		for i := 0; i < len(needle); i++ {
			ss += string(matrix[y+i][x-i])
		}
		if ss == needle {
			sum++
		}
	}

	return sum
}

func find2(matrix [][]rune, needle string, x int, y int) int {
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1
	needleRunes := []rune(needle)
	reversedNeedle := []rune(reverseString(needle))

	if len(needle) != 3 {
		log.Fatalf("Invalid needle")
	}
	if x-1 < 0 || y-1 < 0 {
		return 0
	}
	if x+1 > maxX || y+1 > maxY {
		return 0
	}

	if matrix[y][x] != rune(needle[1]) {
		return 0
	}

	diag1 := []rune{matrix[y-1][x-1], matrix[y][x], matrix[y+1][x+1]}
	diag2 := []rune{matrix[y-1][x+1], matrix[y][x], matrix[y+1][x-1]}

	if (slices.Equal(diag1, needleRunes) || slices.Equal(diag1, reversedNeedle)) &&
		(slices.Equal(diag2, needleRunes) || slices.Equal(diag2, reversedNeedle)) {
		return 1
	}
	return 0
}

func reverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func (d day) Part2(f *os.File) error {
	matrix, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	sizeY := len(matrix)
	if sizeY == 0 {
		return fmt.Errorf("empty matrix")
	}
	sizeX := len(matrix[0])
	if sizeX == 0 {
		return fmt.Errorf("empty matrix")
	}

	log.Infof("Size: %dx%d", sizeX, sizeY)

	var sum int
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			sum += find2(matrix, "MAS", x, y)
		}
	}
	log.Infof("Sum: %d", sum)
	return nil
}
