package day

import "github.com/rafalrothenberger/adventofcode2019/internal/input"

import "fmt"

import "github.com/rafalrothenberger/adventofcode2019/internal/intcode"

// Tile ...
type Tile struct {
	x, y  int
	color int
}

func (t Tile) String() string {
	return fmt.Sprintf("%d %d", t.x, t.y)
}

// Day11Task1 ...
func Day11Task1() interface{} {
	const (
		day      = 11
		task     = 1
		testCode = 0
	)
	input := input.Get(day, task, testCode).([]int)
	tiles := map[string]Tile{}
	in := make(chan int)
	out := make(chan int)

	parser := intcode.NewParser(input, in, out)

	x := 0
	y := 0
	direction := 0
	go func() {
		for true {
			c := 0
			if t, ok := tiles[fmt.Sprintf("%d %d", x, y)]; ok {
				c = t.color
			}
			in <- c
			c = <-out
			tiles[fmt.Sprintf("%d %d", x, y)] = Tile{x: x, y: y, color: c}
			fmt.Println(len(tiles))

			direction = (((direction + 180*<-out - 90) % 360) + 360) % 360
			switch direction {
			case 0:
				y++
			case 90:
				x++
			case 180:
				y--
			case 270:
				x--
			}
		}
	}()

	go func() {
		parser.Parse()
	}()

	for true {
	}

	return nil
}

// Day11Task2 ...
func Day11Task2() interface{} {
	const (
		day      = 11
		task     = 2
		testCode = 0
	)
	input := input.Get(day, task, testCode).([]int)
	tiles := map[string]Tile{}
	in := make(chan int)
	out := make(chan int)

	parser := intcode.NewParser(input, in, out)

	x := 0
	y := 0
	direction := 0
	tiles[fmt.Sprintf("%d %d", x, y)] = Tile{x: x, y: y, color: 1}
	go func() {
		for true {
			c := 0
			if t, ok := tiles[fmt.Sprintf("%d %d", x, y)]; ok {
				c = t.color
			}
			in <- c
			c = <-out
			tiles[fmt.Sprintf("%d %d", x, y)] = Tile{x: x, y: y, color: c}
			fmt.Println(len(tiles))
			// if len(tiles) >= 2416 {
			// draw(tiles)
			// }
			draw(tiles)

			d := 180*<-out - 90
			direction = (((direction + d) % 360) + 360) % 360
			fmt.Println(d)
			switch direction {
			case 0:
				y++
			case 90:
				x++
			case 180:
				y--
			case 270:
				x--
			}
		}
	}()

	go func() {
		parser.Parse()
	}()

	for true {
	}

	return nil
}

func draw(tiles map[string]Tile) {
	fmt.Println("---------------------------")
	var min_x, min_y, max_x, max_y int
	for _, t := range tiles {
		if t.x < min_x {
			min_x = t.x
		}
		if t.y < min_y {
			min_y = t.y
		}
		if t.x > max_x {
			max_x = t.x
		}
		if t.y > max_y {
			max_y = t.y
		}
	}

	for y := max_y; y >= min_y; y-- {
		line := ""
		for x := min_x; x <= max_x; x++ {
			if t, ok := tiles[fmt.Sprintf("%d %d", x, y)]; ok {
				if t.color == 1 {
					line += "*"
				} else {
					line += " "
				}

				continue
			}
			line += " "
		}
		fmt.Println(line)
	}
}
