package day

import "github.com/rafalrothenberger/adventofcode2019/internal/input"

import "github.com/rafalrothenberger/adventofcode2019/internal/intcode"

import "fmt"

import "time"

import "github.com/rafalrothenberger/adventofcode2019/internal/smath"

// Day13Task1 ...
func Day13Task1() interface{} {
	const (
		day      = 13
		task     = 1
		testCode = 0
	)
	input := input.Get(day, task, testCode).([]int)
	out := make(chan int, 1000000)
	intcode.NewParser(input, nil, out).Parse()
	close(out)
	j := 0
	res := 0
	for o := range out {
		if j%3 == 2 && o == 2 {
			res++
		}
		j++
	}
	return res
}

// Day13Task2 ...
func Day13Task2() interface{} {
	paddleX := 0
	ballX := 0
	const (
		day      = 13
		task     = 2
		testCode = 0
	)
	input := input.Get(day, task, testCode).([]int)
	input[0] = 2
	out := make(chan int, 1000000)
	signal := make(chan int, 1000000)
	in := make(chan int, 1000)
	initial := true
	board := make([]string, 21)
	res := 0

	go func() {
		for true {
			<-signal
			time.Sleep(time.Second / 50)

			in <- (ballX - paddleX) / smath.Max(smath.Abs(ballX-paddleX), 1)
		}
	}()

	go func() {
		for true {
			x := <-out
			y := <-out
			r := <-out
			if x == -1 && y == 0 {
				initial = false
				drawBoard(board)
				fmt.Println(res)
				res += r
				fmt.Println("Score", res)
				fmt.Println("Score", r)
				continue
			}
			if initial {
				switch r {
				case 0:
					board[y] += " "
				case 1:
					board[y] += "|"
				case 2:
					board[y] += "X"
					// res++
				case 3:
					board[y] += "-"
					paddleX = x
				case 4:
					board[y] += "O"
					ballX = x
				}
				continue
			}

			switch r {
			case 0:
				board[y] = board[y][:x] + " " + board[y][x+1:]
			case 1:
				board[y] = board[y][:x] + "|" + board[y][x+1:]
			case 2:
				board[y] = board[y][:x] + "X" + board[y][x+1:]
			case 3:
				board[y] = board[y][:x] + "-" + board[y][x+1:]
				paddleX = x
			case 4:
				board[y] = board[y][:x] + "O" + board[y][x+1:]
				ballX = x
			}
			drawBoard(board)

		}
	}()
	intcode.NewParserWithInputSignaling(input, in, out, signal).Parse()
	// close(out)

	for true {
	}
	return nil
}

func drawBoard(board []string) {
	// fmt.Println("123")
	for j := 0; j < 21; j++ {
		fmt.Println(board[j])
	}
}
