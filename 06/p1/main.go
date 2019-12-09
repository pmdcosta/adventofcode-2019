package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	world := readInput("06/p1/input")
	fmt.Println(traverse(world, "COM", 0))
}

func readInput(f string) map[string][]string {
	var world = make(map[string][]string)

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		world[s[0]] = append(world[s[0]], s[1])
	}
	return world
}

func traverse(world map[string][]string, planet string, level int) int {
	var sum int
	for _, p := range world[planet] {
		sum += traverse(world, p, level+1)
	}
	return sum + level
}
