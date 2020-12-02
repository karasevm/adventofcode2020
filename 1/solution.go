package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(ints []int) int {
	for _, valA := range ints {
		for _, valB := range ints {
			if valA+valB == 2020 {
				return valA * valB
			}
		}
	}
	return -1
}

func part2(ints []int) int {
	for _, valA := range ints {
		for _, valB := range ints {
			for _, valC := range ints {
				if valA+valB+valC == 2020 {
					return valA * valB * valC
				}
			}
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
