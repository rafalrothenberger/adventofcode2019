package day

import "github.com/rafalrothenberger/adventofcode2019/internal/input"

// Day8Task1 ...
func Day8Task1() interface{} {
	const (
		day      = 8
		task     = 1
		testCode = 0
	)

	input := input.Get(day, task, testCode).(string)

	no0 := 500
	res := 0
	for i := 0; i < len(input)-25*6; i += 25 * 6 {

		var zeros, ones, twos int
		for _, c := range input[i : i+25*6] {
			switch c {
			case '0':
				zeros++
			case '1':
				ones++
			case '2':
				twos++
			}
		}
		if zeros < no0 {
			no0 = zeros
			res = ones * twos
		}

	}
	return res
}

// Day8Task2 ...
func Day8Task2() interface{} {
	const (
		day         = 8
		task        = 2
		testCode    = 0
		layerLength = 150
		width       = 25
	)

	input := input.Get(day, task, testCode).(string)

	layers := make([]string, 0)
	for i := 0; i <= len(input)-layerLength; i += layerLength {
		layers = append(layers, input[i:i+layerLength])
	}

	res := ""
	x := ""
	for i := 0; i < layerLength; i++ {
		for _, layer := range layers {
			x = string(layer[i])
			if layer[i] != '2' {
				break
			}
		}
		res += x
	}

	var result string
	for i := 0; i <= len(res)-width; i += width {
		for j := 0; j < width; j++ {
			if res[i+j] == '1' {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}
