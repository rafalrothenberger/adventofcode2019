package day

import (
	"github.com/rafalrothenberger/adventofcode2019/internal/input"
	"github.com/rafalrothenberger/adventofcode2019/internal/intcode"
)

// Day2Task1 ...
func Day2Task1() interface{} {
	const day = 2
	const task = 1
	const testCode = 0

	output := intcode.NewParser(input.Get(day, task, testCode).([]int), nil, nil).Parse()

	return output[0]
}

// Day2Task2 ...
func Day2Task2() interface{} {
	const day = 2
	const task = 2
	const testCode = 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			in := input.Get(day, task, testCode).([]int)

			in[1] = i
			in[2] = j

			output := intcode.NewParser(in, nil, nil).Parse()
			if output[0] == 19690720 {
				return 100*i + j
			}
		}
	}
	panic("failed")
}
