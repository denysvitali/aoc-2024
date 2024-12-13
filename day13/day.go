package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/mat"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(13, day{})
}

type day struct{}

type coords struct {
	x, y int64
}

func parseButton(s string) (coords, error) {
	var c coords
	if _, err := fmt.Sscanf(s, "X+%d, Y+%d", &c.x, &c.y); err != nil {
		return c, err
	}
	return c, nil
}

func parsePrize(s string) (coords, error) {
	var c coords
	if _, err := fmt.Sscanf(s, "X=%d, Y=%d", &c.x, &c.y); err != nil {
		return c, err
	}
	return c, nil
}

func parse(f *os.File) ([][]coords, error) {
	var m [][]coords
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "Button A: ") {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		var cArr []coords
		b, err := parseButton(strings.TrimPrefix(line, "Button A: "))
		if err != nil {
			return nil, fmt.Errorf("error parsing button A: %w", err)
		}
		cArr = append(cArr, b)

		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected EOF")
		}
		line = scanner.Text()
		if !strings.HasPrefix(line, "Button B: ") {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		b, err = parseButton(strings.TrimPrefix(line, "Button B: "))
		if err != nil {
			return nil, fmt.Errorf("error parsing button B: %w", err)
		}
		cArr = append(cArr, b)

		// Prize
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected EOF")
		}
		line = scanner.Text()
		if !strings.HasPrefix(line, "Prize: ") {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		p, err := parsePrize(strings.TrimPrefix(line, "Prize: "))
		if err != nil {
			return nil, fmt.Errorf("error parsing prize: %w", err)
		}
		cArr = append(cArr, p)
		m = append(m, cArr)
	}
	return m, nil
}

func (d day) Part1(f *os.File) (int64, error) {
	machines, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	var tokens int64

	for _, m := range machines {

		a, b, err := solve(m)
		if err != nil {
			return 0, fmt.Errorf("error solving system of equations: %w", err)
		}

		if a < 0 && b < 0 {
			continue
		}

		if a < 0 || b < 0 {
			return 0, fmt.Errorf("invalid solution: a=%d, b=%d", a, b)
		}

		tokens += 3*a + b
	}
	return tokens, nil
}

func solve(m []coords) (int64, int64, error) {
	// This is a system of linear equations
	// m[0].x * a + m[1].x * b = m[2].x
	// m[0].y * a + m[1].y * b = m[2].y

	// but we also need to make sure that the solution is an integer and positive
	// a, b >= 0

	// We can solve this using matrices
	A := mat.NewDense(2, 2, []float64{float64(m[0].x), float64(m[1].x), float64(m[0].y), float64(m[1].y)})
	B := mat.NewVecDense(2, []float64{float64(m[2].x), float64(m[2].y)})
	var x mat.VecDense
	if err := x.SolveVec(A, B); err != nil {
		return 0, 0, err
	}

	// check if the solution is an integer
	epsilon := 1e-3 // Make epsilon smaller for more precise integer detection
	a := x.AtVec(0)
	b := x.AtVec(1)

	if math.Abs(a-math.Round(a)) > epsilon || math.Abs(b-math.Round(b)) > epsilon {
		return -1, -1, nil
	}

	aInt := int64(math.Round(a))
	bInt := int64(math.Round(b))

	// Verify the solution
	if aInt*m[0].x+bInt*m[1].x != m[2].x || aInt*m[0].y+bInt*m[1].y != m[2].y {
		log.Debugf("Solution verification failed: %d*%d + %d*%d = %d (expected %d)",
			aInt, m[0].x, bInt, m[1].x, aInt*m[0].x+bInt*m[1].x, m[2].x)
		return -1, -1, nil
	}

	return aInt, bInt, nil
}

func (d day) Part2(f *os.File) (int64, error) {
	machines, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	var tokens int64

	for _, m := range machines {
		m[2].x += 10000000000000
		m[2].y += 10000000000000

		a, b, err := solve(m)
		if err != nil {
			return 0, fmt.Errorf("error solving system of equations: %w", err)
		}

		if a < 0 && b < 0 {
			continue
		}

		if a < 0 || b < 0 {
			return 0, fmt.Errorf("invalid solution: a=%d, b=%d", a, b)
		}
		tokens += 3*a + b
	}
	return tokens, nil
}
