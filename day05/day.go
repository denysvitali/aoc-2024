package day05

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(5, day{})
}

type day struct{}

type orderingRule struct {
	first  int64
	second int64
}

type instructions struct {
	orderingRules []orderingRule
	updates       [][]int64
}

func parse(f *os.File) (*instructions, error) {
	ins := instructions{}
	scanner := bufio.NewScanner(f)
	isOrderingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if isOrderingRules {
			if line == "" {
				isOrderingRules = false
				continue
			}
			var rule orderingRule
			_, err := fmt.Sscanf(line, "%d|%d", &rule.first, &rule.second)
			if err != nil {
				return nil, fmt.Errorf("parsing ordering rule: %w", err)
			}
			ins.orderingRules = append(ins.orderingRules, rule)
			continue
		}

		var update []int64
		for _, v := range strings.Split(line, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing update: %w", err)
			}
			update = append(update, i)
		}
		ins.updates = append(ins.updates, update)
	}

	return &ins, nil
}

func (d day) Part1(f *os.File) error {
	ins, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	orMap := map[int64][]int64{}
	for _, rule := range ins.orderingRules {
		if _, ok := orMap[rule.first]; !ok {
			orMap[rule.first] = []int64{}
		}
		orMap[rule.first] = append(orMap[rule.first], rule.second)
	}
	var validUpdates [][]int64
	for _, u := range ins.updates {
		seen := map[int64]struct{}{}
		ordered := true
	insupd:
		for _, e := range u {
			for _, lt := range orMap[e] {
				if _, ok := seen[lt]; ok {
					fmt.Println("Not ordered")
					ordered = false
					continue insupd
				}
			}
			seen[e] = struct{}{}
		}
		if ordered {
			validUpdates = append(validUpdates, u)
		}
	}

	sumOfMiddles := int64(0)
	for _, u := range validUpdates {
		sumOfMiddles += u[len(u)/2]
	}

	log.Infof("Sum of middle elements: %d", sumOfMiddles)
	return nil
}

func (d day) Part2(f *os.File) error {
	ins, err := parse(f)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	orMap := map[int64][]int64{}
	for _, rule := range ins.orderingRules {
		if _, ok := orMap[rule.first]; !ok {
			orMap[rule.first] = []int64{}
		}
		orMap[rule.first] = append(orMap[rule.first], rule.second)
	}
	var invalidUpdates [][]int64
	for _, u := range ins.updates {
		seen := map[int64]struct{}{}
		ordered := true
	insupd:
		for _, e := range u {
			fmt.Println(e, orMap[e])
			for _, lt := range orMap[e] {
				if _, ok := seen[lt]; ok {
					fmt.Println("Not ordered")
					ordered = false
					continue insupd
				}
			}
			seen[e] = struct{}{}
		}
		if !ordered {
			invalidUpdates = append(invalidUpdates, u)
		}
		fmt.Println("-----")
	}

	fixedUpdates := fixUpdates(invalidUpdates, orMap)
	sumOfMiddles := int64(0)
	for _, u := range fixedUpdates {
		sumOfMiddles += u[len(u)/2]
	}
	log.Infof("Sum of middle elements: %d", sumOfMiddles)
	return nil
}

type customSort struct {
	orMap map[int64][]int64
	u     []int64
}

func (c customSort) Len() int {
	return len(c.u)
}

func (c customSort) Less(i, j int) bool {
	v, ok := c.orMap[c.u[i]]
	if !ok {
		return false
	}
	for _, lt := range v {
		if lt == c.u[j] {
			return true
		}
	}
	return false
}

func (c customSort) Swap(i, j int) {
	c.u[i], c.u[j] = c.u[j], c.u[i]
}

var _ sort.Interface = customSort{}

func fixUpdates(updates [][]int64, orMap map[int64][]int64) [][]int64 {
	for _, u := range updates {
		cs := customSort{orMap: orMap, u: u}
		sort.Sort(cs)
		fmt.Println(cs.u)
	}
	return updates
}
