package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	input := readInput("01/input")

	var result int
	for _, i := range input {
		fuel := calculateFuel(i)
		result += fuel
		for fuel > 0 {
			fuel = calculateFuel(fuel)
			result += fuel
		}
	}
	fmt.Println(result)
}


func calculateFuel(mass int) int {
	fuel := mass / 3
	fuel -= 2
	if fuel < 0 {
		return 0
	}
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