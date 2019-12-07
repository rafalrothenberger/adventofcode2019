package main

import (
	"os"

	"github.com/rafalrothenberger/adventofcode2019/internal/day"
)

func main() {
	dayArg := os.Args[1]
	taskArg := os.Args[2]
	day.Run(dayArg, taskArg)
}
