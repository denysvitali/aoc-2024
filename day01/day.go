package day01

import (
	"bufio"
	"fmt"
	"github.com/denysvitali/aoc-2024/framework"
	"github.com/sirupsen/logrus"
	"os"
	"sort"
)

var log = logrus.StandardLogger()

type day01 struct {
}

type sortInts []int64

func (s sortInts) Len() int {
	return len(s)
}

func (s sortInts) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortInts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = sortInts{}

func (d day01) Part1(f *os.File) error {
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		var a, b int64
		if _, err := fmt.Sscanf(txt, "%d\t\t%d", &a, &b); err != nil {
			return fmt.Errorf("parse: %w", err)
		}
		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}

	sort.Sort(sortInts(arr1))
	sort.Sort(sortInts(arr2))

	distances := make([]int64, len(arr1))

	sum := int64(0)
	for i, a := range arr1 {
		distances[i] = arr2[i] - a
		if distances[i] < 0 {
			distances[i] = -distances[i]
		}
		sum += distances[i]
	}

	log.Infof("Part 1: %d", sum)

	return nil
}
func (d day01) Part2(f *os.File) error {
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		var a, b int64
		if _, err := fmt.Sscanf(txt, "%d\t\t%d", &a, &b); err != nil {
			return fmt.Errorf("parse: %w", err)
		}
		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}

	m1 := map[int64]int{}
	for _, v := range arr2 {
		v2, ok := m1[v]
		if !ok {
			m1[v] = 1
		} else {
			m1[v] = v2 + 1
		}
	}

	simSum := int64(0)
	for _, v := range arr1 {
		v2, ok := m1[v]
		if !ok {
			continue
		}
		simSum = simSum + v*int64(v2)
	}
	log.Infof("Part 2: %d", simSum)

	return nil
}

func init() {
	framework.Registry.Register(1, day01{})
}
