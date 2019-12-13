package intcode

import "fmt"

// Parser ...
type Parser struct {
	pc     int
	input  []int
	stdin  chan int
	stdout chan int
	signal chan int
}

// NewParser ...
func NewParser(input []int, stdin, stdout chan int) *Parser {
	in := make([]int, len(input)*10)
	for i, a := range input {
		in[i] = a
	}
	return &Parser{
		input:  in,
		stdin:  stdin,
		stdout: stdout,
		pc:     0,
	}
}

func NewParserWithInputSignaling(input []int, stdin, stdout chan int, signal chan int) *Parser {
	in := make([]int, len(input)*10)
	for i, a := range input {
		in[i] = a
	}
	return &Parser{
		input:  in,
		stdin:  stdin,
		stdout: stdout,
		pc:     0,
		signal: signal,
	}
}

// Parse ...
func (p *Parser) Parse() []int {
	i := 0
	// n := len(p.input)
	relBase := 0
	for true {
		if p.input[i] == 99 {
			// close(p.stdin)
			// close(p.stdout)
			break
		}
		opcode := p.input[i] % 100
		parMode1 := (p.input[i] / 100) % 10
		parMode2 := (p.input[i] / 1000) % 10
		parMode3 := (p.input[i] / 10000) % 10
		a := (i + 1)
		b := (i + 2)
		c := (i + 3)
		if parMode1 == 0 {
			a = p.input[a]
		}
		if parMode2 == 0 {
			b = p.input[b]
		}
		if parMode3 == 0 {
			c = p.input[c]
		}
		if parMode1 == 2 {
			a = p.input[a] + relBase
		}
		if parMode2 == 2 {
			b = p.input[b] + relBase
		}
		if parMode3 == 2 {
			c = p.input[c] + relBase
		}
		switch opcode {
		case 1:
			p.input[c] = p.input[a] + p.input[b]
			i += 4
		case 2:
			p.input[c] = p.input[a] * p.input[b]
			i += 4
		case 3:
			if p.signal != nil {
				p.signal <- 1
			}
			if p.stdin != nil {
				p.input[a] = <-p.stdin
			} else {
				fmt.Print("< ")
				fmt.Scanf("%d\n", &p.input[a])
			}
			i += 2
		case 4:
			if p.stdout != nil {
				p.stdout <- p.input[a]
			} else {
				fmt.Println(">", p.input[a])
			}
			i += 2
		case 5:
			if p.input[a] != 0 {
				i = p.input[b]
			} else {
				i += 3
			}
		case 6:
			if p.input[a] == 0 {
				i = p.input[b]
			} else {
				i += 3
			}
		case 7:
			if p.input[a] < p.input[b] {
				p.input[c] = 1
			} else {
				p.input[c] = 0
			}
			i += 4
		case 8:
			if p.input[a] == p.input[b] {
				p.input[c] = 1
			} else {
				p.input[c] = 0
			}
			i += 4
		case 9:
			relBase += p.input[a]
			i += 2
		}
	}

	return p.input
}
