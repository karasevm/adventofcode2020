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

// executeOperations executes opcodes provided in lines
// It returns exit code (0 - finished, 1 - looped), and current accumulator value
func executeOperations(lines []string) (int, int) {
	hitInstructions := make(map[int]bool)
	accumulator := 0
	for i := 0; i < len(lines); {
		arg, err := strconv.Atoi(lines[i][4:])
		if err != nil {
			log.Fatal(err)
		}
		if hitInstructions[i] {
			return 1, accumulator
		}
		hitInstructions[i] = true
		switch lines[i][:3] {
		case "jmp":
			i += arg
		case "acc":
			accumulator += arg
			fallthrough
		case "nop":
			i++
		}
	}
	return 0, accumulator
}

func part1(lines []string) int {
	_, acc := executeOperations(lines)
	return acc
}

func part2(lines []string) int {
	for i := 0; i < len(lines); i++ {
		var tempLine string
		switch lines[i][:3] {
		case "jmp":
			tempLine = strings.Replace(lines[i], "jmp", "nop", -1)
		case "nop":
			tempLine = strings.Replace(lines[i], "nop", "jmp", -1)
		default:
			continue
		}
		tempLines := append([]string{}, lines[:i]...)
		tempLines = append(tempLines, tempLine)
		tempLines = append(tempLines, lines[i+1:]...)
		status, acc := executeOperations(tempLines)

		if status == 0 {
			return acc
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

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(lines), part2(lines))
}
