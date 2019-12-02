package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	part1()
}

func part2() {
	input := readInput("01/input")

	var result int
	for _, i := range input {
		i = i / 3
		i -= 2
		result += i
	}
	fmt.Println(result)
}

func part1() {
	input := readInput("01/input")

	var result int
	for _, i := range input {
		result += calculateFuel(i)
	}
	fmt.Println(result)
}


func calculateFuel(mass int) int {
	fuel := mass / 3
	fuel -= 2
	return fuel
}


func readInput(f string) (input []int) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}
	return
}