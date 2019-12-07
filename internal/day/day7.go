package day

import (
	"sync"

	"github.com/rafalrothenberger/adventofcode2019/internal/input"
	"github.com/rafalrothenberger/adventofcode2019/internal/intcode"
	"github.com/rafalrothenberger/adventofcode2019/internal/permutation"
)

// Day7Task2 ...
func Day7Task2() interface{} {
	const (
		day      = 7
		task     = 2
		testCode = 0
	)

	result := 0
	for _, p := range permutation.Gen([]int{5, 6, 7, 8, 9}) {
		toA := make(chan int, 2)
		toA <- p[0]
		toB := make(chan int, 2)
		toB <- p[1]
		toC := make(chan int, 2)
		toC <- p[2]
		toD := make(chan int, 2)
		toD <- p[3]
		toE := make(chan int, 2)
		toE <- p[4]

		toA <- 0

		var wg sync.WaitGroup
		pipes := [][]chan int{
			{toA, toB},
			{toB, toC},
			{toC, toD},
			{toD, toE},
			{toE, toA},
		}

		for _, pipe := range pipes {
			wg.Add(1)
			pipe := pipe
			go func() {
				defer wg.Done()
				intcode.NewParser(input.Get(day, task, testCode).([]int), pipe[0], pipe[1]).Parse()
			}()
		}
		wg.Wait()
		x := <-toA
		if x > result {
			result = x
		}
	}

	return result
}

// Day7Task1 ...
func Day7Task1() interface{} {
	const (
		day      = 7
		task     = 1
		testCode = 0
	)

	result := 0
	for _, p := range permutation.Gen([]int{0, 1, 2, 3, 4}) {
		input := input.Get(day, task, testCode).([]int)
		a := calcAmps(input, p)
		if a > result {
			result = a
		}
	}

	return result
}

func calcAmps(org []int, order []int) int {
	a := 0
	for _, o := range order {
		input := org
		in := make(chan int, 2)
		out := make(chan int, 1)
		in <- o
		in <- a
		intcode.NewParser(input, in, out).Parse()
		a = <-out
	}
	return a
}
