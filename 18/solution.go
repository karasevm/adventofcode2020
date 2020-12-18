package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solveExpressionLtr(expression string) int64 {
	result := int64(0)
	register := int64(0)
	isSum := true
	for i := 0; i < len(expression); i++ {
		if string(expression[i]) == "+" ||
			string(expression[i]) == "*" {
			isSum = (string(expression[i]) == "+")
		} else if string(expression[i]) != " " {
			if string(expression[i]) == "(" {
				pCount := 1
				j := i + 1
				for pCount != 0 {
					switch string(expression[j]) {
					case "(":
						pCount++
					case ")":
						pCount--
					}
					j++
				}
				register = solveExpressionLtr(expression[i+1 : j-1])
				i = j
			} else {
				val, err := strconv.ParseInt(string(expression[i]), 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				register = val
			}
			if isSum {
				result += register
			} else {
				result *= register
			}
		}
	}
	return result
}

func part1(lines []string) int64 {
	result := int64(0)
	for _, line := range lines {
		result += solveExpressionLtr(line)
	}
	return result
}

func solveExpressionInverse(expression string) int64 {
	result := int64(0)
	register := int64(0)
	multipliers := make([]int64, 0)
	for i := 0; i < len(expression); i++ {
		if string(expression[i]) == "+" ||
			string(expression[i]) == "*" {
			if string(expression[i]) == "*" {

				multipliers = append(multipliers, solveExpressionInverse(expression[i+2:]))
				break
			}
		} else if string(expression[i]) != " " {
			if string(expression[i]) == "(" {
				pCount := 1
				j := i + 1
				for pCount != 0 {
					switch string(expression[j]) {
					case "(":
						pCount++
					case ")":
						pCount--
					}
					j++
				}
				register = solveExpressionInverse(expression[i+1 : j-1])
				i = j
			} else {
				val, err := strconv.ParseInt(string(expression[i]), 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				register = val
			}
			result += register
		}
	}
	for _, multiplier := range multipliers {
		result *= multiplier
	}
	return result
}

func part2(lines []string) int64 {
	result := int64(0)
	for _, line := range lines {
		result += solveExpressionInverse(line)
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

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(lines), part2(lines))
}
