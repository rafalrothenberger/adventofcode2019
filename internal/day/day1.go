package day

import (
	"github.com/rafalrothenberger/adventofcode2019/internal/input"
)

// Day1Task2 ...
func Day1Task2() interface{} {
	const day = 1
	const task = 2
	const testCode = 0

	modules := input.Get(day, task, testCode).([]int)

	var fuels []int

	for _, module := range modules {
		sum := 0
		x := (module / 3) - 2
		if x < 0 {
			x = 0
		}

		for x > 0 {
			sum += x
			x = (x / 3) - 2
			if x < 0 {
				x = 0
			}
		}
		fuels = append(fuels, sum)
	}

	sum := 0
	for _, f := range fuels {
		sum += f
	}

	return sum
}

// Day1Task1 ...
func Day1Task1() interface{} {
	const day = 1
	const task = 1
	const testCode = 0

	modules := input.Get(day, task, testCode).([]int)

	sum := 0
	for _, module := range modules {
		x := (module / 3) - 2
		if x < 0 {
			x = 0
		}

		sum += x
	}

	return sum
}
