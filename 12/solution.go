package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type direction struct {
	x, y int
}

// absInt takes an integer and returns it's absolute value
func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part1(lines []string) int {
	directionMultipliers := map[int]direction{0: {0, 1}, 90: {1, 0}, 180: {0, -1}, 270: {-1, 0}}
	currentDirection := 90
	currentX := 0
	currentY := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch string(line[:1]) {
		case "N":
			currentY += value
		case "S":
			currentY -= value
		case "E":
			currentX += value
		case "W":
			currentX -= value
		case "L":
			currentDirection = (currentDirection - value + 360) % 360
		case "R":
			currentDirection = (currentDirection + value) % 360
		case "F":
			currentX += value * directionMultipliers[currentDirection].x
			currentY += value * directionMultipliers[currentDirection].y
		}

	}
	return absInt(currentX) + absInt(currentY)
}

func part2(lines []string) int {
	shipX := 0
	shipY := 0
	waypointX := 10
	waypointY := 1

	for _, line := range lines {
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch string(line[:1]) {
		case "N":
			waypointY += value
		case "S":
			waypointY -= value
		case "E":
			waypointX += value
		case "W":
			waypointX -= value
		case "R":
			for i := 0; i < value/90; i++ {
				waypointX, waypointY = waypointY, -waypointX
			}
		case "L":
			for i := 0; i < value/90; i++ {
				waypointX, waypointY = -waypointY, waypointX
			}
		case "F":
			shipX += waypointX * value
			shipY += waypointY * value
		}

	}
	return absInt(shipX) + absInt(shipY)
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
