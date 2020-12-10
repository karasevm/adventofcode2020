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

func part1(lines []int) int {
	diff1Count := 0
	diff3Count := 0

	sortedLines := append([]int{0}, lines...)
	sort.Slice(sortedLines, func(i, j int) bool {
		return sortedLines[i] < sortedLines[j]
	})

	sortedLines = append(sortedLines, sortedLines[len(sortedLines)-1]+3)
	for i := 0; i < len(sortedLines)-1; i++ {
		switch sortedLines[i+1] - sortedLines[i] {
		case 1:
			diff1Count++
		case 3:
			diff3Count++
		}
	}
	return diff1Count * diff3Count
}

func countWays(number int) int {
	switch number {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		return countWays(number-1) + countWays(number-2) + countWays(number-3)
	}
}

// Get the number of all possible combinations by finding all groups of cosecutive 1 diffs,
// taking their length, and finding all possible combinations of 1,2,3 that sum up to that length
// Multiply them to get all possible combinations
// Can be made more efficient by memoizing the combination count function, but works fast enough as is
func part2(lines []int) int {
	sortedLines := append([]int{0}, lines...)
	sort.Slice(sortedLines, func(i, j int) bool {
		return sortedLines[i] < sortedLines[j]
	})

	sortedLines = append(sortedLines, sortedLines[len(sortedLines)-1]+3)

	result := 1
	oneCount := 0
	for i := 0; i < len(sortedLines)-1; i++ {
		switch sortedLines[i+1] - sortedLines[i] {
		case 1:
			oneCount++
		case 3:
			if oneCount > 0 {
				result *= countWays(oneCount)
			}
			oneCount = 0
		}
	}
	return result
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
	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(inputInts), part2(inputInts))
}
