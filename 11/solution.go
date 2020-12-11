package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

func copyStringSlice(inputSlice []string) []string {
	var outputString []string
	for _, inputString := range inputSlice {
		outputString = append(outputString, inputString)
	}
	return outputString
}

type direction struct {
	y, x int
}

func countOccupied(x int, y int, seats []string) int {
	directions := []direction{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	count := 0
	for i := len(directions) - 1; i >= 0; i-- {
		yToCheck := y + directions[i].y
		xToCheck := x + directions[i].x

		// Don't check if index out of range
		if yToCheck < 0 || yToCheck >= len(seats) ||
			xToCheck < 0 || xToCheck >= len(seats[0]) {
			continue
		}
		if string(seats[yToCheck][xToCheck]) == "#" {
			count++
		}
	}
	return count
}

// compare compares two slices of strings
func compare(sliceA []string, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, element := range sliceA {
		if element != sliceB[index] {
			return false
		}
	}
	return true
}

func countTotalOccupied(seats []string) int {
	count := 0
	for _, line := range seats {
		for _, seat := range line {
			if string(seat) == "#" {
				count++
			}
		}
	}
	return count
}

func part1(lines []string) int {
	sliceWrite := copyStringSlice(lines)
	var sliceRead []string

	for !compare(sliceRead, sliceWrite) {
		sliceRead = copyStringSlice(sliceWrite)
		sliceWrite = nil

		for y := 0; y < len(lines); y++ {
			newString := ""
			for x := 0; x < len(lines[0]); x++ {
				if string(sliceRead[y][x]) == "L" {
					count := countOccupied(x, y, sliceRead)
					if count == 0 {
						newString += "#"
					} else {
						newString += "L"
					}
				} else if string(sliceRead[y][x]) == "#" {
					count := countOccupied(x, y, sliceRead)
					if count >= 4 {
						newString += "L"
					} else {
						newString += "#"
					}
				} else {
					newString += "."
				}
			}
			sliceWrite = append(sliceWrite, newString)
		}
	}
	return countTotalOccupied(sliceWrite)
}

func countOccupiedLines(x int, y int, seats []string) int {
	directions := []direction{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	count := 0
	for multiplier := 1; multiplier <= int(math.Ceil(math.Sqrt(float64(x*x+y*y)))); multiplier++ {
		for i := len(directions) - 1; i >= 0; i-- {
			yToCheck := y + directions[i].y*multiplier
			xToCheck := x + directions[i].x*multiplier

			// Don't check if index out of range
			if yToCheck < 0 || yToCheck >= len(seats) ||
				xToCheck < 0 || xToCheck >= len(seats[0]) {
				continue
			}
			if string(seats[yToCheck][xToCheck]) == "#" ||
				string(seats[yToCheck][xToCheck]) == "L" {
				directions = append(directions[:i], directions[i+1:]...)
				if string(seats[yToCheck][xToCheck]) == "#" {
					count++
				}
			}
		}
	}
	return count
}

func part2(lines []string) int {
	sliceWrite := copyStringSlice(lines)
	var sliceRead []string

	for !compare(sliceRead, sliceWrite) {
		sliceRead = copyStringSlice(sliceWrite)
		sliceWrite = nil
		for y := 0; y < len(lines); y++ {
			newString := ""
			for x := 0; x < len(lines[0]); x++ {
				if string(sliceRead[y][x]) == "L" {
					if countOccupiedLines(x, y, sliceRead) == 0 {
						newString += "#"
					} else {
						newString += "L"
					}
				} else if string(sliceRead[y][x]) == "#" {
					if countOccupiedLines(x, y, sliceRead) >= 5 {
						newString += "L"
					} else {
						newString += "#"
					}
				} else {
					newString += "."
				}
			}
			sliceWrite = append(sliceWrite, newString)
		}
	}
	return countTotalOccupied(sliceWrite)
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
