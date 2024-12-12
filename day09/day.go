package day9

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
	framework.Registry.Register(9, day{})
}

type day struct{}

func parse(f *os.File) ([]int, error) {
	var ints []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		for _, r := range line {
			i, err := strconv.ParseInt(string(r), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing rune %v: %w", r, err)
			}
			ints = append(ints, int(i))

		}
	}
	return ints, nil
}

type file struct {
	id     int
	blocks int
}

func (d day) Part1(f *os.File) (int64, error) {
	ints, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}

	var fs []int
	if len(ints)%2 != 1 {
		return 0, fmt.Errorf("even input")
	}

	isFile := true
	counter := 0
	for _, v := range ints {
		if isFile {
			fs = append(fs, repeatInt(counter, v)...)
			counter++
		} else {
			fs = append(fs, repeatInt(-1, v)...)
		}
		isFile = !isFile
	}

	arrangeBlocks(fs)
	return calculateChecksum(fs), nil
}

func calculateChecksum(fs []int) int64 {
	var checksum int64
	for id, v := range fs {
		if v == -1 {
			continue
		}
		checksum += int64(id * v)
	}
	return checksum
}

func expandFs(fs []file) []int {
	var out []int
	for _, f := range fs {
		out = append(out, repeatInt(f.id, f.blocks)...)
	}
	return out
}

func arrangeBlocks(fs []int) {
	leftC, rightC := 0, len(fs)-1
	for leftC < rightC {
		for leftC < rightC && fs[leftC] != -1 {
			leftC++
		}
		for leftC < rightC && fs[rightC] == -1 {
			rightC--
		}
		if leftC < rightC {
			fs[leftC], fs[rightC] = fs[rightC], fs[leftC]
			leftC++
			rightC--
		}
	}
}

func repeatInt(x int, times int) []int {
	var out []int
	for i := 0; i < times; i++ {
		out = append(out, x)
	}
	return out
}

func arrangeBlocksFrag(fs []file, total int) {
	maxFileId := (total - 1) / 2
	log.Infof("Max file ID: %d", maxFileId)

	for id := maxFileId; id >= 0; id-- {
		foundIdx := -1
		var foundFile *file
		for i, f := range fs {
			if f.id == id {
				foundIdx = i
				foundFile = &f
				break
			}
		}
		if foundFile == nil {
			log.Fatalf("File %d not found", id)
			return
		}

		// Find the first empty space
		for i, v := range fs {
			if v.id == -1 && v.blocks >= foundFile.blocks {
				// We only swap if the index of the spot is less than the index of the file
				if i > foundIdx {
					continue
				}

				newEmptyBlock := file{-1, v.blocks - foundFile.blocks}
				v.blocks = foundFile.blocks
				fs[i], fs[foundIdx] = fs[foundIdx], v

				// We place the new empty block _after_ fs[i]
				if newEmptyBlock.blocks > 0 {
					fs = append(fs[:i+1], append([]file{newEmptyBlock}, fs[i+1:]...)...)
				}
				break
			}
		}

	}
}

type block struct {
	content int
	size    int
}

func (d day) Part2(f *os.File) (int64, error) {
	ints, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}

	var fs []file
	if len(ints)%2 != 1 {
		return 0, fmt.Errorf("even input")
	}

	isFile := true
	counter := 0
	for _, v := range ints {
		if isFile {
			fs = append(fs, file{counter, v})
			counter++
		} else {
			fs = append(fs, file{-1, v})
		}
		isFile = !isFile
	}

	arrangeBlocksFrag(fs, len(ints))

	return calculateChecksum(expandFs(fs)), nil

}
