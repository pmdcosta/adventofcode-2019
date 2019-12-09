package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	world := readInput("06/p2/input")
	traverse(world, "YOU", 0, "")
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
		world[s[1]] = append(world[s[1]], s[0])
	}
	return world
}

func traverse(world map[string][]string, planet string, level int, source string) {
	for _, p := range world[planet] {
		if p == "SAN" {
			fmt.Println(level - 1)
			os.Exit(0)
		}
		if p == planet || p == source {
			continue
		}
		traverse(world, p, level+1, planet)
	}
}
