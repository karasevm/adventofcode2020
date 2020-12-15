package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(digitString string, until int) int {
	currentDigit := 0
	digitsLastSeen := make(map[int]int)
	for index, digit := range strings.Split(digitString, ",") {
		val, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatal(err)
		}
		digitsLastSeen[val] = index
		currentDigit = val
	}
	startingDigitsCount := len(digitsLastSeen)
	for i := startingDigitsCount - 1; i < until-1; i++ {
		age := 0
		val, exists := digitsLastSeen[currentDigit]
		if exists {
			age = i - val
		}
		digitsLastSeen[currentDigit] = i
		currentDigit = age
	}
	return currentDigit
}

func part1(lines []string) int {
	return solve(lines[0], 2020)
}

func part2(lines []string) int {
	return solve(lines[0], 30000000)
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
