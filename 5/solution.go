package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

type seat struct {
	row, col, ID int
}

func parseBSP(bspString string) seat {
	rowStart := 0
	rowEnd := 127
	colStart := 0
	colEnd := 7

	for i := 0; i < 7; i++ {
		pivot := ((rowStart + rowEnd + 1) / 2)
		if string(bspString[i]) == "B" {
			rowStart = pivot
		} else {
			rowEnd = pivot - 1
		}
	}

	for i := 7; i < 10; i++ {
		pivot := ((colStart + colEnd + 1) / 2)
		if string(bspString[i]) == "R" {
			colStart = pivot
		} else {
			colEnd = pivot - 1
		}
	}

	return seat{
		rowStart,
		colStart,
		(rowStart * 8) + colStart}
}

func part1(lines []string) int {
	maxID := -1
	for _, line := range lines {
		if maxID < parseBSP(line).ID {
			maxID = parseBSP(line).ID
		}
	}
	return maxID
}

func part2(lines []string) int {
	var seats []seat
	for _, line := range lines {
		seats = append(seats, parseBSP(line))
	}
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID < seats[j].ID
	})
	for i := 0; i < len(seats)-1; i++ {
		if seats[i].ID+2 == seats[i+1].ID {
			return seats[i].ID + 1
		}
	}
	return -1
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
