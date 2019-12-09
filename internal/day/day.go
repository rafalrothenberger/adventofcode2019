package day

import "fmt"

var days = map[string]interface{}{
	"1_1": Day1Task1,
	"1_2": Day1Task2,
	"2_1": Day2Task1,
	"2_2": Day2Task2,
	"3_1": Day3Task1,
	"3_2": Day3Task2,
	"4_1": Day4Task1,
	"4_2": Day4Task2,
	"5_1": Day5Task1,
	"5_2": Day5Task2,
	"6_1": Day6Task1,
	"6_2": Day6Task2,
	"7_1": Day7Task1,
	"7_2": Day7Task2,
	"8_1": Day8Task1,
	"8_2": Day8Task2,
	"9_1": Day9Task1,
	"9_2": Day9Task2,
}

// Run ...
func Run(day, task string) {
	v, ok := days[fmt.Sprintf("%s_%s", day, task)]
	if !ok {
		panic("Day/task not found")
	}
	fmt.Println(v.(func() interface{})())
}
