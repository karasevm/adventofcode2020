package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func treeFinder(grid []string, right int, down int) int {
	treeCount := 0
	x := 0
	y := 0
	lineLength := len(grid[0])
	for y < len(grid) {
		if string(grid[y][x]) == "#" {
			treeCount++
		}
		x = (x + right) % lineLength
		y += down
	}
	return treeCount
}

func part1(lines []string) int {
	return treeFinder(lines, 3, 1)
}

func part2(lines []string) int {
	return treeFinder(lines, 1, 1) *
		treeFinder(lines, 3, 1) *
		treeFinder(lines, 5, 1) *
		treeFinder(lines, 7, 1) *
		treeFinder(lines, 1, 2)
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("Provide the path to input")
	}
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(lines), part2(lines))
}
