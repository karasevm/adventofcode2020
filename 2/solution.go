package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	expression := regexp.MustCompile(`(\d+)\-(\d+)\ ([a-z]):\ ([a-z]+)$`)
	validCount := 0
	for _, line := range lines {
		match := expression.FindAllStringSubmatch(line, -1)
		count := strings.Count(match[0][4], match[0][3])

		min, err := strconv.Atoi(match[0][1])
		if err != nil {
			log.Fatal(err)
		}

		max, err := strconv.Atoi(match[0][2])
		if err != nil {
			log.Fatal(err)
		}

		if count >= min && count <= max {
			validCount++
		}
	}
	return validCount
}

func part2(lines []string) int {
	expression := regexp.MustCompile(`(\d+)\-(\d+)\ ([a-z]):\ ([a-z]+)$`)
	validCount := 0
	for _, line := range lines {
		match := expression.FindAllStringSubmatch(line, -1)

		first, err := strconv.Atoi(match[0][1])
		if err != nil {
			log.Fatal(err)
		}

		second, err := strconv.Atoi(match[0][2])
		if err != nil {
			log.Fatal(err)
		}

		if (match[0][3] == string(match[0][4][first-1]) && match[0][3] != string(match[0][4][second-1])) ||
			(match[0][3] != string(match[0][4][first-1]) && match[0][3] == string(match[0][4][second-1])) {
			validCount++
		}
	}
	return validCount
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
