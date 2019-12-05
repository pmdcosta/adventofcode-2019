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
	var repeatNumber int
	var repeat []int
	for i := 1; i < len(split); i++ {
		v := split[i]
		if v < current {
			return false
		}

		if v == current {
			repeatNumber++
		} else {
			repeat = append(repeat, repeatNumber)
			repeatNumber = 0
		}
		current = v
	}
	repeat = append(repeat, repeatNumber)

	for _, r := range repeat {
		if r == 1 {
			return true
		}
	}
	return false
}
