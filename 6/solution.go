package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func countUnique(lines []string) int {
	letters := make(map[string]struct{})
	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			sum += len(letters)
			letters = make(map[string]struct{})
			continue
		}
		for _, letter := range line {
			letters[string(letter)] = struct{}{}
		}
	}
	// For the last group in input
	sum += len(letters)
	return sum
}

func intersection(set1 map[string]bool, set2 map[string]bool) map[string]bool {
	resultingIntersection := make(map[string]bool)
	for k := range set1 {
		if set2[k] {
			resultingIntersection[k] = true
		}
	}
	return resultingIntersection
}

func groupIntersection(group []map[string]bool) map[string]bool {
	result := group[0]
	for _, line := range group[1:] {
		result = intersection(result, line)
	}
	return result
}

func countCommon(lines []string) int {
	sum := 0
	var lineGroup []map[string]bool
	for _, line := range lines {
		if len(line) == 0 {
			sum += len(groupIntersection(lineGroup))
			lineGroup = nil
			continue
		}

		letters := make(map[string]bool)
		for _, letter := range line {
			letters[string(letter)] = true
		}
		lineGroup = append(lineGroup, letters)
	}
	// For the last group in input
	sum += len(groupIntersection(lineGroup))
	return sum

}

func part1(lines []string) int {
	return countUnique(lines)
}

func part2(lines []string) int {
	return countCommon(lines)
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
