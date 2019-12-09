package main

import (
	"fmt"
)

var base = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 55, 68, 93, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 2, 9, 101, 5, 9, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 101, 4, 9, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}
var combinations [][]int

type computer struct {
	input  chan int
	output chan int
}

func (c *computer) start(phase int) {
	var state = make([]int, len(base))
	copy(state, base)
	go runComputer(state, phase, c.input, c.output)
}

func main() {
	perm([]int{0, 1, 2, 3, 4}, 0)

	var out int
	for _, c := range combinations {
		current := runPipeline(c)
		if current > out {
			out = current
		}
	}
	fmt.Println(out)
}

func perm(a []int, i int) {
	if i > len(a) {
		var s = make([]int, len(a))
		copy(s, a)
		combinations = append(combinations, s)
		return
	}
	perm(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func runPipeline(phases []int) int {
	var machines []computer

	c := computer{input: make(chan int), output: make(chan int)}
	c.start(phases[0])
	machines = append(machines, c)

	for i := 1; i < len(phases); i++ {
		c := computer{input: machines[i-1].output, output: make(chan int)}
		c.start(phases[i])
		machines = append(machines, c)
	}

	machines[0].input <- 0
	return <-machines[len(phases)-1].output
}

func runComputer(state []int, phase int, input chan int, output chan int) {
	var o int
	var setOutput = func(v int) {
		o = v
		output <- v
	}

	var phaseUsed bool
	var getInput = func() int {
		if !phaseUsed {
			phaseUsed = true
			return phase
		}
		return <-input
	}

	for i := 0; i < len(state); {
		val, done := handleOp(state, i, getInput, setOutput)
		if done {
			return
		}
		i = val
	}
}

func handleOp(state []int, cursor int, in func() int, setOutput func(int)) (int, bool) {
	op := state[cursor] % 100
	switch op {
	case 99:
		return 0, true
	case 1:
		return cursor + sum(state, cursor), false
	case 2:
		return cursor + mul(state, cursor), false
	case 3:
		return cursor + input(state, cursor, in), false
	case 4:
		c, out := output(state, cursor)
		setOutput(out)
		return cursor + c, false
	case 5:
		return jumpTrue(state, cursor), false
	case 6:
		return jumpFalse(state, cursor), false
	case 7:
		return cursor + lessThan(state, cursor), false
	case 8:
		return cursor + equal(state, cursor), false
	}
	panic(fmt.Sprintf("invalid operation: %d (%d)", op, state[cursor]))
}

func sum(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	state[state[cursor+3]] = arg1 + arg2
	return 4
}

func mul(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	state[state[cursor+3]] = arg1 * arg2
	return 4
}

func input(state []int, cursor int, input func() int) int {
	if state[cursor]/100%10 == 0 {
		state[state[cursor+1]] = input()
	} else {
		state[cursor+1] = input()
	}
	return 2
}

func output(state []int, cursor int) (int, int) {
	var arg1 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	return 2, arg1
}

func jumpTrue(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	if arg1 != 0 {
		return arg2
	}
	return cursor + 3
}

func jumpFalse(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	if arg1 == 0 {
		return arg2
	}
	return cursor + 3
}

func lessThan(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	if arg1 < arg2 {
		state[state[cursor+3]] = 1
	} else {
		state[state[cursor+3]] = 0
	}
	return 4
}

func equal(state []int, cursor int) int {
	var arg1, arg2 int
	if state[cursor]/100%10 == 0 {
		arg1 = state[state[cursor+1]]
	} else {
		arg1 = state[cursor+1]
	}
	if state[cursor]/1000%10 == 0 {
		arg2 = state[state[cursor+2]]
	} else {
		arg2 = state[cursor+2]
	}

	if arg1 == arg2 {
		state[state[cursor+3]] = 1
	} else {
		state[state[cursor+3]] = 0
	}
	return 4
}
