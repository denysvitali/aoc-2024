package framework

import "os"

type Day interface {
	Part1(f *os.File) error
	Part2(f *os.File) error
}
