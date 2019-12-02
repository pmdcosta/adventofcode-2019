package main

import "fmt"

var base = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 6, 19, 23, 2, 6, 23, 27, 1, 5, 27, 31, 2, 31, 9, 35, 1, 35, 5, 39, 1, 39, 5, 43, 1, 43, 10, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 6, 59, 1, 5, 59, 63, 2, 63, 6, 67, 1, 5, 67, 71, 1, 71, 6, 75, 2, 75, 10, 79, 1, 79, 5, 83, 2, 83, 6, 87, 1, 87, 5, 91, 2, 9, 91, 95, 1, 95, 6, 99, 2, 9, 99, 103, 2, 9, 103, 107, 1, 5, 107, 111, 1, 111, 5, 115, 1, 115, 13, 119, 1, 13, 119, 123, 2, 6, 123, 127, 1, 5, 127, 131, 1, 9, 131, 135, 1, 135, 9, 139, 2, 139, 6, 143, 1, 143, 5, 147, 2, 147, 6, 151, 1, 5, 151, 155, 2, 6, 155, 159, 1, 159, 2, 163, 1, 9, 163, 0, 99, 2, 0, 14, 0}

func main() {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			if val := findValue(19690720, i, j); val != 0 {
				fmt.Println("found:", val)
				return
			}
		}
	}
	panic("not found")
}

func findValue(value, noun, verb int) int {
	current := make([]int, len(base))
	copy(current, base)
	current[1] = noun
	current[2] = verb

	if runComputer(current) == value {
		return 100*current[1] + current[2]
	}
	return 0
}

func runComputer(input []int) int {
	for i := 0; i < len(input); i = i + 4 {
		if input[i] == 99 {
			break
		}
		var sub = input[i : i+4]
		handleOp(input, sub[0], input[sub[1]], input[sub[2]], sub[3])
	}
	return input[0]
}

func handleOp(input []int, op, arg1, arg2, pos int) {
	switch op {
	case 1:
		input[pos] = arg1 + arg2
	case 2:
		input[pos] = arg1 * arg2
	}
}