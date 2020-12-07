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

type bag struct {
	parents  []string       //parent names
	children map[string]int //children name and count
}

func findUniqueParents(bags map[string]*bag, targetBag string) map[string]struct{} {
	uniqueParents := make(map[string]struct{})
	for _, parent := range bags[targetBag].parents {
		uniqueParents[parent] = struct{}{}
		tempParents := findUniqueParents(bags, parent)
		for k, v := range tempParents {
			uniqueParents[k] = v
		}
	}
	return uniqueParents
}

func findTotalChildrenCount(bags map[string]*bag, targetBag string) int {
	totalCount := 0
	for child, count := range bags[targetBag].children {
		totalCount += count * (findTotalChildrenCount(bags, child) + 1)
	}
	return totalCount
}

func parseBags(lines []string) map[string]*bag {
	bags := make(map[string]*bag)
	expression := regexp.MustCompile(`(\d \w+\s\w+)\sbags?`)

	for _, line := range lines {
		splitLine := strings.Split(line, " bags contain ")

		// Initialize bag struct if it does not exist
		if _, exists := bags[splitLine[0]]; !exists {
			bags[splitLine[0]] = &bag{make([]string, 0), make(map[string]int)}
		}

		children := expression.FindAllStringSubmatch(splitLine[1], -1)

		for _, child := range children {
			count, err := strconv.Atoi(child[1][:1])
			if err != nil {
				log.Fatal(err)
			}
			bags[splitLine[0]].children[child[1][2:]] = count

			// Initialize child struct to populate it's parents
			if _, exists := bags[child[1][2:]]; !exists {
				bags[child[1][2:]] = &bag{make([]string, 0), make(map[string]int)}
			}

			bags[child[1][2:]].parents = append(bags[child[1][2:]].parents, splitLine[0])
		}
	}
	return bags
}

func part1(lines []string) int {
	return len(findUniqueParents(parseBags(lines), "shiny gold"))

}

func part2(lines []string) int {
	return findTotalChildrenCount(parseBags(lines), "shiny gold")
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
