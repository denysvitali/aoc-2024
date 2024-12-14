package day14

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/denysvitali/aoc-2024/framework"
)

var log = logrus.StandardLogger()

func init() {
	framework.Registry.Register(14, day{})
}

type day struct{}
type v2 struct {
	x, y int
}
type pv struct {
	pos v2
	vel v2
}

func parse(f *os.File) ([]pv, error) {
	var r []pv
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var p pv
		if _, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.pos.x, &p.pos.y, &p.vel.x, &p.vel.y); err != nil {
			return nil, fmt.Errorf("error parsing line: %w", err)
		}
		r = append(r, p)
	}
	return r, nil
}

func (d day) Part1(f *os.File) (int64, error) {
	r, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	robotLocations := map[v2][]v2{}
	for _, currRobot := range r {
		p := currRobot.pos
		if _, ok := robotLocations[p]; !ok {
			robotLocations[p] = []v2{}
		}
		robotLocations[p] = append(robotLocations[p], currRobot.vel)
	}

	maxX, maxY := 101, 103
	if len(r) == 12 {
		// Example
		maxX, maxY = 11, 7
	}
	// printMaxX, printMaxY := 11, 7
	//printLocations(robotLocations, printMaxX, printMaxY)

	for i := 0; i < 100; i++ {
		robotLocations = tick(robotLocations, maxX, maxY)
	}

	robotsPerQuad := map[int]int{}
	quadrants := getQuadrants(maxX, maxY)
	for i, quad := range quadrants {
		for y := quad.MinY; y < quad.MaxY; y++ {
			for x := quad.MinX; x < quad.MaxX; x++ {
				if v, ok := robotLocations[v2{x, y}]; ok {
					robotsPerQuad[i] += len(v)
				}
			}
		}
	}

	// printLocations(robotLocations, maxX, maxY)
	mult := 1
	for _, q := range robotsPerQuad {
		mult *= q
	}
	return int64(mult), nil
}

type quad struct {
	MinX, MinY, MaxX, MaxY int
}

func getQuadrants(x int, y int) []quad {
	midX := x / 2
	midY := y / 2

	quadrants := []quad{
		// Top-left quadrant
		{
			MinX: 0,
			MinY: 0,
			MaxX: midX,
			MaxY: midY,
		},
		// Top-right quadrant
		{
			MinX: midX + 1,
			MinY: 0,
			MaxX: x,
			MaxY: midY,
		},
		// Bottom-left quadrant
		{
			MinX: 0,
			MinY: midY + 1,
			MaxX: midX,
			MaxY: y,
		},
		// Bottom-right quadrant
		{
			MinX: midX + 1,
			MinY: midY + 1,
			MaxX: x,
			MaxY: y,
		},
	}

	return quadrants
}

func printLocations(locations map[v2][]v2, maxX int, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if v, ok := locations[v2{x, y}]; ok {
				fmt.Print(len(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printLocations2(locations map[v2][]v2, maxX int, maxY int, i int) {
	// Make the image 4 times bigger in both dimensions
	scaleFactor := 4
	baseImage := image.NewRGBA(image.Rect(0, 0, maxX*scaleFactor, maxY*scaleFactor))
	f, err := os.Create(fmt.Sprintf("out/%d.png", i))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Iterate through the original coordinates
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// Check if this location exists in the map
			isActive := false
			if _, ok := locations[v2{x, y}]; ok {
				isActive = true
			}

			// Fill a 4x4 block for each coordinate
			for dy := 0; dy < scaleFactor; dy++ {
				for dx := 0; dx < scaleFactor; dx++ {
					if isActive {
						baseImage.Set(x*scaleFactor+dx, y*scaleFactor+dy, color.RGBA{G: 255, A: 255})
					} else {
						baseImage.Set(x*scaleFactor+dx, y*scaleFactor+dy, color.RGBA{A: 255})
					}
				}
			}
		}
	}

	if err := jpeg.Encode(f, baseImage, nil); err != nil {
		panic(err)
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func tick(robotLocations map[v2][]v2, maxX, maxY int) map[v2][]v2 {
	newLocations := make(map[v2][]v2)
	for pos, robots := range robotLocations {
		for _, speed := range robots {
			p := pos
			p.x += speed.x
			p.y += speed.y
			p.x = mod(p.x, maxX)
			p.y = mod(p.y, maxY)
			if _, ok := newLocations[p]; !ok {
				newLocations[p] = []v2{}
			}

			if p.x < 0 || p.y < 0 || p.x >= maxX || p.y >= maxY {
				panic("out of bounds")
			}
			newLocations[p] = append(newLocations[p], speed)
		}
	}
	return newLocations
}

// TODO: Implement Part2
func (d day) Part2(f *os.File) (int64, error) {
	r, err := parse(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing file: %w", err)
	}
	if len(r) == 12 {
		// Example, skip
		return int64(-1), nil
	}
	robotLocations := map[v2][]v2{}
	for _, currRobot := range r {
		p := currRobot.pos
		if _, ok := robotLocations[p]; !ok {
			robotLocations[p] = []v2{}
		}
		robotLocations[p] = append(robotLocations[p], currRobot.vel)
	}

	maxX, maxY := 101, 103
	if len(r) == 12 {
		// Example
		maxX, maxY = 11, 7
	}
	// printMaxX, printMaxY := 11, 7
	//printLocations(robotLocations, printMaxX, printMaxY)

	for i := 0; i < 15000; i++ {
		robotLocations = tick(robotLocations, maxX, maxY)
		// Reset screen
		// printLocations2(robotLocations, maxX, maxY, i+1)

		robotsPerQuad := map[int]int{}
		quadrants := getQuadrants(maxX, maxY)
		for i, quad := range quadrants {
			for y := quad.MinY; y < quad.MaxY; y++ {
				for x := quad.MinX; x < quad.MaxX; x++ {
					if v, ok := robotLocations[v2{x, y}]; ok {
						robotsPerQuad[i] += len(v)
					}
				}
			}
		}

		// Detect if there are a lot more robots in a quadrant (compared to the other quadrants)
		// If so, there might be a hidden message
		tot := 0
		for _, q := range robotsPerQuad {
			tot += q
		}
		for _, q := range robotsPerQuad {
			if float32(q) >= float32(tot)*0.5 {
				printLocations2(robotLocations, maxX, maxY, i+1)
				fmt.Println("Possible message at t=%d", i+1)
			}
		}

	}

	return int64(-1), nil
}
