package day

import (
	"github.com/rafalrothenberger/adventofcode2019/internal/input"
	"github.com/rafalrothenberger/adventofcode2019/internal/intcode"
)

func Day9Task1() interface{} {
	const (
		day      = 9
		task     = 1
		testCode = 0
	)

	in := make(chan int, 1)
	in <- 1
	out := make(chan int, 100)
	intcode.NewParser(input.Get(day, task, testCode).([]int), in, out).Parse()

	close(out)
	res := make([]int, 0)
	for o := range out {
		res = append(res, o)
	}

	return res
}

func Day9Task2() interface{} {
	const (
		day      = 9
		task     = 2
		testCode = 0
	)

	in := make(chan int, 1)
	in <- 2
	out := make(chan int, 100)
	intcode.NewParser(input.Get(day, task, testCode).([]int), in, out).Parse()

	close(out)
	res := make([]int, 0)
	for o := range out {
		res = append(res, o)
	}

	return res
}
