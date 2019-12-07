package day

import "github.com/rafalrothenberger/adventofcode2019/internal/intcode"

import "github.com/rafalrothenberger/adventofcode2019/internal/input"

// Day5Task1 ...
func Day5Task1() interface{} {
	const (
		day      = 5
		task     = 1
		testCode = 0
	)

	in := make(chan int, 1)
	out := make(chan int, 10)
	in <- 1
	intcode.NewParser(input.Get(day, task, testCode).([]int), in, out).Parse()

	close(out)
	result := make([]int, 0)
	for c := range out {
		result = append(result, c)
	}
	return result
}

// Day5Task2 ...
func Day5Task2() interface{} {
	const (
		day      = 5
		task     = 1
		testCode = 0
	)

	in := make(chan int, 1)
	out := make(chan int, 10)
	in <- 5
	intcode.NewParser(input.Get(day, task, testCode).([]int), in, out).Parse()

	close(out)
	return <-out
}
