package main

import "fmt"

func main() {
	var valid int
	for i := 357253; i < 892942; i++ {
		if checkValid(i) {
			valid++
		}
	}

	fmt.Println(valid)
}

func checkValid(value int) bool {
	var split []int
	for i := 100000; i >= 10; i = i / 10 {
		split = append(split, value/i%10)
	}
	split = append(split, value%10)

	var current = split[0]
	var repeat bool
	for i, v := range split {
		if v < current {
			return false
		}
		if v == current && i != 0 {
			repeat = true
		}
		current = v
	}
	return repeat
}
