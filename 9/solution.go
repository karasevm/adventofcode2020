package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func findPair(preamble []int, sum int) bool {
	for _, p1 := range preamble {
		for _, p2 := range preamble {
			if p2 != p1 && sum-p1 == p2 {
				return true
			}
		}
	}
	return false
}

func part1(lines []int, preambleSize int) int {
	var preamble []int

	for key, line := range lines[preambleSize:] {
		preamble = lines[key : key+preambleSize]
		if !findPair(preamble, line) {
			return line
		}

	}
	return 0
}

func part2(lines []int, preambleSize int) int {
	badNumber := part1(lines, preambleSize)
	for i := range lines {
		sum := 0
		for k := i; k < len(lines)-i; k++ {
			sum += lines[k]
			if sum > badNumber {
				break
			}
			if sum == badNumber {
				newSlice := append([]int{}, lines[i:k]...)
				sort.Slice(newSlice, func(i, j int) bool {
					return newSlice[i] < newSlice[j]
				})

				return newSlice[0] + newSlice[len(newSlice)-1]
			}

		}
	}
	return 0
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

	var inputInts []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputInts = append(inputInts, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(inputInts, 25), part2(inputInts, 25))
}
